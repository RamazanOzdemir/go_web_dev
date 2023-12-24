module example.com/web_app

go 1.21.5

require (
	example.com/login v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.1
)

replace example.com/login => ./middlewares
