// concurrency is dealing with multiple things at once while not doing multiple things simultaneously

package main

import (
    "fmt"
    "sync"
    "time"
)

var wg sync.WaitGroup


func cleanup(){
    if r := recover(); r != nil {
        fmt.Println("Recovered in cleanup:", r)
        defer wg.Done()
    }
}

func say(s string){
    defer cleanup()
    for i:=0; i<3; i++{
            time.Sleep(100*time.Millisecond)
            fmt.Println(s)
            if i == 2 {
                panic("Oh dear.. a 2")
        
            }
    }
}

func main(){
    wg.Add(1)
    go say("Hey")
    wg.Add(1)
    go say("there")
    wg.Wait()
    
    wg.Add(1)
    go say("ken")
    wg.Wait()
}
