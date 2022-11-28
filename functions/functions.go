package main

import "fmt"

func simpleFunction() {
	fmt.Println("Hello World")
}

func multipleReturns(x, y int) (int, int) {
	return x, y
}

func namedReturns(num1, num2 int) (res int) {
	res = num1 + num2
	return
}

func main() {
	simpleFunction()
	x, y := multipleReturns(1, 2)
	sum := namedReturns(5, 5)
	fmt.Println(x, y)
	fmt.Println(sum)
}
