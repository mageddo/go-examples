package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Age *int
	//Name *string
}

func main() {

	x := Person{}
	PrintValues(x)
	AllocStruct(&x)

}

func PrintValues(i interface{}){
	v := reflect.ValueOf(i)
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		fmt.Println(f.Interface())
	}
}

func AllocStruct(_ interface{}){

	var pe *Person = new(Person)

	v := reflect.ValueOf(pe)
	//for i := 0; i < v.Elem(); i++ {
		f := v.Elem().FieldByName("Age")
		//f := v.Field(i)

		fmt.Println("can set", f.CanSet(), f.CanAddr())
		//p:= unsafe.Pointer(reflect.New(f.Type()).Pointer())

	//}
}
