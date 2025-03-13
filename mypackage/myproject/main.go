package main

import (
    "fmt"
    "myproject/mathutil" // Import the custom package
)

func main() {
    a, b := 10.0, 5.0

    fmt.Println("Addition:", mathutils.Add(a, b))
    fmt.Println("Subtraction:", mathutils.Subtract(a, b))
    fmt.Println("Multiplication:", mathutils.Multiply(a, b))

    result, err := mathutils.Divide(a, b)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Division:", result)
    }
}
