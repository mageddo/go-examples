package crud

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func GetConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://root:root@postgresql-server.dev/pqgotest?sslmode=disable")
	return db, err
}

func Run(fn func(db *sql.DB) (*interface{}, error) ) (interface{},error) {

	db, err := GetConnection()
	if err != nil {
		log.Fatal("could not open connection", err)
		return nil, err
	}
	defer db.Close()

	return fn(db)
}

func Transaction(fn func (db *sql.DB) (*interface{}, *sql.Stmt, error)) (*interface{}, error) {

	db, err := GetConnection()
	if err != nil {
		log.Fatal("could not open connection", err)
		return nil, err
	}
	defer db.Close()

	tx,err := db.Begin()
	if err != nil {
		log.Fatal("could not open transaction", err)
		return nil, err
	}

	stm, err := fn(tx)
	if err != nil {
		tx.Rollback()
		log.Fatal("could not run stm", err)
		return nil, err
	}
	defer stm.Close()

	err = tx.Commit()
	if err != nil {
		log.Fatal("could not commit transaction", err)
		return nil, err
	}


}