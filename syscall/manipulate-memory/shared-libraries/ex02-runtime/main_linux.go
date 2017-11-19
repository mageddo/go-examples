package main

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
import "C"

import (
	"fmt"
	"reflect"
)

func main() {
	//
	//var x func(str *string) int

	handle := C.dlopen(C.CString("lib/libhello.so"), C.RTLD_LAZY)

	p := C.dlsym(handle, C.CString("sum"))
	fmt.Printf("bar is at %v, type=%v\n", p, reflect.TypeOf(p))



	fn := *(*func()int)(p)
	//name := "x"
	ct := fn();
	println(ct)

	//fn := reflect.ValueOf(&x)
	//fn.SetInt(p.)
	//
	//name := "Elvis"
	//
	//syscall.dlo

	//fmt.Println(syscall.Syscall(uintptr(p), 0, 0, 0))
}
