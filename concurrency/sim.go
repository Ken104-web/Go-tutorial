package main

import (
	"fmt"
    "sync"
)

var wg sync.WaitGroup

func main() {
	go helloworld()
	goodbye()
}

func helloworld() {
	fmt.Println("Hello World!")
}

func goodbye() {
	fmt.Println("Good Bye!")
}
