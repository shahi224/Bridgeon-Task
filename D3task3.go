package main

import (
	"fmt"
)

func main() {
	grades := make(map[string]float64)


	grades["Alice"] = 85.5
	grades["Bob"] = 78.0
	grades["Charlie"] = 92.3

	grades["Bob"] = 80.5

	fmt.Println("Student Grades:")
	for name, grade := range grades {
		fmt.Printf("%s: %.2f\n", name, grade)
	}

	grades["David"] = 88.7
	fmt.Println("\nAfter Adding David:")
	for name, grade := range grades {
		fmt.Printf("%s: %.2f\n", name, grade)
	}


	delete(grades, "Alice")
	fmt.Println("\nAfter Removing Alice:")
	for name, grade := range grades {
		fmt.Printf("%s: %.2f\n", name, grade)
	}
}