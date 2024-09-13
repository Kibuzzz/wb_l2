Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Вывод:
```
error
```
Ответ:

Структура интерфейса:
```go
type iface struct {
    tab  *itab           // информация о типе, реализующем интерфейс
    data unsafe.Pointer  // указатель на данные типа, реализующего интерфейс
}
```
Когда интерфейс сравнивается с nil, обе части (и tab, и data) должны быть nil для того, чтобы интерфейс считался nil.

В функции main err != nil потому что когда из функции test возвращается *customError равный nil, пусть значение и равно нил, но тип указан, поэтому интерфейс не равен nil
