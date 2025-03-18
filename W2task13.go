package main
import "fmt"

type Queue struct {
	items []int
}

func(q *Queue) Enqueue(item int) {
	q.items = append (q.items,item)
}

func (q *Queue) Dequeue()int {
	if len(q.items) == 0 {
		fmt.Println("Queue is empty")
		return -1
	}
	Front := q.items[0]
	q.items = q.items[1:]
	return Front
}

func (q *Queue) Front()int {
	if len(q.items) == 0 {
		fmt.Println("Queue is empty")
		return -1
	}
	return q.items[0]
}

func(q *Queue)isEmpty()bool {
	return len(q.items)==0
}

func main() {
	queue := Queue{}

	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	fmt.Println("Queue:",queue.items)

	fmt.Println("Front element:",queue.Front())

	fmt.Println("Dequeued:",queue.Dequeue())
	fmt.Println("Dequeued:",queue.Dequeue())
	fmt.Println("Dequeued:",queue.Dequeue())

	fmt.Println("After dequeing:",queue.items)

	fmt.Println("Is queue is empty:",queue.isEmpty() )
}