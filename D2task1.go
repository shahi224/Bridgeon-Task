package main

import "fmt"

func calculate(a, b float64) (float64, float64, float64, float64) {
	sum := a + b
	diff := a - b
	prod := a * b
	var quotient float64
	if b != 0 {
		quotient = a / b
	} else {
		quotient = 0
	}
	return sum, diff, prod, quotient
}

func main() {
	a, b := 10.0, 5.0
	sum, diff, prod, quotient := calculate(a, b)
	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", diff)
	fmt.Println("Product:", prod)
	fmt.Println("Quotient:", quotient)
}
