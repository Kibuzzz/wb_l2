Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Вывод:
```
Числа от 1 до 8 в случайном порядке и бесконечно много нулей
```

Ответ:
Программа создает две пишущие в канал горутины, после этого она объединяет эти два канала в один - читает данные из каналов a и b и пишет их в общий канал.

После того как канал a и b закрылись функция merge продолжает читать, но уже из пустых каналов и получает дефолтное значение 0.

Чтобы это исправить в select можно проверять закрыть ли канал или нет перед прочтением:
```go
func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	var closed1, closed2 bool
	go func() {
		for {
			if closed1 && closed2 {
				close(c)
				return
			}
			select {
			case v, open := <-a:
				if !open {
					closed1 = true
					continue
				}
				c <- v
			case v, open := <-b:
				if !open {
					closed2 = true
					continue
				}
				c <- v
			}
		}
	}()
	return c
}
```