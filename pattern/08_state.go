package pattern

import (
	"fmt"
)

/*
Паттерн “Состояние” обеспечивает гибкость и чистоту кода, позволяя объекту изменять свое поведение в зависимости от
внутреннего состояния. Это улучшает организацию кода, устраняя громоздкие условные конструкции и делая логику
состояний более прозрачной. Паттерн упрощает добавление новых состояний или изменение существующих, так как каждое
состояние инкапсулируется в отдельный класс, что минимизирует изменения в основном коде и улучшает его поддерживаемость.

Однако, использование паттерна “Состояние” может привести к увеличению количества классов в проекте, что усложняет
его структуру и может затруднить понимание и поддержку кода. Если количество состояний велико, управление переходами
между ними может стать сложным и трудоемким, особенно если логика переходов становится запутанной. Это может потребовать
тщательной документации и дополнительных усилий для обеспечения правильного функционирования всех состояний и переходов.
*/

// State интерфейс для различных состояний автомата
type State interface {
	InsertMoney(context *VendingMachine)
	SelectProduct(context *VendingMachine)
	Dispense(context *VendingMachine)
}

// Конкретное состояние: Ожидание монет
type WaitingForMoneyState struct{}

func (s *WaitingForMoneyState) InsertMoney(context *VendingMachine) {
	fmt.Println("Money inserted.")
	context.SetState(&ProductSelectedState{}) // Переход к следующему состоянию
}

func (s *WaitingForMoneyState) SelectProduct(context *VendingMachine) {
	fmt.Println("Insert money first.")
}

func (s *WaitingForMoneyState) Dispense(context *VendingMachine) {
	fmt.Println("Insert money first.")
}

// Конкретное состояние: Выбор продукта
type ProductSelectedState struct{}

func (s *ProductSelectedState) InsertMoney(context *VendingMachine) {
	fmt.Println("Money already inserted.")
}

func (s *ProductSelectedState) SelectProduct(context *VendingMachine) {
	fmt.Println("Product selected.")
	context.SetState(&DispensingState{}) // Переход к следующему состоянию
}

func (s *ProductSelectedState) Dispense(context *VendingMachine) {
	fmt.Println("Select a product first.")
}

// Конкретное состояние: Выдача продукта
type DispensingState struct{}

func (s *DispensingState) InsertMoney(context *VendingMachine) {
	fmt.Println("Please wait, product is being dispensed.")
}

func (s *DispensingState) SelectProduct(context *VendingMachine) {
	fmt.Println("Product already selected.")
}

func (s *DispensingState) Dispense(context *VendingMachine) {
	fmt.Println("Dispensing product.")
	context.SetState(&FinishedState{}) // Завершение процесса
}

// Конкретное состояние: Завершение
type FinishedState struct{}

func (s *FinishedState) InsertMoney(context *VendingMachine) {
	fmt.Println("Transaction already completed.")
}

func (s *FinishedState) SelectProduct(context *VendingMachine) {
	fmt.Println("Transaction already completed.")
}

func (s *FinishedState) Dispense(context *VendingMachine) {
	fmt.Println("Transaction already completed.")
}

// Контекст: Автомат по продаже напитков
type VendingMachine struct {
	state State
}

func (c *VendingMachine) SetState(state State) {
	c.state = state
}

func (c *VendingMachine) InsertMoney() {
	c.state.InsertMoney(c)
}

func (c *VendingMachine) SelectProduct() {
	c.state.SelectProduct(c)
}

func (c *VendingMachine) Dispense() {
	c.state.Dispense(c)
}

// Пример использования
// func main() {
// 	vm := &VendingMachine{}

// 	vm.SetState(&WaitingForMoneyState{})

// 	vm.InsertMoney()
// 	vm.SelectProduct()
// 	vm.Dispense()
// 	vm.InsertMoney()   // Уже завершено
// 	vm.SelectProduct() // Уже завершено
// 	vm.Dispense()      // Уже завершено
// }
