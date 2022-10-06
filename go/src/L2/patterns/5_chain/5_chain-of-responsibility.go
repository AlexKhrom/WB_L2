package chain

import "fmt"

//Когда применяется цепочка обязанностей?
//Когда имеется более одного объекта, который может обработать определенный запрос
//
//Когда надо передать запрос на выполнение одному из нескольких объект, точно не определяя, какому именно объекту
//
//Когда набор объектов задается динамически
//
//
//Ослабление связанности между объектами. Отправителю и получателю запроса ничего не известно друг о друге.
//	Клиенту неизветна цепочка объектов, какие именно объекты составляют ее, как запрос в ней передается.
//
//В цепочку с легкостью можно добавлять новые типы объектов, которые реализуют общий интерфейс.
//
//В то же время у паттерна есть недостаток: никто не гарантирует, что запрос в конечном счете будет обработан.
//	Если необходимого обработчика в цепочки не оказалось, то запрос просто выходит из цепочки и остается необработанным.

//каждое звено рещеает свою задачу и передает по звену дальше

type Data struct {
	OrdersIsCooked bool
	Order          []string
	Products       []string
}

type Service interface {
	Execute(data *Data)
	SetNext(service Service)
}

type MakeOrderService struct {
	Name string
	Next Service
}

func (s *MakeOrderService) Execute(data *Data) {
	fmt.Println("i want some eat :", data.Order)

	s.Next.Execute(data)
}

func (s *MakeOrderService) SetNext(service Service) {
	s.Next = service
}

type ShopService struct {
	Name string
	Next Service
}

func (s *ShopService) Execute(data *Data) {
	if data.Order == nil {
		fmt.Println("order shouldn't be empty")
		return
	}
	if len(data.Order) == 0 {
		fmt.Println("order shouldn't be empty")
		return
	}
	fmt.Println("buy some products")
	data.Products = append(data.Products, "cheese", "tomato")
	fmt.Println("products :", data.Products)
	s.Next.Execute(data)
}

func (s *ShopService) SetNext(service Service) {
	s.Next = service
}

type CookService struct {
	Name string
	Next Service
}

func (s *CookService) Execute(data *Data) {
	if data.Products == nil {
		fmt.Println("products shouldn't be empty")
		return
	}
	if len(data.Products) == 0 {
		fmt.Println("products shouldn't be empty")
		return
	}
	fmt.Println("cook order :", data.Order)
	data.OrdersIsCooked = true
	s.Next.Execute(data)
}

func (s *CookService) SetNext(service Service) {
	s.Next = service
}

type WaiterService struct {
	Name string
	Next Service
}

func (s *WaiterService) Execute(data *Data) {
	if !data.OrdersIsCooked {
		fmt.Println("some went wrong with cooking order")
		return
	}
	fmt.Println("bring order ", data.Order)
}

func (s *WaiterService) SetNext(service Service) {
	s.Next = service
}

func main() {

	orderService := &MakeOrderService{Name: "order service"}
	shopService := &ShopService{Name: "shop service"}
	cookService := &CookService{Name: "cook service"}
	waiterService := &WaiterService{Name: "waiter service"}

	orderService.SetNext(shopService)
	shopService.SetNext(cookService)
	cookService.SetNext(waiterService)

	data := Data{Order: []string{"spagetti"}, Products: []string{}}

	orderService.Execute(&data)
}
