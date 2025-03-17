package main
import "fmt"

func main () {

	stack := []int{}

	stack = append(stack,10)
	stack = append(stack,20)
	stack = append(stack,30)

	fmt.Println("Stack after pushing:",stack)


	lastIndex := len(stack) -1
	top := stack[lastIndex]
	stack = stack[:lastIndex]

	fmt.Println("popped element:", top)
	fmt.Println("Stack after popping:", stack)




    
}






































