package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"html/template"
)

type Todo struct {
	Title string
	Done bool
}

type TodoPageData struct{
	PageTitle string
	TodoList []Todo
}

func main (){
	r := mux.NewRouter()


	r.HandleFunc("/", func(w http.ResponseWriter, r * http.Request){
		fmt.Fprintf(w,"Welcome to my website" )
	})

	
	todoTemplate := template.Must(template.ParseFiles("templates/todo.html"))
	r.HandleFunc("/todo",  func(w http.ResponseWriter, r * http.Request){
		data := TodoPageData{
			PageTitle: "My Todo List",
			TodoList: []Todo{
				{Title: "Task 1", Done: false},
                {Title: "Task 2", Done: true},
                {Title: "Task 3", Done: true},
			},
		}

		todoTemplate.Execute(w, data)
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