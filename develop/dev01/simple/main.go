package main

import "fmt"

type Human struct {
	name     string
	surename string
	age      int
}

func (h *Human) SayMyFullName() {
	fmt.Printf("my name is %s %s, my age: %d\n", h.name, h.surename, h.age)
}

type Action struct {
	Human
}

func main() {
	a := Action{
		Human: Human{
			name:     "Ivan",
			surename: "Ivanovich",
			age:      10,
		},
	}

	a.SayMyFullName()
	fmt.Printf("%s %s %d", a.name, a.surename, a.age)
}
