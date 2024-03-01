package main

import "fmt"

// Every variable and parameter needs to have a type defined with it. x float64, y float64 means our x and y parameters will be float64
func add(x,y float64) float64 {
// The float64 outside of the parenthesis and before the curly brace is the type for the return        
    return x+y
}

func main() {
        var num1 float64 = 5.6
        var num2 float64 = 9.5

    fmt.Print(add(num1, num2))
}

