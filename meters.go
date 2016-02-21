package main

import "fmt"

func main() {
  fmt.Print("Feet: ")
  var input float64
  fmt.Scanf("%f", &input)

  fmt.Println(input * .3048)
}
