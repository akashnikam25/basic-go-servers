package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	if r.URL.Path != "/hello" {
		http.Error(w, "Path not found", 404)
		return
	}
	fmt.Fprintf(w, "Hello, World!")

}
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), 400)
	}
	fmt.Fprintf(w, "Post Request successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "name: %s\n", name)
	fmt.Fprintf(w, "address: %s\n", address)

}
func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)
	fmt.Println("Starting server...")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}
