package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern

Паттерн “Посетитель” удобен, когда нужно добавлять новые функции или действия для разных объектов (например, фигур),
не меняя сами эти объекты. Ты можешь создать отдельные классы, которые описывают, как выполнять разные операции над объектами,
и это помогает поддерживать код более чистым и легко расширяемым. Например, если нужно вычислить площадь или периметр для фигур,
достаточно создать отдельные “посетители” для каждой операции, и не нужно менять код самих фигур. Это делает добавление новых
операций очень простым.

Однако у “Посетителя” есть и минусы. Если у тебя появятся новые типы объектов (например, новая фигура — треугольник),
придется изменить всех существующих посетителей, чтобы они знали, как работать с этой новой фигурой. Это может быть неудобно,
если структура часто меняется или добавляются новые типы объектов, так как придется много кода менять в разных местах.
*/

type Shape interface {
	getType() string
	accept(Visitor)
}

type Square struct {
	side int
}

func (s *Square) accept(v Visitor) {
	v.visitForSquare(s)
}

func (s *Square) getType() string {
	return "Square"
}

type Circle struct {
	radius int
}

func (c *Circle) accept(v Visitor) {
	v.visitForCircle(c)
}

func (c *Circle) getType() string {
	return "Circle"
}

type Visitor interface {
	visitForSquare(*Square)
	visitForCircle(*Circle)
}

type AreaCalculator struct {
	area int
}

func (a *AreaCalculator) visitForSquare(s *Square) {
	// Calculate area for square.
	// Then assign in to the area instance variable.
	fmt.Println("Calculating area for square")
}

func (a *AreaCalculator) visitForCircle(s *Circle) {
	fmt.Println("Calculating area for circle")
}
