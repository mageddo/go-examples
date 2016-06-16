package crud

import (
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
	"log"
)

var pool *sqlx.DB
func init(){
	db, err := sqlx.Open("postgres", "postgres://root:root@postgresql-server.dev/wiki?sslmode=disable")
	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(95)
	if err != nil{
		log.Println("m=GetPool,msg=connection has failed", err)
	}
}

/**
  abre a conexão com o banco de dados
 */
func GetPool() (*sqlx.DB) {
	return pool
}