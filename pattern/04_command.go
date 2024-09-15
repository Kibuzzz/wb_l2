package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern

Паттерн “Команда” позволяет инкапсулировать запрос или действие в виде отдельного объекта. Это делает код гибким,
так как можно легко создавать команды, которые выполняют разные действия, и передавать их, как данные.
Такой подход особенно удобен, когда нужно сохранять, отменять или повторять действия, поскольку все команды хранятся
в одном месте. Например, в приложениях с кнопками “Отменить” или “Повторить” паттерн “Команда” помогает управлять
этими операциями простым и понятным образом.

Минус паттерна “Команда” в том, что он может привести к усложнению кода, особенно если в системе много различных команд.
Придется создавать отдельные классы для каждой команды, что может увеличить количество кода и привести к перегруженности.
Кроме того, если команды простые, использование этого паттерна может быть избыточным, поскольку для простых задач он
добавляет больше сложности, чем пользы.
*/

type Button struct {
	command Command
}

func (b *Button) press() {
	b.command.execute()
}

type Command interface {
	execute()
}

type OnCommand struct {
	device Device
}

func (c *OnCommand) execute() {
	c.device.on()
}

type OffCommand struct {
	device Device
}

func (c *OffCommand) execute() {
	c.device.off()
}

type Device interface {
	on()
	off()
}

type Tv struct {
	isRunning bool
}

func (t *Tv) on() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *Tv) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}

// Пример использования
// func main() {
//     tv := &Tv{}

//     onCommand := &OnCommand{
//         device: tv,
//     }

//     offCommand := &OffCommand{
//         device: tv,
//     }

//     onButton := &Button{
//         command: onCommand,
//     }
//     onButton.press()

//     offButton := &Button{
//         command: offCommand,
//     }
//     offButton.press()
// }
