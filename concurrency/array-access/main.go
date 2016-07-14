package main

import (
	"math/rand"
	"log"
	"fmt"
	"time"
	"sync"
)

type DB struct {
	mu sync.Mutex
	ids []int
}

/**
 * add item to database
 */
func (db *DB) add(id int){
	db.mu.Lock()
	log.Println("adding: ", id)
	db.ids = append(db.ids, id)
	db.mu.Unlock()
}

/**
 * find item on database, return -1 if not exists
 */
func (db *DB) find(id int) int {
	db.mu.Lock()
	log.Println("finding: ", id)
	var j int = -1
	for i :=0; i < len(db.ids); i++ {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(30)))
		if db.ids[i] == id {
			j = i
			break
		}
	}
	db.mu.Unlock()
	return j
}

/**
 * remove item from database
 */
func (db *DB) remove(id int){
	j := db.find(id)
	db.mu.Lock()
	if j != -1 {
		log.Println("removing: ", id)
		db.ids = append(db.ids[:j], db.ids[j + 1:]...)
	}
	db.mu.Unlock()
}

func main() {

	var db DB
	n := 100
	from := 1

	for i:= from; i <= n; i++ {
		id := i*100
		db.add(id)
	}

	for i:= from; i <= n; i++ {
		go db.find(i)
		go db.remove(i)
	}
	fmt.Scanln()
	fmt.Println(db.ids)

}
