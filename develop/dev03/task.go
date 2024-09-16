package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	column := *flag.Int("k", -1, `-k — указание колонки для сортировки`)        // id после разделения
	numeric := *flag.Bool("n", false, `-n — сортировать по числовому значению`) // разделитель
	reverse := *flag.Bool("r", false, `сортировать в обратном порядке`)         // отображать ли строки без разделителя
	removeDuplicates := *flag.Bool("u", false, `не выводить повторяющиеся строки`)

	flag.Parse()

	if column == -1 {
		column = 0
	}

	fileName := flag.Arg(0)

	lines, err := readFile(fileName)
	if err != nil {
		log.Fatalf("error reading file: %s\n", err.Error())
	}

	if removeDuplicates {
		lines = removeLineDuplicates(lines)
	}

	sortLines(lines, column, numeric)

	if reverse {
		reverseLines(lines)
	}

	fmt.Println(lines)
}

func readFile(fileName string) ([]string, error) {

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func sortLines(lines []string, k int, n bool) {
	slices.SortFunc(lines, func(line1, line2 string) int {
		words1 := strings.Split(line1, " ")
		words2 := strings.Split(line2, " ")
		// любая_строка_без_нужной_колонки < строка_с_нужной_колонкой
		// любая_строка_без_нужной_колонки == любая_строка_без_нужной_колонки
		if k > len(words1) || k > len(words2) {
			switch {
			case k < len(words1) && k > len(words2):
				return 1
			case k > len(words1) && k < len(words2):
				return -1
			case k > len(words1) && k > len(words2):
				return 0
			}
		}
		if n {
			num1, err1 := strconv.Atoi(words1[k])
			num2, err2 := strconv.Atoi(words2[k])
			// строку не получилось преобразовать < любая строка, которую получилось преобразовать
			// строку не получилось преобразовать == строку не получилось преобразовать
			switch {
			case err1 == nil && err2 != nil:
				return 1
			case err1 != nil && err2 == nil:
				return -1
			case err1 != nil && err2 != nil:
				return 0
			}

			switch {
			case num1 < num2:
				return -1
			case num1 > num2:
				return 1
			default:
				return 0
			}
		} else {
			switch {
			case words1[k] < words2[k]:
				return -1
			case words1[k] > words2[k]:
				return 1
			default:
				return 0
			}
		}

	})
}

func reverseLines(lines []string) {
	slices.Reverse(lines)
}

func removeLineDuplicates(lines []string) []string {
	duplicates := make(map[string]struct{})
	var withoutDuplicates []string
	for _, line := range lines {
		if _, duplicate := duplicates[line]; !duplicate {
			withoutDuplicates = append(withoutDuplicates, line)
			duplicates[line] = struct{}{}
		}
	}
	return withoutDuplicates
}
