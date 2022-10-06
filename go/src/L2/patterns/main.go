package main

import (
	state "L2/patterns/8_state"
	"fmt"
)

func main() {
	// 1_facade =====================
	//var (
	//	bank = facade.Bank{
	//		Name:  "Банк",
	//		Cards: []facade.Card{},
	//	}
	//
	//	card1 = facade.Card{
	//		Name:    "card-1",
	//		Balance: 200,
	//		Bank:    &bank,
	//	}
	//
	//	card2 = facade.Card{
	//		Name:    "card-2",
	//		Balance: 5,
	//		Bank:    &bank,
	//	}
	//
	//	user1 = facade.User{
	//		Name: "user1",
	//		Card: &card1,
	//	}
	//
	//	user2 = facade.User{
	//		Name: "user2",
	//		Card: &card2,
	//	}
	//
	//	product = facade.Product{
	//		Name:  "prod",
	//		Price: 150,
	//	}
	//
	//	shop = facade.Shop{
	//		Name: "shop",
	//		Products: []facade.Product{
	//			product,
	//		},
	//	}
	//)
	//
	//bank.Cards = append(bank.Cards, card1, card2)
	//err := shop.Sell(user1, product.Name)
	//
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//
	//err = shop.Sell(user2, product.Name)
	//
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	// end 1_facade =================

	//2_builder ==============
	//comp := 2_builder.NewCompCollector().SetCore(1).SetBrand("asus").SetMemory(256).Get()
	//comp.Print()
	//fmt.Println()
	//end 2_builder =====================

	// 3_visitor ==================
	//s := &visitor.DistanceStrategy{X: 1, Y: 2}
	//p := &visitor.Point{X: 3, Y: 4}
	//res, err := p.Visit(s)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Distance is %g", res.(float64))
	// end 3_visitor ===============

	// commands =========
	//m := 4_command.NewMFC("arbat street")
	//tasks := []4_command.Command{
	//	m.MakePassport("alex", 20),
	//	m.PayTax("alex", 3567.2),
	//	m.MakePassport("alis", 22),
	//	m.PayTax("alis", 18765.4),
	//}
	//
	//exec := 4_command.ExecCommands{Commands: tasks}
	//exec.ExecuteCommands()
	// end commands ================

	// 5_chain ========================
	//orderService := &5_chain.MakeOrderService{Name: "order service"}
	//shopService := &5_chain.ShopService{Name: "shop service"}
	//cookService := &5_chain.CookService{Name: "cook service"}
	//waiterService := &5_chain.WaiterService{Name: "waiter service"}
	//
	//orderService.SetNext(shopService)
	//shopService.SetNext(cookService)
	//cookService.SetNext(waiterService)
	//
	//data := 5_chain.Data{Order: []string{"spagetti"}}
	//
	//orderService.Execute(&data)
	//end 5_chain =========================

	// factory ==================
	//var types = []string{
	//	factory.ServerType,
	//	factory.PersonalComputer,
	//	factory.NoteBookType,
	//	"a",
	//}
	//for _, t := range types {
	//	p := factory.New(t)
	//	if p == nil {
	//		continue
	//	}
	//	fmt.Println("type = ", p.GetType())
	//	p.PrintDetails()
	//}
	// end factory

	// strategy ==================
	//var (
	//	start     = 10
	//	end       = 100
	//	strategys = []strategy.Strategy{&strategy.PublicTransportStrategy{}, &strategy.RoadStrategy{}, &strategy.WalkStrategy{}}
	//)
	//
	//nav := strategy.Navigator{}
	//for _, strat := range strategys {
	//	nav.SetStrategy(strat)
	//	nav.Router(start, end)
	//}
	//end strategy

	// state ===================
	sh := state.NewInternetShop()
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
	// end state ========================
}
