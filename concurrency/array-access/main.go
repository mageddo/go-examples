package main

import (
	"math/rand"
	"log"
	"fmt"
)

type DB struct {

	ids []int

}

func (db *DB) add(id int){
	db.ids = append(db.ids, id)
}

func (db *DB) remove(id int){
	var i int = 0
	var j int = -1
	for ; j < len(db.ids); j++ {
		if db.ids[i] == id {
			j = i
			break
		}
	}
	if j != -1 {
		db.ids = append(db.ids[:i], db.ids[i+1:]...)
	}
}

func main() {

	var db DB
	for i:=1; i <= 100; i++ {
		id := i
		go db.add(id)
		k := rand.Intn(5)
		if k == 3 {
			log.Println("removing: ", id)
			go db.remove(id)
		}
	}

	var str string
	fmt.Scanln(&str)
	fmt.Println(db.ids)

}
