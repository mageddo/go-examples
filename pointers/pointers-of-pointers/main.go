package main

import "fmt"


type Person struct {
	age int8
}

func IncreaseAge(p **Person) int8 {
	(*p).age++
	fmt.Printf("pOfp=%p, p=%p, *age=%p, age=%d\n", &p, p, &(*p).age, (*p).age)
	return (*p).age
}

func main() {
	var p = &Person{age:10}
	fmt.Printf("p=%p\n", &p)
	IncreaseAge(&p)
	IncreaseAge(&p)
	fmt.Printf("p=%p\n", &p)
}
