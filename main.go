package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main (){
	r := mux.NewRouter()


	r.HandleFunc("/", func(w http.ResponseWriter, r * http.Request){
		fmt.Fprintf(w,"Welcome to my website" )
	})

	fs := http.FileServer(http.Dir("static/"))

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		title := vars["title"]
		page := vars["page"]
		        
		
		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)



	})

	
	log.Println("Server will start at localhost")
	http.ListenAndServe(":80", r)
}