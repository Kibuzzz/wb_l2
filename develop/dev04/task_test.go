package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnograms(t *testing.T) {
	tests := []struct {
		input  []string
		wanted map[string][]string
	}{
		{
			input:  []string{"пятка", "пятак", "тяпка", "листок", "слиток", "столик"},
			wanted: map[string][]string{"листок": []string{"листок", "слиток", "столик"}, "пятка": []string{"пятка", "пятак", "тяпка"}},
		},
	}
	for _, test := range tests {
		get := Anograms(test.input)
		assert.Equal(t, test.wanted, get, "wrong answer")
	}
}
