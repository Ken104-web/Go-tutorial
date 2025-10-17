package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io"
	"net/http"
)

type NewsMap struct {
    Name string
    Location string
}

type NewsAggPage struct {
    Title string
    News map[string]NewsMap
}

type Sitemapindex struct {
    Locations []string `xml:"sitemap>loc"`
}

type News struct {
    Titles []string `xml:"url>news>title"`
    Names []string `xml:"url>news>publication>keywords"`
    Locations []string `xml:"url>loc"`
}


func indexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {

    var s Sitemapindex
    var n News
    resp, _ := http.Get("https://www.washingtonpost.com/news-national-sitemap.xml")
    bytes, _ := io.ReadAll(resp.Body)
    xml.Unmarshal(bytes, &s)
    news_map := make(map[string]NewsMap)
    resp.Body.Close()

    for _, Location := range s.Locations {
        resp, _ := http.Get(Location)
        bytes, _ := io.ReadAll(resp.Body)
        xml.Unmarshal(bytes, &n)

        for idx, _ := range n.Titles {
            news_map[n.Titles[idx]] = NewsMap{n.Names[idx], n.Locations[idx]}
        }
    }

    p := NewsAggPage{Title: "Amazing News Aggregator", News: news_map}

    t, _ := template.ParseFiles("aggregatorfinish.html")
    t.Execute(w, p)
}


func main() {
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/agg/", newsAggHandler)
    http.ListenAndServe(":8000", nil) 
}
