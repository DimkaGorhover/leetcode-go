package p0022

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	t.Parallel()
	t.Run(``, func(t *testing.T) {
		result := generateParenthesis(3)
		assert.Contains(t, result, "((()))")
		assert.Contains(t, result, "(()())")
		assert.Contains(t, result, "(())()")
		assert.Contains(t, result, "()(())")
		assert.Contains(t, result, "()()()")
	})
}
