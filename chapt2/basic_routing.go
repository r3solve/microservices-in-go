package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const (
	PORT = ":4000"
)

func pageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageID := vars["page"]
	fileName := "files/" + pageID + ".html"
	if _, err:= os.Stat(fileName); err != nil {
		fileName = "files/404.html"
	}
	
	http.ServeFile(w,r,fileName)

	}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{page}", pageHandler).Methods("GET")
	fmt.Println("Server is running on port " + PORT)
	http.Handle("/", r)
	http.ListenAndServe(PORT, nil)
}