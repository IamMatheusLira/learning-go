package main

import (
	"example/hello/mypackage"
	"example/hello/typecast"
	"fmt"
)

func main() {
  fmt.Println("Hello World") 
  mypackage.SayGreetings("Matheus")

  floatNum := typecast.Cast(123)
  fmt.Printf("Num float = %f", floatNum)
}
