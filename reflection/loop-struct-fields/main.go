package main

import (
"fmt"
"reflect"
)

type Person struct {
Age *int
}

func main() {

x := Person{}

v := reflect.ValueOf(x)

values := make([]interface{}, v.NumField())

for i := 0; i < v.NumField(); i++ {
values[i] = v.Field(i).Interface()
}

fmt.Println(values)
}

