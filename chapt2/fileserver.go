package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting")
	log.Println("Serving files ....")
	
	http.ListenAndServe(":3000", http.FileServer(http.Dir("../../../")))
}
