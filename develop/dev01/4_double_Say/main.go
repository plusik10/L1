package main

import "fmt"

type Speaker interface {
	SayMyFullName()
}

type Human struct {
	name     string
	surename string
	age      int
}

func (h *Human) SayMyFullName() {
	h.say()
}

func (h *Human) say() {
	fmt.Println(h.name + " from human")
}

type Action struct {
	Human
}

func (a *Action) say() {
	fmt.Println(a.name + " from Action")
}

func main() {
	a := &Action{
		Human: Human{
			name:     "Ivan",
			surename: "Ivanovich",
			age:      10,
		},
	}

	a.SayMyFullName()
	fmt.Printf("%s %s %d", a.name, a.surename, a.age)
}
