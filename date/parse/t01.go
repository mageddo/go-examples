package main

import (
	"time"
	"fmt"
)

func main() {

	t := time.Now()
	fmt.Println(t.String())
	fmt.Println(time.Parse("2006-01-02 15:04:05", "2016-03-05 17:10:40")) // work
	fmt.Println(time.Parse("02/01/2006", "03/07/2015")) // work
	fmt.Println(time.Parse("2006/01/02", "2016/07/19")) // work

}
