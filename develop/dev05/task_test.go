package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrep(t *testing.T) {
	tests := []struct {
		name    string
		lines   []string
		pattern string
		flags   flags
		want    []string
	}{
		{
			name:    "Test count",
			lines:   []string{"1112111", "1"},
			pattern: "2",
			flags: flags{
				count: true,
			},
			want: []string{"1"},
		},
		{
			name:    "Test ignoreCase",
			lines:   []string{"A", "B", "C", "a"},
			pattern: "A",
			flags: flags{
				ignoreCase: true,
			},
			want: []string{"A", "a"},
		},
		{
			name:    "Test lineNum",
			lines:   []string{"A", "B", "C", "a"},
			pattern: "A",
			flags: flags{
				lineNum: true,
			},
			want: []string{"1:A"},
		},
		{
			name:    "Test invert",
			lines:   []string{"Ab", "Bb", "C", "a"},
			pattern: "b",
			flags: flags{
				invert: true,
			},
			want: []string{"C", "a"},
		},
		{
			name:    "Test invert and count",
			lines:   []string{"Ab", "Bb", "C", "a"},
			pattern: "b",
			flags: flags{
				invert: true,
				count:  true,
			},
			want: []string{"2"},
		},
		{
			name:    "Test fixed",
			lines:   []string{"abcd", "a"},
			pattern: "a",
			flags: flags{
				fixed: true,
			},
			want: []string{"a"},
		},
		{
			name:    "Test fixed and ignore case",
			lines:   []string{"Abcd", "A", "abcd", "a"},
			pattern: "a",
			flags: flags{
				fixed:      true,
				ignoreCase: true,
			},
			want: []string{"A", "a"},
		},
		{
			name:    "Test fixed and ignore case",
			lines:   []string{"Abcd", "A", "abcd", "a"},
			pattern: "a",
			flags: flags{
				fixed:      true,
				ignoreCase: true,
			},
			want: []string{"A", "a"},
		},
		{
			name:    "Test -A",
			lines:   []string{"a", "1"},
			pattern: "a",
			flags: flags{
				after: 1,
			},
			want: []string{"a", "1"},
		},
		{
			name:    "Test -A строк после больше чем есть",
			lines:   []string{"a", "1"},
			pattern: "a",
			flags: flags{
				after: 5,
			},
			want: []string{"a", "1"},
		},
		{
			name:    "Test -A проверка на дубликаты",
			lines:   []string{"a", "a", "1", "b"},
			pattern: "a",
			flags: flags{
				after: 1,
			},
			want: []string{"a", "a", "1"},
		},
		{
			name:    "Test -A -n",
			lines:   []string{"a", "a", "1", "b"},
			pattern: "a",
			flags: flags{
				after:   1,
				lineNum: true,
			},
			want: []string{"1:a", "2:a", "3:1"},
		},
		{
			name:    "Test -B 1",
			lines:   []string{"1", "b", "a"},
			pattern: "b",
			flags: flags{
				before: 1,
			},
			want: []string{"1", "b"},
		},
		{
			name:    "Test -C 1",
			lines:   []string{"left2", "left1", "a", "right1", "right2"},
			pattern: "a",
			flags: flags{
				context: 1,
			},
			want: []string{"left1", "a", "right1"},
		},
	}
	for _, test := range tests {
		get := Grep(test.lines, test.pattern, test.flags)
		assert.Equal(t, test.want, get, test.name)
		fmt.Printf("%s DONE\n", test.name)
	}
}
