package facade

import (
	"errors"
	"fmt"
	"time"
)

// Паттерн Фасад
// Клиент – использует фасад вместо прямой работы с объектами сложной подсистемы.
//
//Фасад. Применимость.
//
//Когда нужно представить простой или урезанный интерфейс к сложной подсистеме.
//Когда надо уменьшить количество зависимостей между клиентом и сложной системой. Фасадные объекты позволяют отделить,
//изолировать компоненты системы от клиента и развивать и работать с ними независимо.
//Когда вы хотите разложить подсистему на отдельные слои.
//
//Фасад. Преимущества и недостатки.
//
//Преимущество: изолирует клиентов от компонентов сложной подсистемы.
//Недостаток: фасад рискует стать божественным объектом, привязанным ко всем классам программы.

type Product struct {
	Name  string
	Price float64
}

type Shop struct {
	Name     string
	Products []Product
}

func (sh Shop) Sell(user User, product string) error {
	fmt.Println("[Магазин] запрос к пользователю для получения остатка по карте")
	time.Sleep(time.Second)
	err := user.Card.CheckBalance()
	if err != nil {
		return err
	}
	fmt.Printf("[Магазин] проверка может ли пользователь %s купить товар\n", user.Name)

	for _, prod := range sh.Products {
		if prod.Name != product {
			continue
		}
		if prod.Price > user.Card.Balance {
			return errors.New("у пользователя недостаточно средств для покупки товара")
		}
	}
	fmt.Printf("[Магазин] пользователь %s купил товар\n", user.Name)
	return nil
}

type Bank struct {
	Name  string
	Cards []Card
}

func (bank Bank) CheckBalance(cardNumber string) error {
	fmt.Println("[Банк] получение баланса по карте")
	time.Sleep(time.Second)

	for _, card := range bank.Cards {
		if card.Name != cardNumber {
			continue
		}
		if card.Balance <= 0 {
			return errors.New("[Банк] недостаточно средств")
		}
	}
	fmt.Println("[Банк] остаток положительный")
	return nil
}

type Card struct {
	Name    string
	Balance float64
	Bank    *Bank
}

func (card Card) CheckBalance() error {
	fmt.Println("[Карта] запрос в банк для проверки остатка")
	time.Sleep(time.Second)
	return card.Bank.CheckBalance(card.Name)

}

type User struct {
	Name string
	Card *Card
}

func (u *User) GetBalance() float64 {
	return u.Card.Balance
}

var (
	bank = Bank{
		Name:  "Банк",
		Cards: []Card{},
	}

	card1 = Card{
		Name:    "card-1",
		Balance: 200,
		Bank:    &bank,
	}

	card2 = Card{
		Name:    "card-2",
		Balance: 5,
		Bank:    &bank,
	}

	user1 = User{
		Name: "user1",
		Card: &card1,
	}

	user2 = User{
		Name: "user2",
		Card: &card2,
	}

	product = Product{
		Name:  "prod",
		Price: 150,
	}

	shop = Shop{
		Name: "shop",
		Products: []Product{
			product,
		},
	}
)

func main() {
	bank.Cards = append(bank.Cards, card1, card2)
	err := shop.Sell(user1, product.Name)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = shop.Sell(user2, product.Name)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

}
