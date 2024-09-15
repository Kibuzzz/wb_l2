package pattern

import (
	"fmt"
)

/*
Паттерн “Фабричный метод” упрощает создание объектов разных типов, предоставляя единый интерфейс для их создания и управления.
Это делает код более гибким, поскольку добавление нового типа автомобиля возможно без изменения кода, который использует
эти автомобили. Такой подход позволяет легко расширять функциональность, создавая новые фабричные функции для разныхтипов машин.

Однако у этого паттерна есть и недостатки. С увеличением количества типов автомобилей может потребоваться создание множества
фабричных функций, что может привести к росту объема кода и усложнению его поддержки. Кроме того, важно следить за корректностью
имен и полей в коде, чтобы избежать ошибок и сохранить согласованность реализации методов.
*/
type CarInterface interface {
	setName(name string)
	setHorsePower(power int)
	getName() string
	getHorsePower() int
}

type car struct {
	name       string
	horsePower int
}

func (c *car) setName(name string) {
	c.name = name
}

func (c *car) getName() string {
	return c.name
}

func (c *car) setHorsePower(power int) {
	c.horsePower = power
}

func (c *car) getHorsePower() int {
	return c.horsePower
}

type mercedes struct {
	car
}

func newMercedes() CarInterface {
	return &mercedes{
		car: car{
			name:       "new mercedes",
			horsePower: 1200,
		},
	}
}

type bmw struct {
	car
}

func newBmw() CarInterface {
	return &bmw{
		car: car{
			name:       "bmw car",
			horsePower: 1000,
		},
	}
}

func getCar(carType string) (CarInterface, error) {
	if carType == "mercedes" {
		return newMercedes(), nil
	}
	if carType == "bmw" {
		return newBmw(), nil
	}
	return nil, fmt.Errorf("Wrong car type passed")
}

// func printDetails(c CarInterface) {
// 	fmt.Printf("Car Name: %s, Horse Power: %d\n", c.getName(), c.getHorsePower())
// }

// func main() {
// 	mercedes, _ := getCar("mercedes")
// 	bmw, _ := getCar("bmw")
// 	printDetails(mercedes)
// 	printDetails(bmw)
// }
