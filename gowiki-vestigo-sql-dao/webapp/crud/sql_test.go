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

		stm, err := tx.Prepare("SELECT ?;")
		stm.Exec(1)

		return nil, stm, err
	})

	if err != nil {
		log.Fatal("erro ao efetuar a transação", err)
	}

}
