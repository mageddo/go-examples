package main

import "fmt"

type Mutatable struct {
	a int
	b int
}

func (m Mutatable) StayTheSame() {
	m.a = 5
	m.b = 7
	fmt.Printf("by value, p=%p, value=%+v\n", &m, m)
}

func (m *Mutatable) Mutate() {
	fmt.Printf("by pointer, p=%p, value=%+v\n", &m, m)
	m.a = 5
	m.b = 7
}

func main() {
	m := &Mutatable{10, 10}
	fmt.Printf("after create the struct: m=%p, %+v\n\n", m, m)
	m.StayTheSame()
	fmt.Printf("after call #StayTheSame: m=%p, %+v\n\n", m, m)
	m.Mutate()
	fmt.Printf("after call #Mutate: m=%p, %+v\n\n", m, m)

	fmt.Printf("a=%d, b=%d, a=%d\n", m.a, m.b, (*m).a)

}