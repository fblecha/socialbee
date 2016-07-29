package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	routes()
	setupServer()
}

func setupServer() {
	port := ":9091"
	fmt.Printf("Running server on port %v , see http://localhost %v \n", port, port)
	err := http.ListenAndServe(port, nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
