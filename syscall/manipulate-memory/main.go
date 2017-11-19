package main

import (
	"fmt"
	"unsafe"
	"reflect"
)

/*
 * The program only can change values of addresses inside the programm
 * to change values of another program uses syscals
 */
func main(){

	ConversionBetweenStructAndPointer()
	NumberToPointerAndBack()
	//ReadMemoryAddressToInt()
	NumberToMemoryAddressThenConvertItBack()
	ReadStructSpecificField()

}

func ReadStructSpecificField() {

	type Person struct {
		name string
		age int
	}

	person := &Person{name:"Elvis", age:22}
	personPointer := unsafe.Pointer(person)
	ageOffset := unsafe.Offsetof(person.age)

	fmt.Printf("person=%v, personPointer=%x, ageOfsset=%x\n", person, personPointer, ageOffset)

	age := (*int)(unsafe.Pointer(uintptr(personPointer) + ageOffset))

	fmt.Printf("age=%d\n", *age)

}
func NumberToMemoryAddressThenConvertItBack() {
	var age int = 22;
	agePtr := uintptr(unsafe.Pointer(&age));

	agePointer := unsafe.Pointer(agePtr)
	convertedAge := (*int)(agePointer)

	fmt.Printf("originalAge=%d, agePtr=%x, agePointer=%x, convertedAge=%d\n", age, agePtr, agePointer, *convertedAge)
}
func ReadMemoryAddressToInt() {
	xz := uintptr(0xC08200A2E0)
	fmt.Printf("sizeof=%d\n", unsafe.Sizeof(xz))
	//ab := (*[unsafe.Sizeof(xz)]byte)(unsafe.Pointer(xz))
	ab := (*int8)(unsafe.Pointer(xz))
	fmt.Printf("p=%p, msg=%d\n", ab, *ab)
}

func NumberToPointerAndBack() {
	j := 11464646544484
	z := (*int)(unsafe.Pointer(reflect.ValueOf(&j).Pointer()))
	fmt.Printf("z=%d\n", *z)
}

func ConversionBetweenStructAndPointer(){


	type T struct {
		A uint32
		B int16
	}

	t1 := T{123, -321}
	fmt.Printf("originalt=%v\n", t1)

	// converting struct to pointer
	p := unsafe.Pointer(&t1)


	// pointer back to function
	v := (*(*T)(p))
	fmt.Printf("convertedt=%v\n", v)

}
