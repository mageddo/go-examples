package main

import (
	"log"
	"time"
	"fmt"
)

type Payment struct {
	creditor string
	debtor string
	value float64
	status int8 // 0=not processed, 1=processing, 2=processed
}

func main() {

	payments := &[]Payment{
		{creditor:"Bruna Lopes", debtor:"Elvis", value: 20.99},
		{creditor: "Ana Carolina", debtor: "Elvis", value: 60.00},
		{creditor: "Maria Azeli", debtor: "Elvis", value: 10.55},
	}

	c := make(chan Payment, 1)
	go PaymentQueueSender(c, payments)
	go PaymentQueueConsumer(c)

	log.Println(payments)

	var str string;
	fmt.Scanln(&str)

}

/*
 * Send payments to consumers one-by-one
 */
func PaymentQueueSender(c chan<- Payment, payments *[]Payment) {
	for i := 0; i <= len(*payments); i++ {
		c <- (*payments)[i]
		time.Sleep(time.Second * 15)
	}
	log.Println("pinger: all payments sended!")
}

func PaymentQueueConsumer(c <-chan Payment){

	for {
		var payment Payment = <- c
		log.Printf("paying %.2f from %s to %s\n", payment.value, payment.debtor, payment.creditor)
	}
}
