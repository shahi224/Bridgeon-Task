package main
import "fmt"

type Node struct {
	data int
	next *Node
}

func ArrayToLinkedList(arr []int) *Node {
	if len(arr) == 0 {
		return nil 
	}

	head := &Node{data: arr[0]} 
	current := head


	for _, value := range arr[1:] {
		newNode := &Node{data: value}
		current.next = newNode 
		current = newNode    
	}

	return head
}

func PrintLinkedList(head *Node) {
	current := head
	for current != nil {
		fmt.Print(current.data, " -> ")
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	arr := []int{10, 20, 30, 40, 50,60} 
	head := ArrayToLinkedList(arr) 

	fmt.Println("Linked List:")
	PrintLinkedList(head) 
}
