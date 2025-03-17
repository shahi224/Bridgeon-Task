package main
import "fmt"


type Node struct {
	data int
	next *Node
}

type stack struct {
	top *Node
}

func (s *stack) Push(value int) {
	s.top = &Node{data : value, next : s.top}
}

func (s *stack) Pop() int {
	if s.top == nil {
		fmt.Println("Stack is empty")
		return -1
	}
	val := s.top.data
	s.top = s.top.next
	return val
}

func main() {
	stack := &stack{}
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}