package main
import "fmt"
func main(){
	var num1, num2 float64
	var operator string


	fmt.Print("Enter the first number:")
	fmt.Scan(&num1)

	fmt.Print("Enter the operator (+,-,*,/):")
	fmt.Scan(&operator)

	fmt.Print("Enter the second number:")
	fmt.Scan(&num2)


	switch operator {
		case "+" :
			fmt.Printf("Result : % . 2f /n",num1 + num2)
		case "-" :
			fmt.Printf("Result : % . 2f /n",num1 - num2)
		case "*" :
			fmt.Printf("Result : % . 2f /n",num1 * num2)
		case "/" :
			if num2 == 0{
				fmt.Println("Error division by 0 is not allowed")
			}else{
				fmt.Printf("Result : % . 2f /n",num1 / num2)					 

			}
	}
}