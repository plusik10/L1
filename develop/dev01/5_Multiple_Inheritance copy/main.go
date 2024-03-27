package main

import "fmt"

type Speaker interface {
	Speak()
}

type Human struct {
	name string
}

func (h *Human) Speak() {
	fmt.Println("i'm human and i'm speak")
}

type Manager interface {
	Speaker
	Lead()
}

type worker interface {
	Speaker
	Work()
}
type manager struct {
	Human
}

func (m *manager) Lead() {
	fmt.Println("I'm manager and i'm leader")
}

type Worker struct {
	Human
}

func (w *Worker) Work() {
	fmt.Println("I'm worker's")
}

type ingener struct {
	Worker
}

type Director struct {
	manager
}

func checkSpeaker(s Speaker) {
	s.Speak()
}

func checkManager(m Manager) {
	m.Lead()
}

func checkWorker(w worker) {
	w.Work()
}

func main() {
	ingeneer := &ingener{}
	checkWorker(ingeneer)
	checkSpeaker(ingeneer)
	director := &Director{}
	checkManager(director)
	checkSpeaker(director)
}
