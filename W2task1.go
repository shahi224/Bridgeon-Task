package main
import "fmt"

func main() {
	type person struct {
		Name string
		Age int
	}

	p :=new(person)
	fmt.Println(p)
	p.Name = "shahid"
	p.Age = 23
	fmt.Println(*p)


	slice := make([]int, 3)
    slice[0], slice[1], slice[2] = 1, 2, 3
    fmt.Println("Slice using make:", slice)
    
}