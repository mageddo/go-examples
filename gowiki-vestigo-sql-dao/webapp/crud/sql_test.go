package crud

import (
	"testing"
	"database/sql"
	"log"
)


func TestGetConnection(t *testing.T) {
	con, err := GetConnection()
	if(err != nil){
		t.Error("the connection has failed", err)
	}
	con.Close()
}

func TestTransaction(t *testing.T) {
	_, err := Transaction(func(tx *sql.Tx) (*interface {}, *sql.Stmt, error) {

		log.Println("preparing querie")
		stm, err := tx.Prepare("SELECT 1;")
		if err != nil{
			log.Println("error prepare SQL")
			return nil, nil, err
		}
		log.Println("executing querie")
		_, err = stm.Exec()
		log.Println("returning status querie")

		return nil, stm, err
	})

	if err != nil {
		t.Error("erro ao efetuar a transação", err)
	}

}

func TestRun(t *testing.T) {
	_, err := Run(func(db *sql.DB) (*interface{}, error) {

		_, err := db.Exec("SELECT 1;")
		return nil, err

	})

	if err != nil {
		t.Error("could not execute querie", err)
	}
}