package main
import "fmt"


type Node struct {
	data int
	next *Node
}

type Queue struct {
	front *Node
	rear *Node
}

func (q *Queue) Enqueue(value int) {
	newNode := &Node{data: value}
	if q.rear == nil { 
		q.front, q.rear = newNode, newNode
	} else {
		q.rear.next = newNode
		q.rear = newNode
	}
	fmt.Println("Enqueued:", value)
}

func (q *Queue) Dequeue() int {
	if q.front == nil {
		fmt.Println("Queue is empty")
		return -1
	}
	value := q.front.data
	q.front = q.front.next
	if q.front == nil {
		q.rear = nil
	} 
	fmt.Println("Dequeued:", value)
	return value
}

func (q *Queue) Display() {
	if  q.front == nil {
		fmt.Println("Queue is empty")
		return
	}
	fmt.Print("Queue: ")
	for temp := q.front; temp != nil; temp = temp.next {
		fmt.Println(temp.data," ")
	}
	fmt.Println()
}

func main() {
	queue := Queue{}

	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)
	queue.Display()

	queue.Dequeue()
	queue.Display()
	
    
}