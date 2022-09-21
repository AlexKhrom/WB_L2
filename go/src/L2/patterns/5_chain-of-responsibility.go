package main

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

type Service interface {
	Execute(*Data)
	SetNext(next Service)
}

type Data struct {
	GetSource    bool
	UpdateSource bool
}

type Device struct {
	Name string
	Next Service
}

func (d *Device) Execute(data *Data) {
	if data.GetSource {
		fmt.Println("data from device already get")
		d.Next.Execute(data)
	}
	fmt.Println("get data from device")
	data.GetSource = true
	d.Next.Execute(data)
}

func (d *Device) SetNext(next Service) {
	d.Next = next
}

// ================================

type UpdateDataService struct {
	Name string
	Next Service
}

func (d *UpdateDataService) Execute(data *Data) {
	if data.UpdateSource {
		fmt.Println("data from device already update")
		d.Next.Execute(data)
	}
	fmt.Println("update data from device")
	data.UpdateSource = true
	d.Next.Execute(data)
}

func (d *UpdateDataService) SetNext(next Service) {
	d.Next = next
}

//===============================

type DataService struct {
	Name string
	Next Service
}

func (d *DataService) Execute(data *Data) {
	if !data.UpdateSource {
		fmt.Println("data not update")
		return
	}
	fmt.Println("data save")
}

func (d *DataService) SetNext(next Service) {
	d.Next = next
}

func main() {
	device1 := &Device{Name: "device 1"}
	updateService := &UpdateDataService{Name: "update 1"}
	dataService := &DataService{}
	device1.SetNext(updateService)
	updateService.SetNext(dataService)
	data := &Data{}
	device1.Execute(data)
}
