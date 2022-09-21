package main

import (
	"fmt"
	"log"
)

//
//Когда поведение объекта должно зависеть от его состояния и может изменяться динамически во время выполнения
//
//Когда в коде методов объекта используются многочисленные условные конструкции, выбор которых зависит от текущего состояния объекта
//преимущества - избавляет от множества больших условных операторов-состояний
//недостатки - может неоправдано усложнить код, если состояний мало

type State interface {
	AddItem(int) error           // добавляем товары
	RequestItem() error          //
	InsertMoney(money int) error // получаем деньги от пользователя
	DispenseItem() error         // получаем товар
}

type VendingMachine struct {
	hasItem       State
	itemRequested State
	hasMoney      State
	noItem        State
	currentState  State
	itemCount     int
	itemPrice     int
}

func NewVendingMachine(itemCount, itemPrice int) *VendingMachine {
	v := &VendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}
	hasItemState := &HasItemState{
		vendingMachine: v,
	}
	itemRequestedState := &ItemRequestedState{
		vendingMachine: v,
	}
	hasMoneyState := &HasMoneyState{
		vendingMachine: v,
	}
	noItemState := &NoItemState{
		vendingMachine: v,
	}

	v.setState(hasItemState)
	v.hasItem = hasItemState
	v.itemRequested = itemRequestedState
	v.hasMoney = hasMoneyState
	v.noItem = noItemState
	return v
}

func (v *VendingMachine) RequestItem() error {
	return v.currentState.RequestItem()
}

func (v *VendingMachine) AddItem(count int) error {
	return v.currentState.AddItem(count)
}

func (v *VendingMachine) InsertMoney(money int) error {
	return v.currentState.InsertMoney(money)
}

func (v *VendingMachine) DispenseItem() error {
	return v.currentState.DispenseItem()
}

func (v *VendingMachine) setState(s State) {
	v.currentState = s
}

func (v *VendingMachine) incrementItemCount(count int) {
	v.itemCount += count
}

// нет товаров ===============================

type NoItemState struct {
	vendingMachine *VendingMachine
}

func (v *NoItemState) RequestItem() error {
	return fmt.Errorf("item out of stack")
}

func (v *NoItemState) AddItem(count int) error {
	v.vendingMachine.incrementItemCount(count)
	v.vendingMachine.setState(v.vendingMachine.hasItem)
	return nil
}

func (v *NoItemState) InsertMoney(money int) error {
	return fmt.Errorf("item out of stack")
}

func (v *NoItemState) DispenseItem() error {
	return fmt.Errorf("item out of stack")
}

// ===========================================

type ItemRequestedState struct {
	vendingMachine *VendingMachine
}

func (v *ItemRequestedState) RequestItem() error {
	return fmt.Errorf("item already requested")
}

func (v *ItemRequestedState) AddItem(count int) error {
	return fmt.Errorf("item dispense in progress")

}

func (v *ItemRequestedState) InsertMoney(money int) error {
	if money < v.vendingMachine.itemPrice {
		return fmt.Errorf("inserted money is less")
	}
	fmt.Println("money is ok")
	v.vendingMachine.setState(v.vendingMachine.hasMoney)
	return nil
}

func (v *ItemRequestedState) DispenseItem() error {
	return fmt.Errorf("please insert money")
}

//==============================================

type HasMoneyState struct {
	vendingMachine *VendingMachine
}

func (v *HasMoneyState) RequestItem() error {
	return fmt.Errorf("item dispense in progress")
}

func (v *HasMoneyState) AddItem(count int) error {
	return fmt.Errorf("item dispense in progress")

}

func (v *HasMoneyState) InsertMoney(money int) error {
	return fmt.Errorf("item out of stack")
}

func (v *HasMoneyState) DispenseItem() error {
	fmt.Println("DispenseItem")
	v.vendingMachine.itemCount = v.vendingMachine.itemCount - 1
	if v.vendingMachine.itemCount == 0 {
		v.vendingMachine.setState(v.vendingMachine.noItem)
	} else {
		v.vendingMachine.setState(v.vendingMachine.hasItem)
	}
	return nil
}

//==============================================

type HasItemState struct {
	vendingMachine *VendingMachine
}

func (v *HasItemState) RequestItem() error {
	if v.vendingMachine.itemCount == 0 {
		v.vendingMachine.setState(v.vendingMachine.noItem)
		return fmt.Errorf("No item Present")
	}
	fmt.Println("item requested")
	v.vendingMachine.setState(v.vendingMachine.itemRequested)
	return nil
}

func (v *HasItemState) AddItem(count int) error {
	fmt.Println("item added")
	v.vendingMachine.incrementItemCount(count)
	return nil
}

func (v *HasItemState) InsertMoney(money int) error {
	return fmt.Errorf("please select item first")
}

func (v *HasItemState) DispenseItem() error {
	return fmt.Errorf("please select item first")
}

func main() {
	vendingMachine := NewVendingMachine(1, 10)
	err := vendingMachine.RequestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = vendingMachine.InsertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = vendingMachine.DispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = vendingMachine.AddItem(2)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = vendingMachine.RequestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = vendingMachine.InsertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
