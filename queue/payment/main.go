package main

import (
	"log"
	"time"
	"fmt"
	"container/list"
	"math/rand"
)

type Payment struct {
	id int
	creditor string
	debtor string
	value float64
	status int8 // 0=not processed, 1=processing, 2=processed
}

var keep int = 1; // 1=keep populate the database, !1=do not populate the database

type DB struct {
	payments list.List
}

var db DB

func main() {

	db.payments.PushBack(&Payment{id: 1, creditor:"Bruna Lopes", debtor:"Elvis", value: 20.99})
	db.payments.PushBack(&Payment{id: 2, creditor: "Ana Carolina", debtor: "Elvis", value: 60.00})
	db.payments.PushBack(&Payment{id: 3, creditor: "Maria Azeli", debtor: "Elvis", value: 10.55})

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
				fmt.Print(p)
				fmt.Print(", ")
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

	for i := 4; ; i++ {
		time.Sleep(time.Millisecond * 300)
		if keep != 1 {
			continue
		}
		db.payments.PushBack(&Payment{
			id: i,
			debtor: fmt.Sprintf("debitor: %d", i),
			creditor: fmt.Sprintf("credtor: %d", i),
			value: 1.99,
		})

	}
}

/*
 * Send payments to consumers one-by-one
 */
func PaymentQueuePoolSender(c chan<- *Payment) {
	var found bool
	for {
		found = false
		for e := db.payments.Front(); e != nil; e = e.Next() {
			p := e.Value.(*Payment)
			if(p.status == 0){
				log.Printf("sending payment=%d: ", p.id)
				c <- p
				p.status = 1
				log.Printf("sent payment=%d: ", p.id)
				found = true
			}
		}
		if !found {
			log.Println("nothing to process, hibernating...")
			time.Sleep(time.Second * 30)
		}
	}
	log.Println("pinger: all payments sended!")
}

/*
 * Consumes the queue pool
 */
func PaymentQueueConsumer(c <-chan *Payment, i int){
	for {
		var payment *Payment = <- c
		log.Printf("received payment=%d queue=pay-%d, paying %.2f from %s to %s\n", payment.id, i, payment.value, payment.debtor, payment.creditor)
		// taking a time to execute the very long payment process
		time.Sleep(time.Second * time.Duration(rand.Int31n(10)))
		log.Printf("payed payment=%d, queue=pay-%d", i, payment.id)
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