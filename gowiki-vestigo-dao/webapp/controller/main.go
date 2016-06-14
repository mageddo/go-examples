package controller

import (
	"log"
	"github.com/mageddo/vestigo"
)


var App = func() *vestigo.Router {
	log.Println("loading controllers....")
	return vestigo.NewRouter()
}()

func LoadControllers(){
	log.Println("controllers are loaded!")
}