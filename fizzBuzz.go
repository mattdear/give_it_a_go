package main

import "fmt"
import "time"
import "strconv"

func main() {

  output := ""

  for i := 1; i < 101; i++ {
    if i % 3 == 0 {output += "Fizz"}
    if i % 5 == 0 {output += "Buzz"}
    if output == "" {output = strconv.Itoa(i)}
    fmt.Println(output)
    output = ""
    duration := time.Second
    time.Sleep(duration)
  }

  duration := time.Duration(10)*time.Second
  time.Sleep(duration)

}
