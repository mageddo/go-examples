package main

import "log"

func main() {

	type Payment struct {
		creditor string
		debtor string
		value float64
	}


	payments := &[]Payment{
		{creditor:"Bruna Lopes", debtor:"Elvis", value: 20.99},
		{creditor: "Ana Carolina", debtor: "Elvis", value: 60.00},
		{creditor: "Maria Azeli", debtor: "Elvis", value: 10.55},
	}


	log.Println(payments)

}
