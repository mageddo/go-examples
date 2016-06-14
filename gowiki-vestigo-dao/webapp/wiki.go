package main

import (
	"net/http"
	"log"
	"github.com/mageddo/go-examples/gowiki-dao/webapp/controller"
)

func main() {
	controller.LoadControllers()
	log.Println("starting app")
	http.ListenAndServe(":8080", nil)
}