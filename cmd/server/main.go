package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Employee struct {
	ID         int    `json:"-"`
	Name       string `json:"name"`
	Department string `json:"speciality"`
	ProjectID  int    `json:"project"`
}

var employees = []Employee{
	{1, "Gaurav", "LnD", 1001},
	{2, "Anupam", "Cloud", 10002},
	{3, "Udbhav", "SRE", 20002},
}

func EmployeesIndexHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func EmployeeCreateHandler(w http.ResponseWriter, req *http.Request) {
	var newEmp Employee
	err := json.NewDecoder(req.Body).Decode(&newEmp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	newEmp.ID = len(employees) + 1

	employees = append(employees, newEmp)

	json.NewEncoder(w).Encode(newEmp)
}

func EmployeesHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		EmployeeCreateHandler(w, req)
	} else {
		EmployeesIndexHandler(w, req)
	}
}

func LoggingMiddleware(handler http.Handler) http.Handler {
	h := func(w http.ResponseWriter, req *http.Request) {
		begin := time.Now()

		handler.ServeHTTP(w, req)

		log.Printf("%s %s tooks %s\n", req.Method, req.URL, time.Since(begin))
	}

	return http.HandlerFunc(h)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		msg := "Hello, World!"

		fmt.Fprintln(w, msg)
	})

	r.HandleFunc("/employees", EmployeesIndexHandler).Methods("GET")
	r.HandleFunc("/employees", EmployeeCreateHandler).Methods("POST")

	http.ListenAndServe("localhost:8000", LoggingMiddleware(r))
}
