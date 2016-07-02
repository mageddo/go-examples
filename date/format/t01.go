package main

import (
	"time"
	"fmt"
)

func main() {

	t := time.Now()
	fmt.Println(t.String())
	fmt.Println(t.Format("2006-01-02 15:04:05"))

}
