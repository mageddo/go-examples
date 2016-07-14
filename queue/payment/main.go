package main

import (
	"log"
	"time"
	"fmt"
	"sync"
	"container/list"
)

type Payment struct {
	creditor string
	debtor string
	value float64
	status int8 // 0=not processed, 1=processing, 2=processed
}

var keep int = 1; // 1=keep populate the database, !1=do not populate the database

type DB struct {
	mu sync.Mutex
	payments list.List
}

var db DB

func main() {

	db.payments.PushBack(&Payment{creditor:"Bruna Lopes", debtor:"Elvis", value: 20.99})
	db.payments.PushBack(&Payment{creditor: "Ana Carolina", debtor: "Elvis", value: 60.00})
	db.payments.PushBack(&Payment{creditor: "Maria Azeli", debtor: "Elvis", value: 10.55})

	c := make(chan *Payment, 1)

	// one guy reading from database and putting on the queue pool
	go PaymentQueuePoolSender(c)

	for i:=1; i <= 10; i++ { // 10 consumers of the queue
		go PaymentQueueConsumer(c, i)
	}

	// ever populating the database with new requests
	go dataBasePopulator()

	// print the database stats time to times
	go stats()

	// some queue useful commands
	for cmd:="" ; ; fmt.Scanln(&cmd) {
		switch cmd {
		case "1":
			keep = 1
			log.Println("populate activated!")
			break
		case "0":
			keep = 0
			log.Println("populate disabled!")
			break
		case "2":
			for e := db.payments.Front(); e != nil; e = e.Next() {
				p := e.Value.(*Payment)
				log.Print(p)
			}
			log.Println()
			log.Println("database prited!")
		}
	}

}

/**
 * Keep inserting itens on database(array) to be consumed by queues
 */
func dataBasePopulator(){

	for i := 1; ; i++ {
		time.Sleep(time.Millisecond * 100)
		if keep != 1 {
			continue
		}
		db.mu.Lock()
		db.payments.PushBack(&Payment{
			debtor: fmt.Sprintf("debitor: %d", i),
			creditor: fmt.Sprintf("credtor: %d", i),
			value: 1.99,
		})
		db.mu.Unlock()

	}
}

/*
 * Send payments to consumers one-by-one
 */
func PaymentQueuePoolSender(c chan<- *Payment) {
	for {
		db.mu.Lock()
		for e := db.payments.Front(); e != nil; e = e.Next() {
			p := e.Value.(*Payment)
			if(p.status == 0){
				c <- p
				log.Println("sending to pay")
				time.Sleep(time.Millisecond * 500)
			}
		}
		db.mu.Unlock()
	}
	log.Println("pinger: all payments sended!")
}

/*
 * Consumes the queue pool
 */
func PaymentQueueConsumer(c <-chan *Payment, i int){
	for {
		var payment *Payment = <- c
		payment.status = 1;
		log.Printf("queue=pay-%d, paying %.2f from %s to %s\n", i, payment.value, payment.debtor, payment.creditor)
		// taking a time to execute the very long payment process
		time.Sleep(time.Second * 10)
		log.Printf("queue=pay-%d, payed!", i)
		payment.status = 2;
	}
}


/**
 * All items on database that are not processed yet
 */
func countNotProcessedPayments() int {
	var count int = 0;

	for e := db.payments.Front(); e != nil; e = e.Next() {
		p := e.Value.(*Payment)
		if p.status == 0 {
			count++;
		}
	}
	return count
}

/**
 * All items on database that are processing now
 */
func countProcessingPayments() int {
	var count int = 0;
	for e := db.payments.Front(); e != nil; e = e.Next() {
		p := e.Value.(*Payment)
		if p.status == 1 {
			count++;
		}
	}
	return count
}

func countProcessedPayments() int {
	var count int = 0;
	for e := db.payments.Front(); e != nil; e = e.Next() {
		p := e.Value.(*Payment)
		if p.status == 2 {
			count++;
		}
	}
	return count
}

func stats(){

	for ;; {
		log.Println("===========================")
		log.Printf("not processed=%d, processing=%d, processed=%d\n", countNotProcessedPayments(),
			countProcessingPayments(), countProcessedPayments())
		log.Println("===========================")
		time.Sleep(time.Second * 5)
	}
}