package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello web")
	greet()
	res := sum(1, 6)
	fmt.Println(res)

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":6969", r))
}

func greet() {
	fmt.Println("Hello from greet")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello from router </h1>"))
}

func sum(a int, b int) int {
	return a + b
}
