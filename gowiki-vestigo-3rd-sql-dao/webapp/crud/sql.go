package crud

import (
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

var (
	pool *sqlx.DB
	err error
	)
func init(){
	pool, err = sqlx.Open("postgres", "postgres://root:root@postgresql-server.dev/wiki?sslmode=disable")
	pool.SetMaxIdleConns(0)
	pool.SetConnMaxLifetime(2 * time.Minute)
	pool.SetMaxOpenConns(95)
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