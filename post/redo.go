package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
    "io"
)


type Sitemapindex struct {
  Locations []string `xml:"url>loc"`
}

type News struct {
	Titles    []string `xml:"url>news:news>news:title"`
	Keywords  []string `xml:"url>news:news>news:keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct{
    Keyword string
    Location string
}


func main(){

    var s Sitemapindex
    var n News
	resp, _ := http.Get("https://www.washingtonpost.com/news-world-sitemap.xml")
    bytes, _ := io.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)
    news_map := make(map[string]NewsMap)

    for _, Location := range s.Locations{
        resp, _ := http.Get(Location)
        bytes, _ := io.ReadAll(resp.Body)
        xml.Unmarshal(bytes, &n)
		 // fmt.Println(s.Locations)
    
	
		fmt.Println("Found Keywords:", len(n.Keywords))
		
        for index, _ := range n.Keywords{
            news_map[n.Titles[index]] = NewsMap{n.Keywords[index], n.Locations[index]}
        }
    }
    for index, data := range news_map {
        fmt.Println("\n\n\n\n\n",index)
        fmt.Println("\n",data.Keyword)
        fmt.Println("\n",data.Location)
    }
}
