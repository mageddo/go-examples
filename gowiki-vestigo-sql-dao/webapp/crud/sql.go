package crud

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

/**
  abre a conexão com o banco de dados
 */
func GetConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://root:root@postgresql-server.dev/wiki?sslmode=disable")
	//db.SetMaxIdleConns(8)
	db.SetMaxOpenConns(95)
	return db, err
}

/**
  the the specified querie then close the connection
 */
func Run(fn func(db *sql.DB) (interface{}, error) ) (interface{},error) {

	db, err := GetConnection()
	if err != nil {
		log.Println("could not open connection", err)
		return nil, err
	}
	defer func(){
		log.Println("m=Run,msg=closing connection")
		db.Close()
	}()

	log.Println("m=Run, msg=calling cb")
	i, err := fn(db)
	log.Println("m=Run, msg=cb called")
	return i,err
}

/**
  create a transaction, try run the querie, then commit or rollback the transaction and close the connection
 */
func Transaction(fn func (tx *sql.Tx) (interface{}, *sql.Stmt, error)) (interface{}, error) {

	defer func(){
		str := recover()
		if str != nil{
			log.Println("panic occur: ", str)
		}
	}()

	log.Println("opening connection database")

	db, err := GetConnection()
	if err != nil {
		log.Println("could not open connection", err)
		return nil, err
	}
	defer func(){
		log.Println("m=Transaction,msg=closing connection")
		db.Close()
	}()

	log.Println("begin transaction")
	tx,err := db.Begin()
	if err != nil {
		log.Println("could not open transaction", err)
		return nil, err
	}

	log.Println("calling callback")
	it, stm, err := fn(tx)
	if err != nil {
		tx.Rollback()
		log.Println("could not run stm", err)
		return nil, err
	}
	defer func(){
		log.Println("closed stm")
		stm.Close()
	}()
	log.Println("callback successfull called, commiting")
	err = tx.Commit()
	log.Println("transaction commited")
	if err != nil {
		log.Println("could not commit transaction", err)
		return nil, err
	}

	return it, err
}