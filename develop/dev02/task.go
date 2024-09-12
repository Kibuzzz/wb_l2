package main

import (
	"errors"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	ErrorInvalidString = errors.New("некорректная строка")
)

func isValid(runes []rune) bool {
	for i := 0; i < len(runes)-1; i++ {
		cur := runes[i]
		next := runes[i+1]
		if unicode.IsDigit(cur) && unicode.IsDigit(next) {
			return false
		}
	}
	return true
}

func Unpack(input string) (string, error) {
	runes := []rune(input)

	if len(runes) == 0 {
		return "", nil
	}

	if !isValid(runes) {
		return "", ErrorInvalidString
	}

	l := len(runes)
	var result string
	for i := 0; i < l-1; i++ {
		cur := runes[i]
		next := runes[i+1]
		if unicode.IsLetter(cur) && unicode.IsDigit(next) {
			n := int(next - '0')
			result += strings.Repeat(string(cur), n)
			i++ // потому что уже посмотрели - это цифра
		} else if unicode.IsLetter(cur) {
			result += string(cur)
		}
	}
	result += string(runes[l-1])
	return result, nil
}

func main() {

}
