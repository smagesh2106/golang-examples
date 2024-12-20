package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/foo", handleFoo).Methods("GET")
	r.HandleFunc("/bar/{category}/bar/{id}", handleBar).Methods("GET", "POST")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Println(err)
	}
}

func handleFoo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("GET Method")
	}
	fmt.Printf("%v, %s\n", r.URL, r.Method)
}

func handleBar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if r.Method == "GET" {
		fmt.Printf("GET Method, category :%s, id: %s", params["category"], params["id"])
	}
	if r.Method == "POST" {
		fmt.Printf("POST Method, category :%s, id: %s", params["category"], params["id"])
		//fmt.Println("Body :" + r.GetBody())

	}

	fmt.Printf(" %v, %s\n", r.URL, r.Method)
}
