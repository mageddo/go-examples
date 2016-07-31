package main

import (
	"log"
	"net/http"
	"fmt"
	"time"
	"math/rand"
	"github.com/mageddo/go-examples/id"
)

var i int = 0;
func handler(w http.ResponseWriter, r *http.Request) {
	i++
	var d int = i
	fmt.Printf("Hi there, i=%d!\n", d)
	time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
	fmt.Printf("Bye there, i=%d!\n", d)

	fmt.Fprint(w, "hi")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":5555", nil)
}

func main2() {

	/* simulating 10000 users requests  */
	i := 1;
	for ;i < 10000; i++ {
		go getHello(i)
	}

}

func getHello(d int){

	log.Printf("tid=%d, m=getHello, i=%d,begin", id.Id(), d)
	log.Printf("tid=%d, m=getHello, i=%d, processing the hello message", id.Id(), d)
	log.Printf("tid=%d, m=getHello, i=%d, end", id.Id(), d)

}