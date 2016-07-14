package main

import (
	"log"
	"fmt"
	"sync"
	"time"
	"math/rand"
)

// based on http://stackoverflow.com/a/10735763/2979435

type Counter struct {
	mu  sync.Mutex
	x   int64
}

func (c *Counter) Add(x int64) {
	c.mu.Lock()
	c.x += x
	c.mu.Unlock()
}

func (c *Counter) Value() (x int64) {
	c.mu.Lock()
	x = c.x
	c.mu.Unlock()
	return
}
func (c *Counter) IncreaseAndReturn(sync bool) (x int64) {
	if sync {
		c.mu.Lock()
	}
	xbefore := c.x
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(30)))
	c.x = xbefore + 1
	x = c.x
	if sync {
		c.mu.Unlock()
	}

	return
}

func main() {
	var views Counter

	for i:=0; i < 100; i++ {
		go func(){
			log.Printf("Counting %d so far.\n", views.IncreaseAndReturn(false))
		}()
	}
	var str string;
	fmt.Scanln(&str)
	fmt.Println("in the end: ", views.Value())

}
