package crud

import (
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
	"log"
)

/**
  abre a conexão com o banco de dados
 */
func GetConnection() (*sqlx.DB) {
	db, err := sqlx.Open("postgres", "postgres://root:root@postgresql-server.dev/wiki?sslmode=disable")
	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(95)
	if err != nil{
		log.Println("m=GetConnection,msg=connection has failed", err)
	}
	return db
}