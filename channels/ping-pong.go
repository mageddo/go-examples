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
}


func main() {
  var c chan string = make(chan string, 10)

  go pinger(c)
  go ponger(c)
  go printer(c)

  var input string
  fmt.Scanln(&input)
}
