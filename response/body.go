package main

import ("fmt"
"encoding/xml"
"net/http" //  encapsulates the request-response pattern in one method
"io" // access the Body property of a response 
)


type Sitemapindex struct {
    Locations []Location `xml: "sitemap"`
}

type Location struct {
    Loc string `xml:"loc"`
} 

func main(){
    resp, err := http.Get("https://www.samsung.com/sitemap.xml")
    if err != nil {
        fmt.Println("Error making http request:", err)
    }
    bytes, err := io.ReadAll(resp.Body)
    if err != nil{
        fmt.Println(err)
    }
    var s Sitemapindex
    xml.Unmarshal(bytes, &s)

    resp.Body.Close()

    fmt.Println(s.Locations)

}
