package main

import "fmt"

//Шаблон «Посетитель» — это способ отделения алгоритма от структуры объекта, в которой он оперирует.
//	Результат отделения — возможность добавлять новые операции в существующие структуры объектов без их
//модифицирования. Это один из способов соблюдения принципа открытости/закрытости (open/closed principle).
// все новые методы мы выделяем в отдельную структуру, которая уже реализует методы для всех структур, к которым применим
//наш метод. Получается , что метод вызывается путем передачи структуры метода в виде аргумента
// GeometryVisitor - интерфейс - определяюзий методы у структуры-метода которую мы создаем

type Geometry interface {
	Visit(GeometryVisitor) (interface{}, error)
}

type GeometryVisitor interface {
	VisitPoint(p *Point) (interface{}, error)
	VisitLine(l *Line) (interface{}, error)
	VisitCircle(c *Circle) (interface{}, error)
}

type Point struct {
	X, Y float32
}

func (point *Point) Visit(v GeometryVisitor) (interface{}, error) {
	return v.VisitPoint(point)
}

type Line struct {
	X1, Y1 float32
	X2, Y2 float32
}

func (line *Line) Visit(v GeometryVisitor) (interface{}, error) {
	return v.VisitLine(line)
}

type Circle struct {
	X, Y, R float32
}

func (circle *Circle) Visit(v GeometryVisitor) (interface{}, error) {
	return v.VisitCircle(circle)
}

type DistanceStrategy struct {
	X, Y float32
}

func (s *DistanceStrategy) VisitPoint(p *Point) (interface{}, error) {
	// Evaluate distance from point(X, Y) to point p
	return 10.1, nil
}

func (s *DistanceStrategy) VisitLine(l *Line) (interface{}, error) {
	// Evaluate distance from point(X, Y) to line l
	return 10.1, nil
}

func (s *DistanceStrategy) VisitCircle(c *Circle) (interface{}, error) {
	// Evaluate distance from point(X, Y) to circle c
	return 10.1, nil
}

func main() {
	s := &DistanceStrategy{X: 1, Y: 2}
	p := &Point{X: 3, Y: 4}
	res, err := p.Visit(s)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Distance is %g", res.(float64))
}
