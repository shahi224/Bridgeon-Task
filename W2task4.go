package main
import "fmt"

func reverseArr(arr []int) {
	left, right := 0, len(arr)-1

	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}


func main() {
	arr := []int {1,2,3,4,5}
	fmt.Println("original array:",arr)

	reverseArr(arr)
	fmt.Println("Reversed arr:",arr)
}