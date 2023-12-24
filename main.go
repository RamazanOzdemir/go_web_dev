package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"html/template"

	"example.com/login"
)

type Todo struct {
	Title string
	Done bool
}

type TodoPageData struct{
	PageTitle string
	TodoList []Todo
}

type ContactDetail struct {
	Email string
	Subject string
	Message string
}

func main (){
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("static/"))

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))


	// HOME
	r.HandleFunc("/", func(w http.ResponseWriter, r * http.Request){
		fmt.Fprintf(w,"Welcome to my website" )
	})

	
	// TodoList PAGE
	todoTemplate := template.Must(template.ParseFiles("templates/todo.html"))

	todoHandler := func(w http.ResponseWriter, r * http.Request){
		data := TodoPageData{
			PageTitle: "My Todo List",
			TodoList: []Todo{
				{Title: "Task 1", Done: false},
                {Title: "Task 2", Done: true},
                {Title: "Task 3", Done: true},
			},
		}

		todoTemplate.Execute(w, data)
	}

	// apply login middleware to todo service
	r.HandleFunc("/todo",  login.Chain(todoHandler, login.Logging()) )

	contactFormTemplate := template.Must(template.ParseFiles("templates/contact.html"))
	// CONTACT PAGE
	r.HandleFunc("/contact", func(w http.ResponseWriter, r * http.Request){
		
		if r.Method != http.MethodPost {
            contactFormTemplate.Execute(w, nil)
            return
        }
		details := ContactDetail{
            Email:   r.FormValue("email"),
            Subject: r.FormValue("subject"),
            Message: r.FormValue("message"),
        }

        // do something with details
        _ = details

        contactFormTemplate.Execute(w, struct{ Success bool }{true})
	})




	// BOOKS PAGE with dynamic router parameters
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		title := vars["title"]
		page := vars["page"]
		        
		
		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)



	})

	
	log.Println("Server will start at localhost")
	http.ListenAndServe(":80", r)
}