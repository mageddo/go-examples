package main

import (
  "fmt"
  "time"
)

func pinger(c chan string) {
  for i := 0; i < 20; i++ {
    c <- fmt.Sprintf("ping %d", i)
  }
  fmt.Println("pinger: all pings sended!")
}

func printer(c chan string) {
  for {
    msg := <- c
    fmt.Println(msg)
    time.Sleep(time.Second * 1)
  }
}

func main() {
  var c chan string = make(chan string)

  go pinger(c)
  go printer(c)

  var input string
  fmt.Scanln(&input)
}
