package main

import ("fmt"
"net/http"
"io"
)

func main(){
    resp, err := http.Get("https://reqres.in/api/products")
    if err != nil {
        fmt.Println(err)
    }
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
    }

    resp.Body.Close()

    fmt.Println(string(body))

}
