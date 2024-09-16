package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func sortString(word string) string {
	word = strings.ToLower(word)
	s := strings.Split(word, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
func Anograms(words []string) map[string][]string {
	anograms := make(map[string][]string)
	ind := make(map[string]string)
	for _, word := range words {
		sorted := sortString(word)
		firstWord, ok := ind[sorted]
		if !ok {
			ind[sorted] = word
			anograms[word] = append(anograms[word], word)
		} else {
			anograms[firstWord] = append(anograms[firstWord], word)
		}
	}
	return anograms
}

func main() {
	input := []string{"пятка", "пятак", "тяпка", "листок", "слиток", "столик"}

	fmt.Println(Anograms(input))
}
