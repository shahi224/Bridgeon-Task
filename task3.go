package main
import "fmt"
func main(){
	var n int
	fmt.Println("Enter number of terms :")
	fmt.Scan(&n)

	first,second := 0,1

	fmt.Println("Fibonacci series:")
	
	for i := 0; i < n; i++ {
		fmt.Print(first," ")

		next := first + second
		first = second
		second = next
	}
}