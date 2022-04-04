package main

import (
	"fmt"
	"time"
)

type Human struct {
	Name string
	Age int
	/*
	other field
	*/
}

func (h Human) GetAge() int {
	return h.Age
}

func (h Human) GetName() string {
	return h.Name
}

type Action struct {
	Human // встраивание(внедрение) структуры Human
	Time time.Time
}

func main() {

	action := Action{
		Human: Human{
			Name: "Nikita",
			Age: 15,
		},
		Time: time.Now(),
	}
	fmt.Printf("Action method GetAge: %d \n", action.GetAge())
	fmt.Printf("Action method GetName: %s \n", action.GetName())
	fmt.Printf("Action time: %v", action.Time)
}