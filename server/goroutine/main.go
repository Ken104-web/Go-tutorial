// Goroutine are functions or methods that run concurrently(same time) with other functions and methods
// they communicate using channels
// When a new Goroutine is started, the goroutine call returns immediately
// The main Goroutine should be running for any other Goroutines to run
package main

import (
	"fmt"
	"time"
)

func hello(){
    fmt.Println("Hello from here")
    
}
func main(){
    go hello()
    time.Sleep(1 * time.Second)
    fmt.Println("Here I am")
}

