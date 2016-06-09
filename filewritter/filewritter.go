package main

import (
	"fmt"
	"time"
	"os"
)


func main(){

	start := time.Now()

	fmt.Println("started")
	file, err := os.Create("/tmp/dat2")
	check(err)

	fmt.Printf("writing to %s\n", file.Name())
	var i int;

	for i = 100000; i > 0; i-- {
		str := fmt.Sprintf("linha de teste: %d\n", makeTimestamp())
		file.WriteString(str)
	}
	file.Sync()
	//os.Remove(file.Name())
	fmt.Printf("file writed in %.2f seconds \n", (start.Sub(time.Now()).Seconds()))

       time.Sleep(3000 * time.Hour)
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}


func check(e error) {
	if e != nil {
		panic(e)
	}
}