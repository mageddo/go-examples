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

	c := make(chan *Payment, 10)
	go PaymentQueuePoolSender(c, payments)
	go PaymentQueueConsumer(c)

	dataBasePopulator(payments)

	var str string;
	fmt.Scanln(&str)

}

func dataBasePopulator(payments *[]Payment){
	for i := 1; ; i++ {
		time.Sleep(time.Second * 20)
		*payments = append((*payments), Payment{
			debtor: fmt.Sprintf("debitor: %d", i),
			creditor: fmt.Sprintf("debitor: %d", i),
			value: 1.99,
		})

	}
}

/*
 * Send payments to consumers one-by-one
 */
func PaymentQueuePoolSender(c chan<- *Payment, payments *[]Payment) {
	for {
		for i := 0; i < len(*payments); i++ {
			p := &(*payments)[i]
			if(p.status == 0){
				c <- p
			}
			time.Sleep(time.Millisecond * 500)
		}
	}
	log.Println("pinger: all payments sended!")
}

/*
 * Consumes the queue pool
 */
func PaymentQueueConsumer(c <-chan *Payment){
	for {
		var payment *Payment = <- c
		payment.status = 1;
		log.Printf("paying %.2f from %s to %s\n", payment.value, payment.debtor, payment.creditor)
		payment.status = 2;
	}
}
