package state

import (
	"fmt"
)

//
//Когда поведение объекта должно зависеть от его состояния и может изменяться динамически во время выполнения
//
//Когда в коде методов объекта используются многочисленные условные конструкции, выбор которых зависит от текущего состояния объекта
//преимущества - избавляет от множества больших условных операторов-состояний
//недостатки - может неоправдано усложнить код, если состояний мало

// состояние покупки товара
// сотояния :
// добавление товара на склад ->
// адрес доставки - регистрация заказа ->
// оплата заказа ->
// доставляем заказ ->

type State interface {
	AddItem(int) error
	RegistrationOrder(address string) error
	PayOrder(float64) error
	DeliverOrder() error
}

type InternetShop struct {
	HasItem      State
	NoItem       State
	Registr      State
	HasMoney     State
	CurrentState State
	money        float64
	address      string
	items        int
}

func (sh *InternetShop) SetState(s State) {
	sh.CurrentState = s
}

func NewInternetShop() *InternetShop {
	sh := &InternetShop{}
	hasItem := &HasItemState{shop: sh}
	NoItem := &NoItemState{shop: sh}
	Registr := &RegistrState{shop: sh}
	HasMoney := &HasMoneyState{shop: sh}

	sh.SetState(NoItem)
	sh.HasMoney = HasMoney
	sh.NoItem = NoItem
	sh.HasItem = hasItem
	sh.Registr = Registr

	return sh
}

func (sh *InternetShop) AddItem(count int) error {
	return sh.CurrentState.AddItem(count)
}

func (sh *InternetShop) RegistrationOrder(address string) error {
	return sh.CurrentState.RegistrationOrder(address)
}

func (sh *InternetShop) PayOrder(money float64) error {
	return sh.CurrentState.PayOrder(money)
}

func (sh *InternetShop) DeliverOrder() error {
	return sh.CurrentState.DeliverOrder()
}

type NoItemState struct {
	shop *InternetShop
}

func (sh *NoItemState) AddItem(count int) error {
	sh.shop.items += count
	sh.shop.SetState(sh.shop.HasItem)
	fmt.Println("items added")
	return nil
}

func (sh *NoItemState) RegistrationOrder(address string) error {
	return fmt.Errorf("item out of stack")
}

func (sh *NoItemState) PayOrder(money float64) error {
	return fmt.Errorf("item out of stack")
}

func (sh *NoItemState) DeliverOrder() error {
	return fmt.Errorf("item out of stack")
}

type HasItemState struct {
	shop *InternetShop
}

func (sh *HasItemState) AddItem(count int) error {
	sh.shop.items += count
	fmt.Println("items added")
	return nil
}

func (sh *HasItemState) RegistrationOrder(address string) error {
	if sh.shop.items == 0 {
		sh.shop.SetState(sh.shop.NoItem)
		return fmt.Errorf("No item Present")
	}
	sh.shop.address = address
	sh.shop.SetState(sh.shop.Registr)
	fmt.Println("item was registr")
	return nil
}

func (sh *HasItemState) PayOrder(money float64) error {
	return fmt.Errorf("item not registr")
}

func (sh *HasItemState) DeliverOrder() error {
	return fmt.Errorf("item not registr")
}

type RegistrState struct {
	shop *InternetShop
}

func (sh *RegistrState) AddItem(count int) error {
	return fmt.Errorf("process for registration item")
}

func (sh *RegistrState) RegistrationOrder(address string) error {
	return fmt.Errorf("item already registration")
}

func (sh *RegistrState) PayOrder(money float64) error {
	if money < 1000 {
		return fmt.Errorf("not enough money")
	}
	sh.shop.SetState(sh.shop.HasMoney)
	return nil
}

func (sh *RegistrState) DeliverOrder() error {
	return fmt.Errorf("item not payed")
}

type HasMoneyState struct {
	shop *InternetShop
}

func (sh *HasMoneyState) AddItem(count int) error {
	return fmt.Errorf("process for deliver item")
}

func (sh *HasMoneyState) RegistrationOrder(address string) error {
	return fmt.Errorf("process for deliver item")
}

func (sh *HasMoneyState) PayOrder(money float64) error {
	return fmt.Errorf("item already payd")
}

func (sh *HasMoneyState) DeliverOrder() error {
	fmt.Println("item delivering")
	sh.shop.items--
	if sh.shop.items == 0 {
		sh.shop.SetState(sh.shop.NoItem)
	} else {
		sh.shop.SetState(sh.shop.HasItem)
	}
	return nil
}

func main() {
	sh := NewInternetShop()
	//vendingMachine := NewVendingMachine(1, 10)
	err := sh.RegistrationOrder("arbat")
	if err != nil {
		fmt.Println("err: ", err)
	}
	err = sh.PayOrder(10)
	if err != nil {
		fmt.Println("err: ", err)
	}
	err = sh.DeliverOrder()
	if err != nil {
		fmt.Println("err: ", err)
	}
	err = sh.AddItem(2)
	if err != nil {
		fmt.Println("err: ", err)
	}
	err = sh.RegistrationOrder("arbat")
	if err != nil {
		fmt.Println("err: ", err)
	}
	err = sh.PayOrder(10001)
	if err != nil {
		fmt.Println("err: ", err)
	}
	err = sh.DeliverOrder()
	if err != nil {
		fmt.Println("err: ", err)
	}
}
