package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Page struct {
    Title string
    Content string
    Date string
}

const (
    dbUser = "yoshua"
    dbPass = "password"
    dbHost = "localhost"
    dbPort = "3306"
    dbName = "Go"
)

var db *sql.DB
// var err error

func servePage(w http.ResponseWriter, r *http.Request) {
    var current_page Page
    vars := mux.Vars(r)
    pageID := vars["guid"]

    err := db.QueryRow("SELECT page_title, page_content,page_date FROM pages WHERE page_guid=?", pageID).Scan(&current_page.Title, &current_page.Content, &current_page.Date)

    html := `<html><head><title>` + current_page.Title +
    `</title></head><body><h1>` + current_page.Title + `</h1><div>` +
    current_page.Content + `</div></body></html>`

    if err !=nil {
        w.WriteHeader(404)
        http.Error(w, http.StatusText(404), http.StatusNotFound)
        
    } else {
        fmt.Fprintf(w, html)
    }
    
    
}


func main() {
    var err error
    // Define the data source name (DSN)
    con_string := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	// dsn := "yoshua:password@tcp(127.0.0.1:3306)/Go"
    // Open a connection to the database
    db, err = sql.Open("mysql", con_string)
    if err != nil {
        log.Fatalf("Error opening database: %v", err)
    }else {
        fmt.Println("Connected to MYSQl")
    }
    defer db.Close()

    r := mux.NewRouter()

    r.HandleFunc("/blog/{guid}", servePage).Methods("GET")
    fmt.Println("Serving on port 3000")
    http.ListenAndServe(":3000", r)
    
    

}
