package some

import "fmt"

type Some struct {
	Name string
}

func (s Some) Hello() {
	fmt.Println("name = ", s.Name)
}
