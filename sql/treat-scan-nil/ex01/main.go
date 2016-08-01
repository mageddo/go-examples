package main

type Person struct {
	Name *string
	Age *int
}

func (p *Person)GetName() string {
	if p.Name == nil {
		return ""
	}
	return *p.Name
}

func (p *Person)GetAge() int {
	if p.Age == nil {
		return 0
	}

	return *p.Age
}

func main() {

	// 1 - create a select querie
	// 2 - scan to struct pointer
	// 3 - manipulate using get/set methods instead access property directly
}