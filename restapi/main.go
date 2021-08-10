package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

// Init books var as a slice Book struct
var books []Book

func main() {
	// initialize router
	router := mux.NewRouter()

	//Mock data
	books = append(books, Book{Id:"1",Isbn:"8779911",Title:"Midsummer's dream",
		Author:&Author{FirstName:"James",LastName:"Chase"}})
	books = append(books, Book{Id:"2",Isbn:"3129914",Title:"No longer at ease",
		Author:&Author{FirstName:"Nkiru",LastName:"Soprinye"}})

	// route handlers and endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))

}

//Book struct
type Book struct {
	Id     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}


// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(books)
}

// Get Single Book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)

	//loop through books and find with id
	for _, item := range books{
		if item.Id == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Create Book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.Id = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)

}

// Update Book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index, item := range books{
		if item.Id == params["id"]{
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.Id = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
		json.NewEncoder(w).Encode(books)
	}

}

// Delete Book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index, item := range books{
		if item.Id == params["id"]{
			books = append(books[:index], books[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(books)
	}
}
