package main

import (
	"fmt"
)

type AnimalStruct struct {
	food       string
	locomotion string
	noise      string
}

type animal interface {
	Eat()
	Move()
	Speak()
}

var animalMap = map[string]AnimalStruct{
	"cow":   {"grass", "walk", "moo"},
	"bird":  {"worms", "fly", "peep"},
	"snake": {"mice", "slither", "hsss"},
}

func (animal AnimalStruct) Eat() {
	fmt.Println(animal.food)
	return
}

func (animal AnimalStruct) Move() {
	fmt.Println(animal.locomotion)
	return
}

func (animal AnimalStruct) Speak() {
	fmt.Println(animal.noise)
	return
}

func main() {
	var request string
	var animalName string
	var animalType string
	var animalInterface animal
	for {
		fmt.Print("Prompt > ")
		fmt.Scan(&request, &animalName, &animalType)
		if request == "newanimal" {
			animalMap[animalName] = animalMap[animalType]
			fmt.Println("Created it!")
		} else if request == "query" {
			animalInterface = animalMap[animalName]
			switch animalType {
			case "eat":
				animalInterface.Eat()
			case "move":
				animalInterface.Move()
			case "speak":
				animalInterface.Speak()
			}
		} else {
			fmt.Println("wrong query")
		}

	}
}