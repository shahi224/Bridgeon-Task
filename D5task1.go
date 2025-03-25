package main

import (
	"fmt"
	"sync"
)

func printEvenNumbers(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		fmt.Println("Even:", i)
	}
}

func printOddNumbers(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		fmt.Println("Odd:", i)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2) 

	go printEvenNumbers(&wg)
	go printOddNumbers(&wg)

	wg.Wait()
	fmt.Println("Finished printing numbers")
}
