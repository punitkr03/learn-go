package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello web")
	greet()
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

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all the courses.")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]string{"course1", "course2", "course3"})
}
