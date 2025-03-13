package main

import "fmt"

func findMax(arr []int) int {
    max := arr[0]
    for _, num := range arr {
        if num > max {
            max = num
        }
    }
    return max
}

func main() {
    numbers := []int{3, 7, 2, 9, 5}
    fmt.Println("Highest number:", findMax(numbers))
}
