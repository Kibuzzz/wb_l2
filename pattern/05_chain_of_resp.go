package pattern

import (
	"errors"
)

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern

Паттерн “Цепочка обязанностей” удобен тем, что позволяет передавать запрос через цепочку обработчиков до тех пор,
пока один из них не справится с задачей. Это упрощает добавление новых обработчиков и изменение логики обработки запросов,
поскольку все обработчики могут быть связаны в цепочку без необходимости изменения других частей системы.
Паттерн также делает код более гибким и расширяемым, так как позволяет легко изменять порядок обработки запросов
или добавлять новые шаги.

Однако у этого паттерна есть и минусы. Если в цепочке много обработчиков, запрос может долго перемещаться по цепочке,
что может снизить производительность. Кроме того, если цепочка не правильно настроена или не все обработчики реализуют
необходимую логику, запрос может “застрять” и не получить обработки. Это может усложнить отладку и поддержку кода,
особенно если обработчики зависят друг от друга.
*/

type httpRequest struct {
	validPassword bool
	restrictedIP  bool
}

type Handler interface {
	process(r httpRequest) error
	setNext(Handler)
}

type passwordValidator struct {
	next Handler
}

func (pv *passwordValidator) process(r httpRequest) error {
	if !r.validPassword {
		return errors.New("invalid password")
	}
	if pv.next != nil {
		return pv.next.process(r)
	}
	return nil
}

func (pv *passwordValidator) setNext(handler Handler) {
	pv.next = handler
}

type ipChecker struct {
	next Handler
}

func (ic *ipChecker) process(r httpRequest) error {
	if r.restrictedIP {
		return errors.New("restricted IP")
	}
	if ic.next != nil {
		return ic.next.process(r)
	}
	return nil
}

func (ic *ipChecker) setNext(handler Handler) {
	ic.next = handler
}

// Пример использования
// func main() {
// 	passwordHandler := &passwordValidator{}
// 	ipHandler := &ipChecker{}

// 	passwordHandler.setNext(ipHandler)

// 	request := httpRequest{validPassword: true, restrictedIP: false}

// 	err := passwordHandler.process(request)
// 	if err != nil {
// 		fmt.Println("Request failed:", err)
// 	} else {
// 		fmt.Println("Request succeeded")
// 	}
// }
