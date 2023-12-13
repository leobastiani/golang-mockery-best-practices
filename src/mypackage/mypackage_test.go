package mypackage_test

import (
	"context"
	"testing"

	"github.com/leobastiani/golang-mockery-best-practices/src/mypackage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMyComponent(t *testing.T) {
	// Create a new instance of the mock object
	myInterface := mypackage.NewMockMyInterface[string](t)
	myInterface.EXPECT().MyMethod(
		mock.MatchedBy(func(_ context.Context) bool { return true }),
		mock.MatchedBy(func(s string) bool {
			return len(s) > 0
		}),
	).Return(10, nil)

	// Create an instance of the component under test
	c := mypackage.NewMyComponent(myInterface)

	// Call the method being tested
	result, err := c.MyMethod(context.Background(), "some non zero string")

	// Check that the result and error match the expected values
	assert.NoError(t, err)
	assert.Equal(t, 10, result)

	myInterface.AssertExpectations(t)
}
