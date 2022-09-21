package main

import "fmt"

//Строитель (Builder) - шаблон проектирования, который инкапсулирует создание объекта и позволяет разделить его на различные этапы.
//
//Когда использовать паттерн Строитель?
//Когда процесс создания нового объекта не должен зависеть от того, из каких частей этот объект состоит и как эти части связаны между собой
//
//Когда необходимо обеспечить получение различных вариаций объекта в процессе его создания

//мы хотим создать компьютер определенной модели, для каждой модели у нас есть свой сборщик, который реализует интерфейс сборщика для компьютера
//создаем объект конкретного сборщика, у всех сборщиков один и тот же интерфейс

type Collector interface {
	SetCore()
	SetBrand()
	SetMemory()
	SetMonitor()
	SetGraphicCard()
	GetComputer() Comp
}

type Comp struct {
	Core        int
	Brand       string
	Memory      int
	Monitor     int
	GraphicCard int
}

func (c Comp) print() {
	fmt.Printf("brand: %s, core:%d, memory: %d,monitor:%d ,graphicalCard: %d\n", c.Brand, c.Core, c.Memory, c.Monitor, c.GraphicCard)
}

const (
	AsusCollectorType = "asus"
	HPCollectorType   = "hp"
)

func GetCollector(collectorType string) Collector {
	switch collectorType {
	case AsusCollectorType:
		return &AsusCollector{}
	case HPCollectorType:
		return &AsusCollector{}
	default:
		return nil
	}
}

type AsusCollector struct {
	Core        int
	Brand       string
	Memory      int
	Monitor     int
	GraphicCard int
}

func (collector *AsusCollector) SetCore() {
	collector.Core = 4
}
func (collector *AsusCollector) SetBrand() {
	collector.Brand = "asus"
}
func (collector *AsusCollector) SetMemory() {
	collector.Memory = 8
}
func (collector *AsusCollector) SetMonitor() {
	collector.Monitor = 1
}
func (collector *AsusCollector) SetGraphicCard() {
	collector.GraphicCard = 1
}
func (collector *AsusCollector) GetComputer() Comp {
	return Comp{
		Core:        collector.Core,
		Brand:       collector.Brand,
		Memory:      collector.Memory,
		Monitor:     collector.Monitor,
		GraphicCard: collector.GraphicCard,
	}
}

type HpCollector struct {
	Core        int
	Brand       string
	Memory      int
	Monitor     int
	GraphicCard int
}

func (collector *HpCollector) SetCore() {
	collector.Core = 4
}
func (collector *HpCollector) SetBrand() {
	collector.Brand = "hp"
}
func (collector *HpCollector) SetMemory() {
	collector.Memory = 16
}
func (collector *HpCollector) SetMonitor() {
	collector.Monitor = 2
}
func (collector *HpCollector) SetGraphicCard() {
	collector.GraphicCard = 1
}
func (collector *HpCollector) GetComputer() Comp {
	return Comp{
		Core:        collector.Core,
		Brand:       collector.Brand,
		Memory:      collector.Memory,
		Monitor:     collector.Monitor,
		GraphicCard: collector.GraphicCard,
	}
}

type Factory struct {
	collector Collector
}

func NewFactory(collector Collector) *Factory {
	return &Factory{collector: collector}
}

func (factory *Factory) SetCollector(collector Collector) {
	factory.collector = collector
}

func (factory *Factory) CreatComputer() Comp {
	factory.collector.SetCore()
	factory.collector.SetBrand()
	factory.collector.SetMemory()
	factory.collector.SetMonitor()
	factory.collector.SetGraphicCard()
	return factory.collector.GetComputer()
}

func main() {

}
