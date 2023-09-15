package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	empHTTP "algogrit.com/empserver/employees/http"
	"algogrit.com/empserver/employees/repository"
	"algogrit.com/empserver/employees/service"
)

var empRepo = repository.NewInMem()
var empV1Svc = service.NewV1(empRepo)
var empHandler = empHTTP.NewHandler(empV1Svc)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		msg := "Hello, World!"

		fmt.Fprintln(w, msg)
	})

	empHandler.SetupRoutes(r)

	http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, r))
}
