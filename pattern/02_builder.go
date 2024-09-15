package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern

Паттерн “Строитель” предоставляет гибкий способ создания сложных объектов, разделяя процесс их создания на этапы.
Это делает код более читабельным и понятным, особенно если объект имеет много параметров или настройки.
“Строитель” позволяет удобно настраивать объект пошагово, без необходимости передавать множество параметров в конструктор,
а также делает возможным создание разных вариаций объекта с различными комбинациями параметров.
Этот подход облегчает тестирование и поддержку кода, так как упрощает добавление новых свойств и улучшений.

Однако у паттерна “Строитель” есть и минусы. Он может усложнять код, добавляя новые классы и методы,
что может быть избыточным для простых объектов с малым количеством параметров. Кроме того, создание цепочек
методов может привести к большому количеству кода, если объект имеет множество свойств. Это также может увеличить
вероятность ошибок, если методы используются неправильно, так как порядок вызова не всегда очевиден.
*/

type robot struct {
	color  string
	name   string
	power  float64
	armor  float64
	hasHat bool
}

type robotBuilder struct {
	robot robot
}

func NewRobotBuilder() *robotBuilder {
	return &robotBuilder{robot: robot{}}
}

func (rb *robotBuilder) SetName(name string) *robotBuilder {
	rb.robot.name = name
	return rb
}

func (rb *robotBuilder) SetColor(color string) *robotBuilder {
	rb.robot.color = color
	return rb
}

func (rb *robotBuilder) SetPower(power float64) *robotBuilder {
	rb.robot.power = power
	return rb
}

func (rb *robotBuilder) SetArmor(armor float64) *robotBuilder {
	rb.robot.armor = armor
	return rb
}

func (rb *robotBuilder) SetHasHat(hasHat bool) *robotBuilder {
	rb.robot.hasHat = hasHat
	return rb
}

func (rb *robotBuilder) Build() robot {
	return rb.robot
}
