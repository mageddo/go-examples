package main

import "fmt"

type Mutatable struct {
	a int
	b int
}

func (m Mutatable) StayTheSame() {
	fmt.Printf("by value=%p\n", &m)
	m.a = 5
	m.b = 7
}

func (m *Mutatable) Mutate() {
	fmt.Printf("by value=%p\n", &m)
	m.a = 5
	m.b = 7
}

func main() {
	m := &Mutatable{0, 0}
	fmt.Println(m)
	m.StayTheSame()
	fmt.Println(m)
	m.Mutate()
	fmt.Println(m)

	fmt.Println(m.a, m.b, (*m).a)

}