package main

import (
	"net/http"
	"log"
	"github.com/mageddo/go-examples/gowiki-vestigo-3rd-sql-dao/webapp/controller"
)

func main() {
	controller.LoadControllers()
	log.Println("starting app")
	log.Fatal("can not start app: ", http.ListenAndServe(":8989", controller.App))
}