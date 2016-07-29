package main

import (
	"syscall"
	"fmt"
)

func main(){
	path, found := syscall.Getenv("PATH")
	fmt.Printf("path=%s, found=%t", path, found)
}
