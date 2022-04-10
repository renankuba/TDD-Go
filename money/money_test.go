package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := new(Dollar)
	five.Init(5)
	five.times(2)
	assert.Equal(t, five.amount, 10.0)
}
