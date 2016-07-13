package main

import "fmt"

type Person interface {
	IncreaseAge() int8
}

type OldPerson struct {
	age int8
}

func (p *OldPerson) IncreaseAge() int8 {
	p.age++
	fmt.Printf("p=%p, *age=%p, age=%d\n", &p, &p.age, p.age)
	return p.age
}

func IncreaseAndPrint(p Person){
	p.IncreaseAge()
}

func main() {

	var p = &OldPerson{age:10}
	p.IncreaseAge()
	p.IncreaseAge()
	p.IncreaseAge()
	p.IncreaseAge()
	p.IncreaseAge()
	p.IncreaseAge()
	IncreaseAndPrint(p)
	IncreaseAndPrint(p)
	IncreaseAndPrint(p)
	IncreaseAndPrint(p)
	IncreaseAndPrint(p)
	IncreaseAndPrint(p)
	IncreaseAndPrint(p)

}
