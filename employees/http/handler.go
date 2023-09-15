package http

import (
	"algogrit.com/empserver/employees/service"
	"github.com/gorilla/mux"
)

type EmployeeHandler struct {
	*mux.Router // Embedding
	v1Svc       service.EmployeeService
}

func (h *EmployeeHandler) SetupRoutes(r *mux.Router) {
	r.HandleFunc("/v1/employees", h.IndexV1).Methods("GET")
	r.HandleFunc("/v1/employees", h.CreateV1).Methods("POST")

	h.Router = r
}

func NewHandler(v1Svc service.EmployeeService) EmployeeHandler {
	h := EmployeeHandler{v1Svc: v1Svc}

	h.SetupRoutes(mux.NewRouter())

	return h
}
