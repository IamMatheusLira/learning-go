package main

import "fmt"

func sum(x, y int) int {
  return x + y
}

func main() {
  if (2 + 2) > 5 {
      fmt.Println("2 + 2 é maior que 5")
  } else if ((2 + 2) == 5) {
      fmt.Println("2 + 2 é igual a 5")
  } else {
    fmt.Println("2 + 2 é menor que 5")
  }


  if x := sum(3, 2); x < 10 {
    fmt.Println("3 + 2 é menor que 10")
  }
}
