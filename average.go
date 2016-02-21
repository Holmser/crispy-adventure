package main

import "fmt"

func main() {
  var x [5]float64
  x[0] = 50
  x[1] = 78
  x[2] = 94
  x[3] = 89
  x[4] = 100
  //fmt.Println(x)

  var total float64 = 0
  for i := 0; i< len(x); i++ {
    total += x[i]
  }
  fmt.Println(total / float64(len(x)))
}
