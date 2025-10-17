package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)



func main(){
    url := "https://restcountries.com/v3.1/name/kenya"

    if resp, err := http.Get(url); err == nil {
        defer resp.Body.Close()
        if body, err := io.ReadAll(resp.Body);
            err == nil{
            fmt.Println(string(body))
            } else {
                log.Fatal(err)
            } 
            fmt.Println("Done")
    }

}
