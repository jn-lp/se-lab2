package lab2

import (
	"testing"
	"strings"
	"bytes"

	"github.com/stretchr/testify/assert"
)

func TestSuccess(t *testing.T) {
	output := new(bytes.Buffer)

	handler := ComputeHandler{
		Input: strings.NewReader("1 2 + 3 - 4 * 5 / 6 ^ 7 + 8 - 9 * 10 /"),
		Output: output,
	}

	err := handler.Compute()

	assert.Nil(t, err)
	assert.Equal(t, "(((1+2-3)*4/5)^6+7-8)*9/10\n", output.String())
}

func TestFailure(t *testing.T) {
	handler := ComputeHandler{
		Input: strings.NewReader("1 @ +"),
		Output: new(bytes.Buffer),
	}

	err := handler.Compute()

	assert.NotNil(t, err)
}
