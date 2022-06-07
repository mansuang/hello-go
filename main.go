package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var books []Book

type Book struct {
	Id     string  `json:"id"`
	Title  string  `json:"title"`
	Isbn   string  `json:"isbn"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range books {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(Book{})
}

func main() {
	path := os.Getenv("APIPATH")
	r := mux.NewRouter()

	book1 := Book{Id: "1", Title: "Book One", Isbn: "1221", Author: &Author{Firstname: "Vivek", Lastname: "Singh"}}
	book2 := Book{Id: "2", Title: "Book Two", Isbn: "2211", Author: &Author{Firstname: "Abhash", Lastname: "Kumar"}}

	books = append(books, book1)
	books = append(books, book2)

	r.HandleFunc(path, getBooks).Methods("GET")
	r.HandleFunc(path+"/{id}", getBook).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}