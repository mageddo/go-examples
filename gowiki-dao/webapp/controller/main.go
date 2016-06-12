package controller

import "log"

var msg = func()int{
	log.Println("loading controllers....")
	return 0
}()

func LoadControllers(){
	log.Println("controllers are loaded!")
}