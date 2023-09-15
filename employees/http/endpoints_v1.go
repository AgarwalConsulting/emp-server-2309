package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"algogrit.com/empserver/entities"
)

func (h EmployeeHandler) IndexV1(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	employees, err := h.v1Svc.Index()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	json.NewEncoder(w).Encode(employees)
}

func (h EmployeeHandler) CreateV1(w http.ResponseWriter, req *http.Request) {
	var newEmp entities.Employee
	err := json.NewDecoder(req.Body).Decode(&newEmp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	createdEmp, err := h.v1Svc.Create(newEmp)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	json.NewEncoder(w).Encode(createdEmp)
}
