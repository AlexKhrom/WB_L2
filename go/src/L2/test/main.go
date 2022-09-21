package main

import "fmt"

type A interface {
	hello()
}

type AStruct struct {
}

func (s *AStruct) hello() {
	fmt.Println("hello!")
}

func hi(inter A) {
	fmt.Println("hi")
}

func main() {
	var s AStruct
	hi(&s)
}
