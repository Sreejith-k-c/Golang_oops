package main

import "fmt"

type Engine interface {
	Start()
}

type PetrolEngine struct {
	Name string
}

func (e PetrolEngine) Start() {
	fmt.Println("starting...", e.Name)
}

type gasEngine struct {
	Name string
}

func (e gasEngine) Start() {
	fmt.Println("starting...", e.Name)
}
func Start(e Engine) {
	e.Start()
}
func main() {
	fmt.Println("Interfaces")
	pEngine := PetrolEngine{
		Name: "Petrol",
	}
	pEngine.Start()

	fmt.Println("Interfaces")
	gEngine := gasEngine{
		Name: "Gas",
	}
	Start(pEngine)
	Start(gEngine)
}
