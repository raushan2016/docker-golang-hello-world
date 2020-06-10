package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	init := os.Getenv("INIT")
	if init != "" {
		fmt.Println("Hello world, init done!!!!!!")
		return
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := mux.NewRouter()
	r.HandleFunc("/", hello)

	http.Handle("/", r)
	fmt.Println("Starting up on " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}
