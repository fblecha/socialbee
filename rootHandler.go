package main

import (
	//"fmt"
	"html/template"
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("root handler")
	t, _ := template.ParseFiles("templates/index.html")
	err := t.Execute(w, nil)
	//fmt.Println("execute done")
	if err != nil {
		log.Fatal(err)
	}
}
