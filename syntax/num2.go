// incase we use float32
// if you have a function that returns multiple values, you specify their types inside parenthesis
package main

import "fmt"

func add(x, y float32) float32{
        return x+y
}

fun main(){
    num1, num2 := 5.6, 9.5
  fmt.Print(add(num1, num2))
}
