/*
* Задача: «Система уведомлений»
Контекст:
Вы разрабатываете простую систему уведомлений. Система должна уметь отправлять сообщения через разные каналы (Email, SMS).
Техническое задание:
Интерфейс: Создайте интерфейс Notifier с методом Send(message string).
Структуры:
Создайте структуру EmailNotifier с полем EmailAddress.
Создайте структуру SMSNotifier с полем PhoneNumber.
Методы: Реализуйте метод Send для каждой структуры:
EmailNotifier.Send должен выводить в консоль: "Отправка Email на [адрес]: [сообщение]"
SMSNotifier.Send должен выводить в консоль: "Отправка SMS на номер [номер]: [сообщение]"
Основная функция (main):
Создайте слайс (slice) типа []Notifier.
Добавьте в него один экземпляр EmailNotifier и один SMSNotifier.
В цикле пройдитесь по слайсу и отправьте сообщение "Hello Go!" через каждый нотификатор.
Ожидаемый вывод:
Отправка Email на user@example.com: Hello Go!
Отправка SMS на номер +79990000000: Hello Go!
*/
package main

import "fmt"

type (
	Notifier interface {
		Send(message string)
	}

	EmailNotifier struct {
		EmailAddress string
	}

	SMSNotifier struct {
		PhoneNumber string
	}
)

func (e EmailNotifier) Send(message string) {
	fmt.Printf("Отправка на Email на %s: %s\n", e.EmailAddress, message)
}

func (s SMSNotifier) Send(message string) {
	fmt.Printf("Отправка на SMS на %s: %s\n", s.PhoneNumber, message)
}

func main() {
	s := []Notifier{}

	e := EmailNotifier{
		EmailAddress: "user@example.com",
	}

	sn := SMSNotifier{
		PhoneNumber: "+79990000000",
	}

	s = append(s, e, sn)

	for _, v := range s {
		v.Send("Hello Go!")
	}
}
