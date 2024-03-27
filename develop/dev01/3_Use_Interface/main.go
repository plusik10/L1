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
	fmt.Printf("my name is %s %s, my age: %d\n", h.name, h.surename, h.age)
}

type Action struct {
	Human
}

func Speak(s Speaker) {
	s.SayMyFullName()
}

func main() {
	a := &Action{
		Human: Human{
			name:     "Ivan",
			surename: "Ivanovich",
			age:      10,
		},
	}

	Speak(a)
	fmt.Printf("%s %s %d", a.name, a.surename, a.age)
}
