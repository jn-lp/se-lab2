package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostfixToInfix(t *testing.T) {
	res, err := PostfixToInfix("1 22 + 333 * 4444 - 55555 / 6 ^")
	if assert.Nil(t, err) {
		assert.Equal(t, "(((1+22)*333-4444)/55555)^6", res)
	}
}

func ExamplePostfixToInfix() {
	res, _ := PostfixToInfix("1 2 +")
	fmt.Println(res)

	// Output:
	// 1+2
}
