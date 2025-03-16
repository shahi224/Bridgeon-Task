package main
import "fmt"

type Node struct {
	data int
	prev *Node
	next *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func (dll *DoublyLinkedList) InsertFront(value int) {
	newNode := &Node{data: value}

	if dll.head == nil { 
		dll.head = newNode
		dll.tail = newNode
	} else {
		newNode.next = dll.head
		dll.head.prev = newNode
		dll.head = newNode
	}
}

func (dll *DoublyLinkedList) DeleteLast() {
	if dll.tail == nil { 
		fmt.Println("List is empty! Nothing to delete.")
		return
	}

	if dll.head == dll.tail { 
		dll.head = nil
		dll.tail = nil
	} else {
		dll.tail = dll.tail.prev 
		dll.tail.next = nil      
	}
}

func (dll *DoublyLinkedList) PrintList() {
	current := dll.head
	fmt.Print("Doubly Linked List: ")
	for current != nil {
		fmt.Print(current.data, " <-> ")
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	dll := &DoublyLinkedList{} 

	dll.InsertFront(10)
	dll.InsertFront(20)
	dll.InsertFront(30)

	fmt.Println("After inserting elements at front:")
	dll.PrintList()

	dll.DeleteLast()
	fmt.Println("After deleting last element:")
	dll.PrintList()
}
