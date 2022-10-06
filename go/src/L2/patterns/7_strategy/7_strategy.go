package strategy

import "fmt"

//
//Design patter Strategy  или шаблон проектирования Стратегия  относится к поведенческим шаблонам проектирования.
//	Его задача - выделить схожие алгоритмы, решающие конкретную задачу. Реализация алгоритмов выносится в отдельные классы
//и предоставляется возможность выбирать алгоритмы во время выполнения программы.
//
//Шаблон дает возможность в процессе выполнения выбрать стратегию (алгоритм, инструмент, подход) решения задачи.

type Strategy interface {
	Router(startPoint, endPoint int)
}

type Navigator struct {
	Strategy
}

func (nav *Navigator) SetStrategy(strategy Strategy) {
	nav.Strategy = strategy
}

type RoadStrategy struct {
}

func (r *RoadStrategy) Router(startPoint, endPoint int) {
	avgSpeed := 30
	trafficJam := 2
	total := endPoint - startPoint
	totalTime := total / avgSpeed * trafficJam
	fmt.Printf("a: %d, b : %d, avgSpeed: %d, trafficjam: %d, total: %d, totalTime:%d\n", startPoint, endPoint,
		avgSpeed, trafficJam, total, totalTime)
}

type PublicTransportStrategy struct {
}

func (r *PublicTransportStrategy) Router(startPoint, endPoint int) {
	avgSpeed := 40
	total := endPoint - startPoint
	totalTime := total / avgSpeed
	fmt.Printf("a: %d, b : %d, avgSpeed: %d, total: %d, totalTime:%d", startPoint, endPoint,
		avgSpeed, total, totalTime)
}

type WalkStrategy struct {
}

func (r *WalkStrategy) Router(startPoint, endPoint int) {
	avgSpeed := 4
	total := endPoint - startPoint
	totalTime := total * 60
	fmt.Printf("a: %d, b : %d, avgSpeed: %d, total: %d, totalTime:%d", startPoint, endPoint,
		avgSpeed, total, totalTime)
}

var (
	start     = 10
	end       = 100
	strategys = []Strategy{&PublicTransportStrategy{}, &RoadStrategy{}, &WalkStrategy{}}
)

func main() {
	nav := Navigator{}
	for _, strategy := range strategys {
		nav.SetStrategy(strategy)
		nav.Router(start, end)
	}
}
