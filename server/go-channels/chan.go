// Channels can be thought of as pipes using which Goroutines communicate
// sending and receiving from a channel can be implied as such
// read from channel a data := <- a 
// write to channel a a <- data
package main

import(
    "fmt"
)

func hello(done chan bool){
    fmt.Println("hello from the other side")
    done <- true
}
func main (){
    done := make(chan bool)
    go hello(done)
    <- done
    fmt.Println("main func")

}
