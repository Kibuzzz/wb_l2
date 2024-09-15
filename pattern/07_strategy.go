package pattern

/*
Паттерн “Стратегия” предоставляет гибкость в выборе алгоритма или поведения во время выполнения, благодаря инкапсуляции
различных алгоритмов в отдельные классы. Это упрощает добавление новых стратегий и изменение существующих без модификации
клиентского кода, что делает систему более расширяемой и поддерживаемой. Например, если нужно изменить способ фильтрации
данных, можно просто создать новый класс стратегии и настроить его в контексте, не меняя остальной код.

Однако, у паттерна “Стратегия” есть и свои недостатки. Он может привести к увеличению числа классов в проекте, так как
для каждой новой стратегии требуется создать отдельный класс. Это может усложнить проект и сделать его более трудным
для понимания и поддержки, особенно если количество стратегий велико. Также, все стратегии должны реализовывать общий
интерфейс, что может быть ограничением, если алгоритмы сильно различаются по своему поведению или имеют разные требования.
*/

type Filter interface {
	Filter([]int) []int
}

type filterEven struct{}

func (f *filterEven) Filter(arr []int) []int {
	var result []int
	for _, v := range arr {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	return result
}

type filterOdd struct{}

func (f *filterOdd) Filter(arr []int) []int {
	var result []int
	for _, v := range arr {
		if v%2 != 0 {
			result = append(result, v)
		}
	}
	return result
}

type FilterContext struct {
	filter Filter
}

func (c *FilterContext) SetFilter(filter Filter) {
	c.filter = filter
}

func (c *FilterContext) ExecuteFilter(arr []int) []int {
	return c.filter.Filter(arr)
}

// func main() {
// 	arr := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}

// 	evenFilter := &filterEven{}
// 	oddFilter := &filterOdd{}

// 	context := &FilterContext{}

// 	context.SetFilter(evenFilter)
// 	evenFilteredArr := context.ExecuteFilter(arr)
// 	fmt.Println("Even Filtered Array:", evenFilteredArr)

// 	context.SetFilter(oddFilter)
// 	oddFilteredArr := context.ExecuteFilter(arr)
// 	fmt.Println("Odd Filtered Array:", oddFilteredArr)
// }
