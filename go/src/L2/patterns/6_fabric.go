package main

import (
	"fmt"
)

//Фабричный метод (Factory Method) - это паттерн, который определяет интерфейс для создания объектов некоторого класса,
//но непосредственное решение о том, объект какого класса создавать происходит в подклассах. То есть паттерн предполагает,
//что базовый класс делегирует создание объектов классам-наследникам.
//
//Когда надо применять паттерн
//Когда заранее неизвестно, объекты каких типов необходимо создавать
//
//Когда система должна быть независимой от процесса создания новых объектов и расширяемой: в нее можно легко вводить
//новые классы, объекты которых система должна создавать.
//
//Когда создание новых объектов необходимо делегировать из базового класса классам наследникам
// основная проблема - появляется божественный интерфейс Computer от которого мы потом не сможем отказаться
var (
	ServerType       = "server"
	PersonalComputer = "personal"
	NoteBookType     = "notebook"
)

type Computer interface {
	GetType() string
	PrintDetails()
}

func New(typeName string) Computer {
	switch typeName {
	case ServerType:
		return NewServer()
	case PersonalComputer:
		return NewPersonalComputer()
	case NoteBookType:
		return NewNoteBook()
	default:
		fmt.Printf("нет такого типа %s", typeName)
		return nil

	}
}

type Server struct {
	Type   string
	Core   int
	Memory string
}

func NewServer() Computer {
	serv := new(Server)
	serv.Type = "server"
	serv.Core = 16
	serv.Memory = "256gb"
	return serv
}

func (s Server) GetType() string {
	return s.Type
}

func (s Server) PrintDetails() {
	fmt.Printf("type = %s, core = %d, memory = %s\n", s.Type, s.Core, s.Memory)
}

type PersonalPC struct {
	Type    string
	Core    int
	Memory  string
	Monitor bool
}

func NewPersonalComputer() Computer {
	serv := new(PersonalPC)
	serv.Type = "personal"
	serv.Core = 4
	serv.Memory = "128gb"
	serv.Monitor = false
	return serv
}

func (s PersonalPC) GetType() string {
	return s.Type
}

func (s PersonalPC) PrintDetails() {
	fmt.Printf("type = %s, core = %d, memory = %s, monitor: %s\n", s.Type, s.Core, s.Memory, s.Monitor)
}

type NoteBook struct {
	Type    string
	Core    int
	Memory  string
	Monitor bool
}

func NewNoteBook() Computer {
	serv := new(NoteBook)
	serv.Type = "notebook"
	serv.Core = 4
	serv.Memory = "128gb"
	serv.Monitor = true
	return serv
}

func (s NoteBook) GetType() string {
	return s.Type
}

func (s NoteBook) PrintDetails() {
	fmt.Printf("type = %s, core = %d, memory = %s, monitor: %s\n", s.Type, s.Core, s.Memory, s.Monitor)
}

var types = []string{
	ServerType,
	PersonalComputer,
	NoteBookType,
	"a",
}

func main() {
	for _, t := range types {
		p := New(t)
		if p == nil {
			continue
		}
		fmt.Println("type = ", p.GetType())
		p.PrintDetails()
	}
}
