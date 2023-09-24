package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CanExecuteCB(t *testing.T) {
	assert.False(t, CheckBraces("("))
	assert.False(t, CheckBraces("}"))
	assert.False(t, CheckBraces("(("))
	assert.False(t, CheckBraces(")("))
	assert.False(t, CheckBraces("({)}"))

	assert.True(t, CheckBraces("{}"))
	assert.True(t, CheckBraces("{()}"))
	assert.True(t, CheckBraces("{[]()}"))
	assert.True(t, CheckBraces("{[(())]()}"))
}
