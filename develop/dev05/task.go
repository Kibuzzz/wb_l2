package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type flags struct {
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
}

func main() {
	after := flag.Int(`A`, 0, `"after" печатать +N строк после совпадения`)
	before := flag.Int(`B`, 0, `-B - "before" печатать +N строк до совпадения`)
	context := flag.Int(`C`, 0, `-C - "context" (A+B) печатать ±N строк вокруг совпадения`)
	count := flag.Bool(`c`, false, `-c - "count" (количество строк)`)
	ignoreCase := flag.Bool(`i`, false, `-i - "ignore-case" (игнорировать регистр)`)
	invert := flag.Bool(`v`, false, `-v - "invert" (вместо совпадения, исключать)`)
	fixed := flag.Bool(`F`, false, `-F - "fixed", точное совпадение со строкой, не паттерн`)
	lineNum := flag.Bool(`n`, false, `-n - "line num", печатать номер строки`)

	flag.Parse()

	flags := flags{
		after:      *after,
		before:     *before,
		context:    *context,
		count:      *count,
		ignoreCase: *ignoreCase,
		invert:     *invert,
		fixed:      *fixed,
		lineNum:    *lineNum,
	}

	if *after < 0 {
		log.Fatal("after flag should be >= 0")
	}

	// Чтение аргументов командной строки
	pattern := flag.Arg(0)
	fileName := flag.Arg(1)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("failed to open file %s\n", err.Error())
	}
	file.Close()

	lines := []string{}
	Grep(lines, pattern, flags)
}

func Grep(lines []string, pattern string, flags flags) []string {

	var counter int
	var output []string
	linesInOutput := make(map[int]bool) // чтобы не допустить повторного добавления уже довабленной строки

	if flags.context != 0 {
		flags.after = flags.context
		flags.before = flags.context
	}

	for lineIndex, line := range lines {

		// Для проверки есть ли совпадение в строке
		match := findSubString(line, pattern, flags)

		var appendLine bool

		switch {
		case flags.invert && !match:
			appendLine = true
		case !flags.invert && match:
			appendLine = true
		}

		if appendLine {
			counter++
			switch {
			case flags.context > 0:
				fallthrough
			case flags.before > 0:
				rightBarier := lineIndex
				leftBarier := lineIndex - flags.before
				if leftBarier < 0 {
					leftBarier = 0
				}
				output = appendLines(lines, output, linesInOutput, leftBarier, rightBarier, flags)
				fallthrough
			case flags.after > 0:
				leftBarier := lineIndex
				rightBarier := lineIndex + flags.after
				if rightBarier >= len(lines) {
					rightBarier = len(lines) - 1
				}
				output = appendLines(lines, output, linesInOutput, leftBarier, rightBarier, flags)
			default:
				if flags.lineNum {
					line = fmt.Sprintf("%d:%s", lineIndex+1, line)
				}
				output = append(output, line)
			}
		}
	}

	if flags.count {
		numLines := strconv.Itoa(counter)
		output = []string{numLines}
	}

	return output
}

func appendLines(lines []string, output []string, duplicates map[int]bool, leftBarier int, rightBarier int, flags flags) []string {
	for i := leftBarier; i <= rightBarier; i++ {
		line := lines[i]
		if flags.lineNum {
			line = fmt.Sprintf("%d:%s", i+1, line)
		}
		if _, duplicate := duplicates[i]; !duplicate {
			output = append(output, line)
			duplicates[i] = true
		}
	}
	return output
}

func findSubString(line string, pattern string, flags flags) bool {
	if flags.fixed {

		var words []string

		if flags.ignoreCase {
			words = strings.Split(strings.ToLower(line), " ")
			pattern = strings.ToLower(pattern)
		} else {
			words = strings.Split(line, " ")
		}

		for _, word := range words {
			if word == pattern {
				return true
			}
		}
		return false
	}
	if flags.ignoreCase {
		return strings.Contains(strings.ToLower(line), strings.ToLower(pattern))
	}
	return strings.Contains(line, pattern)
}
