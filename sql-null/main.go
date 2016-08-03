package main

import (
	"bitbucket.org/mageddo/web-kmymoney/kmymoney-api/crud"
	"fmt"
)

type Person struct {

	Name string
	Age int8

}

func main() {

	var person Person
	p := crud.GetPool()

	row := p.QueryRowx("SELECT 'hello' AS Name, 12 AS Age")

	var name *string = &person.Name;
	var age *int8 = &person.Age

	row.Scan(&name, &age)

	fmt.Printf("person=%v, name=%s, age=%d", person, "", *age)

}
