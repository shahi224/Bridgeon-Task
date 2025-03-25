package main

import (
	"fmt"
	"errors"
)

func calculate(a, b float64) (float64, float64, float64, float64, error) {
	sum := a + b
	diff := a - b
	prod := a * b

	if b == 0 {
		return sum, diff, prod, 0, errors.New("division by zero is not allowed")
	}

	quotient := a / b
	return sum, diff, prod, quotient, nil
}

func main() {
	a, b := 10.0, 0.0
	sum, diff, prod, quotient, err := calculate(a, b)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Sum:", sum)
		fmt.Println("Difference:", diff)
		fmt.Println("Product:", prod)
		fmt.Println("Quotient:", quotient)
	}
}
