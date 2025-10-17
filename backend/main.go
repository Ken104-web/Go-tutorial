package main

import (
        "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"

    _ "github.com/lib/pq" 
)

const (
    host = "localhost"
    port = "1234"
    user = "Kenneth"
    password = "ken104"
    dbname = "mydb"
)

var db *sql.DB

type User struct {
    ID int `json:"id"`
    Name string `json:"name"`
    Email string `json:"email"`
}

func main() {
        pgConnStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
        conn, err := sql.Open("postgres", pgConnStr)
        if err != nil {
            log.Fatalf("Err opening database connection: %v", err)
        }
        db = conn
        defer db.Close()

        err = db.Ping()
        if err != nil{
            log.Fatalf("Err opening database connection: %v", err)
        }
        fmt.Println("Connected to postgres database")
        
    }

