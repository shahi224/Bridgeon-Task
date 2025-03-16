package main

import "fmt"

type Node struct {
    data int
    next *Node
}

type SinglyLinkedList struct {
    head *Node
}

func (sll *SinglyLinkedList) Insert(value int) {
    newNode := &Node{data: value}
    if sll.head == nil {
        sll.head = newNode
        return
    }
    current := sll.head
    for current.next != nil {
        current = current.next
    }
    current.next = newNode
}

func (sll *SinglyLinkedList) Delete(value int) {
    if sll.head == nil {
        return
    }
    if sll.head.data == value {
        sll.head = sll.head.next
        return
    }
    current := sll.head
    for current.next != nil && current.next.data != value {
        current = current.next
    }
    if current.next != nil {
        current.next = current.next.next
    }
}

func (sll *SinglyLinkedList) Search(value int) bool {
    current := sll.head
    for current != nil {
        if current.data == value {
            return true
        }
        current = current.next
    }
    return false
}

func (sll *SinglyLinkedList) Display() {
    current := sll.head
    for current != nil {
        fmt.Printf("%d -> ", current.data)
        current = current.next
    }
    fmt.Println("nil")
}

func main() {
    sll := &SinglyLinkedList{}

    sll.Insert(10)
    sll.Insert(20)
    sll.Insert(30)
	sll.Insert(40)


    fmt.Println("Linked List:")
    sll.Display()

    fmt.Println("Search 20:", sll.Search(20))
    fmt.Println("Search 40:", sll.Search(40)) 

    
    sll.Delete(20)
    fmt.Println("After deleting 20:")
    sll.Display()
}
