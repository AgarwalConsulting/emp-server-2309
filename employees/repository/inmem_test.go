package repository_test

import (
	"sync"
	"testing"

	"algogrit.com/empserver/employees/repository"
	"algogrit.com/empserver/entities"
	"github.com/stretchr/testify/assert"
)

func TestConsistency(t *testing.T) {
	sut := repository.NewInMem()

	initialEmps, err := sut.ListAll()

	assert.Nil(t, err)
	assert.NotNil(t, initialEmps)

	assert.Equal(t, 3, len(initialEmps))

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			emp := entities.Employee{Name: "Gaurav"}
			sut.Save(emp)
			sut.ListAll()
		}()
	}

	wg.Wait()

	finalEmps, err := sut.ListAll()

	assert.Nil(t, err)
	assert.NotNil(t, finalEmps)

	assert.Equal(t, 103, len(finalEmps))
}
