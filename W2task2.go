package main
import "fmt"


type person struct {
	Name string
	Age int
}

func main(){

	p := &person {"shahid", 23}
	fmt.Println("Befor:", p.Name, p.Age)

	p.Age = 24
	fmt.Println("After:",p.Name, p.Age)
}