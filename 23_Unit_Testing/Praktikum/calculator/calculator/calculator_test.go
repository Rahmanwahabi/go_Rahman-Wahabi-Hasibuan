package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddition(t *testing.T) {
	result := Addition(3, 5)
	assert.Equal(t, float64(8), result)
}

func TestSubtraction(t *testing.T) {
	result := Subtraction(10, 7)
	assert.Equal(t, float64(3), result)
}

func TestDivision(t *testing.T) {
	result, err := Division(15, 3)
	assert.NoError(t, err)
	assert.Equal(t, float64(5), result)

	// Test division by zero
	_, err = Division(10, 0)
	assert.Error(t, err)
	assert.EqualError(t, err, "division by zero is not allowed")
}

func TestMultiplication(t *testing.T) {
	result := Multiplication(4, 6)
	assert.Equal(t, float64(24), result)
}
