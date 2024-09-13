Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
Будет выведено: 2 и 1.

Если отложенная(deffered) функция это литерал функции, а окружающая его функция имеет именнованное возвращаемое значение, отложенная функция может получить доступ и изменить параметр перед тем как вернет значение.

Из спецификации:
https://go.dev/ref/spec#Defer_statements
```
