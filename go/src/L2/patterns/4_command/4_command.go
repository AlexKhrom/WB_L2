package command

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

type ExecCommands struct {
	Commands []Command
}

// The executeCommands method executes all the commands
// one by one

func (e *ExecCommands) ExecuteCommands() {
	for i, c := range e.Commands {
		fmt.Println("exec 4_command ", i+1)
		c.execute()
	}
}

type MFC struct {
	Adress       string
	TotalClients int
	Workers      int
}

func NewMFC(adress string) *MFC {
	return &MFC{
		Adress:       adress,
		TotalClients: 0,
		Workers:      3,
	}
}

func (m *MFC) MakePassport(name string, age int) Command {
	return &MakePassportCommand{
		MFC:  m,
		Age:  age,
		Name: name,
	}
}

func (m *MFC) PayTax(name string, val float64) Command {
	return &PayTaxCommand{
		MFC:   m,
		Name:  name,
		Value: val,
	}
}

type MakePassportCommand struct {
	MFC  *MFC
	Name string
	Age  int
}

func (m *MakePassportCommand) execute() {
	fmt.Printf("make passport for %s with %d years\n", m.Name, m.Age)
}

type PayTaxCommand struct {
	MFC   *MFC
	Name  string
	Value float64
}

func (p *PayTaxCommand) execute() {
	fmt.Printf("%s paid a tax of %f rubles\n", p.Name, p.Value)
}

func main() {
	m := NewMFC("arbat street")
	tasks := []Command{
		m.MakePassport("alex", 20),
		m.PayTax("alex", 3567.2),
		m.MakePassport("alis", 22),
		m.PayTax("alis", 18765.4),
	}

	for _, t := range tasks {
		t.execute()
	}
}
