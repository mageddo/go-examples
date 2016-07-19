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

	var u1 *int8 = (*int8)(unsafe.Pointer(uintptr(0xC08200A2E0)))

	fmt.Printf("p=%p, pv=%p\n", &u1, u1)
	fmt.Printf("p=%p, msg=%s\n", k, *k)
	z := (*int)(unsafe.Pointer(reflect.ValueOf(&j).Pointer()))
	fmt.Printf("p=%p, msg=%d\n", k, *z)
	//
	//var v int = 1
	//var tmp int
	//for ;;{
	//	fmt.Printf("v=%d, vp=%p\n", v, &v)
	//	fmt.Scanln(&tmp);
	//	if(tmp != -1){
	//		v = tmp
	//	}
	//}

	//xz := uintptr(0xc08200a290)
	//xz := uintptr(0x024670FD)
	xz := uintptr(0xC08200A2E0)
	fmt.Printf("sizeof=%d\n", unsafe.Sizeof(xz))
	//ab := (*[unsafe.Sizeof(xz)]byte)(unsafe.Pointer(xz))
	ab := (*int8)(unsafe.Pointer(xz))
	fmt.Printf("p=%p, msg=%d\n", ab, *ab)
	//*ab = 5

	//ab2 := (*int)(unsafe.Pointer(uintptr(0xc82000a330+3)))
	//fmt.Printf("p=%p, msg=%d\n", ab2, *ab2)



}
