package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// "a4bc2d5e" => "aaaabccddddde"
// "abcd" => "abcd"
// "45" => "" (некорректная строка)
// "" => ""

func TestUnpack(t *testing.T) {
	tests := []struct {
		input  string
		wanted string
		err    error
	}{
		{
			"a4bc2d5e",
			"aaaabccddddde",
			nil,
		},
		{
			"abcd",
			"abcd",
			nil,
		},
		{
			"45",
			"",
			ErrorInvalidString,
		},
		{
			"",
			"",
			nil,
		},
	}
	for _, test := range tests {
		result, err := Unpack(test.input)
		assert.ErrorIs(t, err, test.err, "wrong error")
		assert.EqualValues(t, test.wanted, result, "wrong result")
	}
}
