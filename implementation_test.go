package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostfixToInfix(t *testing.T) {
	res, err := PostfixToInfix("1")
	if assert.Nil(t, err) {
		assert.Equal(t, "1", res)
	}
	res, err = PostfixToInfix("1 2 +")
	if assert.Nil(t, err) {
		assert.Equal(t, "1+2", res)
	}
	res, err = PostfixToInfix("1 22 + 333 * 4444 - 55555 / 6 ^")
	if assert.Nil(t, err) {
		assert.Equal(t, "(((1+22)*333-4444)/55555)^6", res)
	}
	res, err = PostfixToInfix("1 2+	3 - 4 * 5 / 6 ^ 7 + 8 - 9 *10	/")
	if assert.Nil(t, err) {
		assert.Equal(t, "(((1+2-3)*4/5)^6+7-8)*9/10", res)
	}
	res, err = PostfixToInfix("1 +")
	if !assert.NotNil(t, err) {
		t.Fatal("Unary operator error not implemented")
	}
	res, err = PostfixToInfix("")
	if !assert.NotNil(t, err) {
		t.Fatal("Empty string error not implemented")
	}
}

func ExamplePostfixToInfix() {
	res, _ := PostfixToInfix("1 2 +")
	fmt.Println(res)

	// Output:
	// 1+2
}
