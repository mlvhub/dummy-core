package core_test

import (
	"testing"

	"github.com/mlvhub/dummy-core/core"
	"github.com/stretchr/testify/assert"
)

func TestToLowercase(t *testing.T) {
	tests := map[string]string{
		"hello": "hello",
		"HELLO": "hello",
		"HeLlO": "hello",
	}

	for str, expected := range tests {
		actual := core.ToLowercase(str)
		assert.Equal(t, actual, expected)
	}
}
