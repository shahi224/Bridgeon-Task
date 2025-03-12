package main

import (
	"fmt"
)
func removeDuplicates(nums []int) []int {
	seen := make(map[int]bool)
	unique := []int{}

	for _, num := range nums {
		if !seen[num] {
			seen[num] = true
			unique = append(unique, num)
		}
	}

	return unique
}

func main() {
	numbers := []int{1, 2, 2, 3, 4, 4, 5, 6, 6, 7}
	fmt.Println("Original slice:", numbers)
	fmt.Println("Without duplicates:", removeDuplicates(numbers))
}