package main

import (
    "fmt"
    "net/http"
    "html/template"
)

type NewsAggHandler struct{
    Title string
    News string

}

func newsAggHandler(w http.ResponseWriter, r *http.Request){
    p := NewsAggHandler{Title: "Amazing News Aggregator", News: "some news"}
    t, _ := template.ParseFiles("basictemplating.html")
    t.Execute(w, p)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Go is neat!</h1>")
}


func main(){
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/agg/", newsAggHandler)
    http.ListenAndServe(":8000", nil)
}