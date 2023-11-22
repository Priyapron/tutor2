package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	// Handle requests with Gorilla Mux
	r := mux.NewRouter()

	// Route for adding two numbers
	r.HandleFunc("/cal/{num1}/plus/{num2}", AddHandler)

	// Route for handling books
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		b := []byte("<h1>Book: " + title + "</h1><br><h2>page: " + page + "</h2>")
		w.Write(b)
	})

	// Route for the root ("/") URL
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	// Route for the "/about" URL
	r.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "My name is Priyapron Ophatnithiwat")
	})

	// Start the HTTP server on port 8080 with the router
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}

// AddHandler handles the addition of two numbers
func AddHandler(w http.ResponseWriter, r *http.Request) {
	// Get the values from the route parameters
	vars := mux.Vars(r)
	num1, err := strconv.Atoi(vars["num1"])
	if err != nil {
		http.Error(w, "Invalid number format for num1", http.StatusBadRequest)
		return
	}

	num2, err := strconv.Atoi(vars["num2"])
	if err != nil {
		http.Error(w, "Invalid number format for num2", http.StatusBadRequest)
		return
	}

	// Perform the addition
	total := num1 + num2

	// Respond with the result
	fmt.Fprintf(w, "Total = %d\n", total)
}
