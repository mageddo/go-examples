package main

import (
	"time"
	"fmt"
)

func main() {

	t := time.Now()
	fmt.Println(t.String())
	fmt.Println(t.Format("2006-01-02 15:04:05")) // work: year-month-day
	fmt.Println(t.Format("2015-01-02 15:04:05")) // not work
	fmt.Println(t.Format("02/01/2006 15:04:05")) // work day/month/year
	fmt.Println(t.Format("02/01/2006")) // work day/month/year

}
