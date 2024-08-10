package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID        int     `json:"id"`
	Author    string  `json:"author"`
	Publisher string  `json:"publisher"`
	Date      string  `json:"date"`
	Price     float64 `json:"price"`
	ISBN      string  `json:"isbn"`
}

var all_books = []Book{
	{ID: 1, Author: "Abdul Haque", Publisher: "Rex Books", Date: "2022", Price: 30.90, ISBN: "978-3-16-148410-0"},
	{ID: 2, Author: "J.K. Rowling", Publisher: "Bloomsbury", Date: "1997", Price: 20.99, ISBN: "978-0-7475-3269-9"},
	{ID: 3, Author: "George Orwell", Publisher: "Secker & Warburg", Date: "1949", Price: 15.99, ISBN: "978-0-452-28423-4"},
	{ID: 4, Author: "F. Scott Fitzgerald", Publisher: "Charles Scribner's Sons", Date: "1925", Price: 10.99, ISBN: "978-0-7432-7356-5"},
	{ID: 5, Author: "Harper Lee", Publisher: "J.B. Lippincott & Co.", Date: "1960", Price: 18.99, ISBN: "978-0-06-112008-4"},
	{ID: 6, Author: "J.R.R. Tolkien", Publisher: "George Allen & Unwin", Date: "1954", Price: 25.99, ISBN: "978-0-618-00222-8"},
	{ID: 7, Author: "Herman Melville", Publisher: "Harper & Brothers", Date: "1851", Price: 12.99, ISBN: "978-0-14-243724-7"},
	{ID: 8, Author: "Jane Austen", Publisher: "T. Egerton", Date: "1813", Price: 14.99, ISBN: "978-0-19-953556-9"},
	{ID: 9, Author: "Leo Tolstoy", Publisher: "The Russian Messenger", Date: "1869", Price: 19.99, ISBN: "978-0-14-303500-8"},
	{ID: 10, Author: "Mark Twain", Publisher: "Chatto & Windus", Date: "1884", Price: 13.99, ISBN: "978-0-14-243717-9"},
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	var filteredBooks []Book
	qParams := r.URL.Query()
	author := qParams.Get("author")
	publisher := qParams.Get("publisher")

	if author != "" || publisher != "" {
		for _, book := range all_books {
			if book.Author == author || book.Publisher == publisher {
				filteredBooks = append(filteredBooks, book)
			}

		}
		if len(filteredBooks) == 0 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"message": "No books found"})
			return
		}
	} else {
		filteredBooks = all_books
	}

	booksJSON, err := json.Marshal(filteredBooks)
	if err != nil {
		http.Error(w, "Unable to load data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(booksJSON)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	if idStr == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var foundBook Book
	for _, book := range all_books {
		if book.ID == id {
			foundBook = book
			break
		}
	}

	// if foundBook  {
	// 	http.Error(w, "Book not found", http.StatusNotFound)
	// 	return
	// }

	jsonResponse, err := json.Marshal(foundBook)
	if err != nil {
		http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/book/{id}", getBook).Methods("GET")
	r.HandleFunc("/", index).Methods("GET")
	fmt.Println("Server running on port 8000...")
	http.ListenAndServe(":3000", r)
}
