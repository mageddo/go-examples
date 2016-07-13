package main

import "fmt"

type Person struct {
	name string
	age int
	car Car
}

type Car struct {
	model string
}

func (m Person) StayTheSame() {
	fmt.Printf("by value, person=%p, car=%p\n", &m, &m.car)
	m.age = 5
	m.name = "Elvis"
	m.car.model = "Corvette"
}

func (m *Person) Mutate() {
	fmt.Printf("by reference, person=%p, car=%p\n", &m, &m.car)
	m.age = 5
	m.name = "Elvis"
	m.car.model = "Ford Mustang"
}

func main() {
	m := &Person{"Bob", 12, Car{"Corcel"}}
	fmt.Println(m)
	fmt.Println("============================")

	m.StayTheSame()
	fmt.Println("============================")
	m.StayTheSame()
	fmt.Println(m)

	fmt.Println("\n")

	m.Mutate()
	fmt.Println("============================")
	m.Mutate()
	fmt.Println(m)

}