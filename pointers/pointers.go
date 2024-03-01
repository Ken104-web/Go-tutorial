package main

import "fmt"

func main(){
    x := 15

    a := &x
    fmt.Println(a)
    fmt.Println(*a)// memory address
    
    *a = 5          // sets the value pointed at to 5, which means x is modified (since x is stored at the mem addr)
	fmt.Println(x)  // see the new value of x
	*a = *a**a      // what is the value of x now? 
	fmt.Println(x)  // prints a value
	fmt.Println(*a) // prints a value

	fmt.Println(&x) // prints a memory address
	fmt.Println(a) // prints a memory address
}

