package main

import (
  "fmt"
  "time"
)

func pinger(c chan<- string) {  // escreve no channel
  for i := 0; i <= 10; i++ {
    c <- fmt.Sprintf("ping %d", i)
  }
  fmt.Println("pinger: all pings sended!")
}

func ponger(c chan string) { // bidirecional, pode ler ou escrever no channel
  for i := 0; i <= 10; i++ {
    c <- fmt.Sprintf("pong %d", i)
    time.Sleep(time.Second * 1)
  }
  fmt.Println("ponger: all pings sended!")

}

func printer(c <-chan string) { // le o channel
  for {
    msg := <- c
    fmt.Println(msg)
    time.Sleep(time.Second * 1)
  }
  fmt.Println("all messages read") // nunca virá aqui
}


func main() {
  /* Esse valor 10 é o buffer, ou seja, quando um valor maior que 1 é passado
   * o channel funciona de forma asincrona, uma rotina consegue rodar n vezes
   * e só irá parar quando o buffer encher, quando está em 1 ela só consegue 
   * enviar 1, e o leitor só consegue ler 1
   *
   */
  var c chan string = make(chan string, 10) 

  go pinger(c)
  go ponger(c)
  go printer(c)

  var input string
  fmt.Scanln(&input)
}
