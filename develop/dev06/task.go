package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	fields := flag.String("f", "", `"fields" - выбрать поля (колонки)`)                  // id после разделения
	delimiter := flag.String("d", "\t", `"delimiter" - использовать другой разделитель`) // разделитель
	separated := flag.Bool("s", false, `"separated" - только строки с разделителем`)     // отображать ли строки без разделителя

	flag.Parse()

	if *fields == "" {
		log.Fatal("flag f required")
	}

	requiredFieldStr := strings.Split(*fields, ",")
	var requiredFieldInt []int
	for _, field := range requiredFieldStr {
		i, err := strconv.Atoi(field)
		if err != nil || i < 0 {
			log.Fatal("bad field")
		}
		requiredFieldInt = append(requiredFieldInt, i-1)
	}

	lines := input()
	var output []string
	// filter words
	for _, line := range lines {
		lineWords := strings.Split(line, *delimiter)
		if len(lineWords) == 1 && *separated {
			continue
		}
		var outputString string
		for _, i := range requiredFieldInt {
			if i < len(lineWords) {
				outputString += lineWords[i] + " "
			}
		}
		output = append(output, outputString)
	}
	// print output
	for _, line := range output {
		fmt.Println(line)
	}
}

func input() []string {
	reader := bufio.NewReader(os.Stdin)

	var lines []string

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if len(strings.TrimSpace(line)) == 0 {
			return lines
		}
		lines = append(lines, line)
	}
}
