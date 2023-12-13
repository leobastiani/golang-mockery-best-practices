package mypackage_test

import (
	"testing"

	"github.com/leobastiani/golang-mockery-best-practices/src/mypackage"
	mypackage_mocks "github.com/leobastiani/golang-mockery-best-practices/src/mypackage/mocks"
	"github.com/stretchr/testify/assert"
)

func TestMyComponent(t *testing.T) {
	// Create a new instance of the mock object
	mock := mypackage_mocks.NewMyInterface(t)
	mock.EXPECT().MyMethod().Return(10)

	// Create an instance of the component under test
	c := mypackage.NewMyComponent(mock)

	// Call the method being tested
	result := c.MyMethod()

	// Check that the result and error match the expected values
	assert.Equal(t, 10, result)
}
