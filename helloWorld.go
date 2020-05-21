package main

import "fmt"
import "time"

func main() {
  fmt.Println("Hello, world!")
  duration := time.Duration(30)*time.Second
  time.Sleep(duration)
}
