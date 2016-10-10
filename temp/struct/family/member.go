package family

import "fmt"

type Member interface {
	Like()
}

type Dad struct {
	Name string
	Age  int
}

func (m *Dad) Like() {
	fmt.Printf("My name is %s. I'm %d years old. I like reading books.\n", m.Name, m.Age)
}

type Mom struct {
	Name string
	Age  int
}

func (m *Mom) Like() {
	fmt.Printf("My name is %s. I'm %d years old. I like watching Youtube.\n", m.Name, m.Age)
}

type Child struct {
	Name string
	Age  int
}

func (m *Child) Like() {
	fmt.Printf("My name is %s. I'm %d years old. I like playing with friends.\n", m.Name, m.Age)
}
