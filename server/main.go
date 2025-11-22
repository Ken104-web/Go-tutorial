package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    // "strconv"
    "sync"
)

type User struct{
    Name string `json: "name"`
}
var userCache = make(map[int]User)
var cacheMutex sync.RWMutex


func main(){
    max := http.NewServeMux()
    max.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
        w.Write([]byte("Hi ken here"))
    })
    fmt.Println("running good")
    // post users
    max.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request){
        fmt.Println("debug")
        var user User
        err := json.NewDecoder(r.Body).Decode(&user)
        fmt.Println("debug")
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        if user.Name == ""{
                http.Error(w, "Name is needed", http.StatusBadRequest)
                return
        }
        cacheMutex.Lock()
        userCache[len(userCache)+1] = user
        cacheMutex.Unlock()
        w.Write([]byte("User Created"))

    })

    fmt.Println("Server running at port 2020")
    http.ListenAndServe(":2020", max)
}

