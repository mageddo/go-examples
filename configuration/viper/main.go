package main

import (
	"github.com/spf13/viper"
	"fmt"
	"log"
)

func main(){

	viper.SetConfigName("application-prod")

	// o path que ele achar primeiro, ele ja para, ou seja, nao mergia os diretorios
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("Fatal error config file: %s \n", err))
	}


	log.Printf("name=%s", viper.Get("name"))
	log.Printf("version=%s", viper.Get("version"))
	log.Printf("user.name=%s", viper.Get("user.name"))
	log.Printf("url=%s", viper.Get("url"))

}
