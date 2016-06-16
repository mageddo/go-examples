package main

import "log"

func CbCaller(fn func()){

	defer func(){
		log.Println("defer has happen")
	}()

	log.Println("before call");
	fn()
	log.Println("after call");

}

func CbReturnCaller( fn func() (interface{}) ) interface{} {

	defer func(){
		log.Println("defer was runned")
	}()

	log.Println("before call");
	msg := fn()
	log.Println("after call");
	return msg

}


func main(){

	CbCaller(func(){
		log.Println(">> function called")
	})

	log.Println("==============================")

	msg := CbReturnCaller(func() interface{} {
		log.Println(">> function with return called")
		return "the cb msg"
	})
	log.Println("the msg was=", msg)


}
