package main

import "log"


// see more at http://stackoverflow.com/questions/38666433/how-to-work-with-concurrent-logs-at-golang
func main() {

	/* simulating 10000 users requests  */
	i := 1;
	for ;i < 10000; i++ {
		go getHello(i)
	}

}

func getHello(d int){
	log.Printf("m=getHello, i=%d,begin", d)
	log.Println("m=getHello, processing the hello message")
	log.Printf("m=getHello, i=%d, end", d)
}