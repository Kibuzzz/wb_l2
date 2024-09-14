package pattern

import "fmt"

/*
Паттерн “Фасад” используется для предоставления удобного и простого интерфейса, который скрывает сложную внутреннюю
логику подсистем. “Фасад” может не реализовывать весь функционал подсистем, которые в него входят, но предоставляет
только необходимые методы для взаимодействия с ними.

Плюсы:
  -  Простой интерфейс.
  -  Улучшенная читаемость и поддерживаемость кода.
  -  Снижение связанности между клиентом и сложными подсистемами.
Минусы:
  -  Может стать слишком большим, если система расширяется.
  -  Усложнение отладки из-за скрытия деталей работы подсистем.
  -  Добавление фасада увеличивает количество классов, что может привести к излишнему
     усложнению структуры приложения.
  -  Может вызвать дублирование функционала, если для разных клиентов нужны разные интерфейсы,
     что приведет к созданию дополнительных фасадов.
*/

// система для обработки транзакций
type transactionProvider struct {
}

func (tp *transactionProvider) processTransaction(amount float64) {
	fmt.Printf("Процессинг транзакции на сумму %.2f\n", amount)
}

// система для ввода и вывода денег из банкомата
type ioSystem struct {
}

func (io *ioSystem) withdrawCash(amount float64) {
	fmt.Printf("Выдача наличных: %.2f\n", amount)
}

func (io *ioSystem) depositCash(amount float64) {
	fmt.Printf("Приём наличных: %.2f\n", amount)
}

// система проверки карты
type cardValidator struct {
}

func (cv *cardValidator) validateCard(cardNumber string) bool {
	fmt.Printf("Проверка карты: %s\n", cardNumber)
	// Допустим, карта всегда валидна для упрощения
	return true
}

// Фасад банкомата, который упрощает работу с подсистемами
type ATMFacade struct {
	transaction transactionProvider
	io          ioSystem
	validator   cardValidator
}

// Метод снятия наличных через фасад
func (atm *ATMFacade) Withdraw(cardNumber string, amount float64) {
	if atm.validator.validateCard(cardNumber) {
		atm.transaction.processTransaction(amount)
		atm.io.withdrawCash(amount)
	}
}

// Метод пополнения счета через фасад
func (atm *ATMFacade) Deposit(cardNumber string, amount float64) {
	if atm.validator.validateCard(cardNumber) {
		atm.transaction.processTransaction(amount)
		atm.io.depositCash(amount)
	}
}

// Пример использования
// func main() {
// 	atm := ATMFacade{
// 		transaction: transactionProvider{},
// 		io:          ioSystem{},
// 		validator:   cardValidator{},
// 	}

// 	atm.Withdraw("1234-5678-9012", 100.0)
// 	atm.Deposit("1234-5678-9012", 50.0)
// }
