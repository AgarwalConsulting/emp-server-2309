package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"

	"algogrit.com/empserver/employees/repository"
	"algogrit.com/empserver/employees/service"
	"algogrit.com/empserver/entities"
)

func TestIndexV1(t *testing.T) {
	ctrl := gomock.NewController(t)
	// defer ctrl.Finish()

	mockRepo := repository.NewMockEmployeeRepository(ctrl)
	sut := service.NewV1(mockRepo)

	expectedEmps := []entities.Employee{
		{1, "Gaurav", "LnD", 1001},
	}

	mockRepo.EXPECT().ListAll().Return(expectedEmps, nil)

	actualEmps, err := sut.Index()

	assert.Nil(t, err)
	assert.NotNil(t, actualEmps)

	assert.Equal(t, expectedEmps, actualEmps)
}
