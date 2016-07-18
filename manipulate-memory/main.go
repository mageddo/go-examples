package main

import (
	"fmt"
	"unsafe"
	"reflect"
)

func main(){

	j := 11464646544484
	x := "ok"
	var k *string = &x

	fmt.Printf("p=%p, msg=%s\n", k, *k)
	z := (*int)(unsafe.Pointer(reflect.ValueOf(&j).Pointer()))
	fmt.Printf("p=%p, msg=%d\n", k, *z)



	ab := (*string)(unsafe.Pointer(uintptr(0xc82000a330)))

	fmt.Printf("p=%p, msg=%s\n", ab, *ab)



}
