package main
import "fmt"

func insertAt(arr []int, index, value int)[]int {
	arr = append(arr[:index], append([]int {value}, arr[index:]...)...)
	return arr
}

func deleteAt(arr []int, index int)[]int {
	return append(arr[:index],arr[index+1:]...) 
}


func main(){

	arr := []int {1,2,3,4,5}
	fmt.Println("original array:",arr)

	arr = insertAt(arr, 2,77)
	fmt.Println("inserting:",arr)

	arr = deleteAt(arr, 3)
	fmt.Println("deleting:",arr)
}