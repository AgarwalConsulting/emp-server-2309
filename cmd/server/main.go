package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"algogrit.com/empserver/employees/repository"
	"algogrit.com/empserver/employees/service"
	"algogrit.com/empserver/entities"
)

var empRepo = repository.NewInMem()
var empV1Svc = service.NewV1(empRepo)

func EmployeesIndexHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	employees, err := empV1Svc.Index()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	json.NewEncoder(w).Encode(employees)
}

func EmployeeCreateHandler(w http.ResponseWriter, req *http.Request) {
	var newEmp entities.Employee
	err := json.NewDecoder(req.Body).Decode(&newEmp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	createdEmp, err := empV1Svc.Create(newEmp)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	json.NewEncoder(w).Encode(createdEmp)
}

func EmployeesHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		EmployeeCreateHandler(w, req)
	} else {
		EmployeesIndexHandler(w, req)
	}
}

// func LoggingMiddleware(handler http.Handler) http.Handler {
// 	h := func(w http.ResponseWriter, req *http.Request) {
// 		begin := time.Now()

// 		handler.ServeHTTP(w, req)

// 		log.Printf("%s %s tooks %s\n", req.Method, req.URL, time.Since(begin))
// 	}

// 	return http.HandlerFunc(h)
// }

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		msg := "Hello, World!"

		fmt.Fprintln(w, msg)
	})

	r.HandleFunc("/employees", EmployeesIndexHandler).Methods("GET")
	r.HandleFunc("/employees", EmployeeCreateHandler).Methods("POST")

	http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, r))
}
