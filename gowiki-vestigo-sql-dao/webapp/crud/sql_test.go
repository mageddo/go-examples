package crud

import (
	"testing"
	"database/sql"
	"log"
	"reflect"
)


func TestGetConnection(t *testing.T) {
	con, err := GetConnection()
	if(err != nil){
		t.Error("the connection has failed", err)
	}
	con.Close()
}

func TestTransaction(t *testing.T) {
	_, err := Transaction(func(tx *sql.Tx) (interface {}, *sql.Stmt, error) {

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
	ra, err := Run(func(db *sql.DB) (interface{}, error) {

		r, err := db.Exec("SELECT 10;")
		n,_ := r.RowsAffected()
		return n, err

	})

	if err != nil {
		t.Error("could not execute querie", err)
	}

	log.Println("ra type: ", reflect.TypeOf(ra))

	if ra != int64(1) {
		t.Error("affect rows number is not correct", ra)
	}
}