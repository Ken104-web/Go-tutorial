package main

import "fmt"

func foo(){
    defer fmt.Println("Done!")
    defer fmt.Println("Are you done yet?")
    fmt.Println("Done some stuff, who knows what?")

    // for i:=0; i<5; i++{
    //     defer fmt.Println(i)
    // }
}

func main(){
    foo()
}
