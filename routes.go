package main

import (
	"net/http"
)

func routes() {
	http.HandleFunc("/", rootHandler) // setting router rule
	http.HandleFunc("/checksocial", checkSocialHandler)
}
