package main

import "fmt"

type Speaker interface {
	say()
}

type Runer interface {
	run()
}

type Human struct {
	name     string
	surename string
	age      int
}

type Dog struct {
	nickname string
	breed    string
}

func (b *Human) say() {
	fmt.Println("Say func")
}

func (b *Human) run() {
	fmt.Println("Run func")
}

type Action struct {
	Human // встраиваем
	Dog   // встраиваем
}

func speak(b Speaker) {
	b.say()
}

func walk(b Runer) {
	b.run()
}

func main() {
	a := &Action{
		Human: Human{
			name:     "Ivan",
			surename: "Ivanovich",
			age:      10,
		},
		Dog: Dog{
			nickname: "pesik",
			breed:    "pudel",
		},
	}
	a.say()
	a.run()

	speak(a)
	walk(a)
}
