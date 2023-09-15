package http_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	empHTTP "algogrit.com/empserver/employees/http"
	"algogrit.com/empserver/employees/service"
	"algogrit.com/empserver/entities"
)

func TestCreateV1Handler(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := service.NewMockEmployeeService(ctrl)

	sut := empHTTP.NewHandler(mockService)

	expectedEmp := entities.Employee{Name: "Gaurav", Department: "LnD", ProjectID: 10001}
	createdEmp := expectedEmp
	createdEmp.ID = 1

	mockService.EXPECT().Create(expectedEmp).Return(&createdEmp, nil)

	jsonString := `{"name": "Gaurav", "speciality": "LnD", "project": 10001}` // Type: string

	reqBody := strings.NewReader(jsonString)
	req := httptest.NewRequest("POST", "/v1/employees", reqBody) // io.Reader
	respRec := httptest.NewRecorder()

	// sut.CreateV1(respRec, req)
	sut.ServeHTTP(respRec, req)

	assert.Equal(t, http.StatusOK, respRec.Code)

	var actualEmp entities.Employee

	resp := respRec.Result()

	json.NewDecoder(resp.Body).Decode(&actualEmp)

	assert.Equal(t, expectedEmp, actualEmp)
}
