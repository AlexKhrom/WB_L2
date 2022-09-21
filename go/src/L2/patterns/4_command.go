package main

import "fmt"

//Паттерн "Команда" (Command) позволяет инкапсулировать запрос на выполнение определенного действия в виде отдельного
//объекта. Этот объект запроса на действие и называется командой. При этом объекты, инициирующие запросы на выполнение
//действия, отделяются от объектов, которые выполняют это действие.
//
//Команды могут использовать параметры, которые передают ассоциированную с командой информацию. Кроме того,
//команды могут ставиться в очередь и также могут быть отменены.
//
//Когда использовать команды?
//Когда надо передавать в качестве параметров определенные действия, вызываемые в ответ на другие действия. То есть
//когда необходимы функции обратного действия в ответ на определенные действия.
//
//Когда необходимо обеспечить выполнение очереди запросов, а также их возможную отмену.
//
//Когда надо поддерживать логгирование изменений в результате запросов. Использование логов может помочь восстановить
//состояние системы - для этого необходимо будет использовать последовательность запротоколированных команд.

type Command interface {
	execute()
}

type Restaurant struct {
	TotalDishes   int
	CleanedDishes int
}

// `NewRestaurant` constructs a new restaurant instance with 10 dishes,
// all of them being clean

func NewResteraunt() *Restaurant {
	const totalDishes = 10
	return &Restaurant{
		TotalDishes:   totalDishes,
		CleanedDishes: totalDishes,
	}
}

// The MakePizzaCommand is a struct which contains
// the number of pizzas to make, as well as the
// restaurant as its attributes
type MakePizzaCommand struct {
	n          int
	restaurant *Restaurant
}

func (c *MakePizzaCommand) execute() {
	// Reduce the total clean dishes of the restaurant
	// and print a message once done
	c.restaurant.CleanedDishes -= c.n
	fmt.Println("made", c.n, "pizzas")
}

// The MakeSaladCommand is similar to the MakePizza command
type MakeSaladCommand struct {
	n          int
	restaurant *Restaurant
}

func (c *MakeSaladCommand) execute() {
	c.restaurant.CleanedDishes -= c.n
	fmt.Println("made", c.n, "salads")
}

type CleanDishesCommand struct {
	restaurant *Restaurant
}

func (c *CleanDishesCommand) execute() {
	// Reset the cleaned dishes to the total dishes
	// present, and print a message once done
	c.restaurant.CleanedDishes = c.restaurant.TotalDishes
	fmt.Println("dishes cleaned")
}

func (r *Restaurant) MakePizza(n int) Command {
	return &MakePizzaCommand{
		restaurant: r,
		n:          n,
	}
}

func (r *Restaurant) MakeSalad(n int) Command {
	return &MakeSaladCommand{
		restaurant: r,
		n:          n,
	}
}

func (r *Restaurant) CleanDishes() Command {
	return &CleanDishesCommand{
		restaurant: r,
	}
}

type Cook struct {
	Commands []Command
}

// The executeCommands method executes all the commands
// one by one
func (c *Cook) executeCommands() {
	for _, c := range c.Commands {
		c.execute()
	}
}

func main() {
	// initialize a new resaurant
	r := NewResteraunt()

	// create the list of tasks to be executed
	tasks := []Command{
		r.MakePizza(2),
		r.MakeSalad(1),
		r.MakePizza(3),
		r.CleanDishes(),
		r.MakePizza(4),
		r.CleanDishes(),
	}

	// create the cooks that will execute the tasks
	cooks := []*Cook{
		&Cook{},
		&Cook{},
	}

	// Assign tasks to cooks alternating between the existing
	// cooks.
	for i, task := range tasks {
		// Using the modulus of the current task index, we can
		// alternate between different cooks
		cook := cooks[i%len(cooks)]
		cook.Commands = append(cook.Commands, task)
	}

	// Now that all the cooks have their commands, we can call
	// the `executeCommands` method that will have each cook
	// execute their respective commands
	for i, c := range cooks {
		fmt.Println("cook", i, ":")
		c.executeCommands()
	}
}
