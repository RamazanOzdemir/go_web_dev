package login

import (
	"log"
	"net/http"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Logging() Middleware{
	return func(f http.HandlerFunc) http.HandlerFunc {
		// Define the http.HandlerFunc

		return func( w http.ResponseWriter, r *http.Request){
			
			// Do middleware things
			start := time.Now()
			defer func(){ log.Println(r.URL.Path, time.Since(start))}()

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

// applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares{
		f=m(f)
	}

	return f
}