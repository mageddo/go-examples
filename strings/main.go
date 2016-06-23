package main

import (
	"strings"
	"fmt"
)

func main(){
	arr := strings.Split("a,b,c,d,e", ",")

	for i:=0; i < len(arr); i++ {
		fmt.Printf("linha: %d, valor=%s \n", i, arr[i])
	}

}