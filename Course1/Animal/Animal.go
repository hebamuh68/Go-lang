package main

import (
	"fmt"
)

//====================================================
type Animal struct{
	food, locomotion, noise string
}

//====================================================
var Cow = Animal{"grass", "walk", "moo"}
var Bird = Animal{"worms", "fly", "peep"}
var Snake = Animal{"mice", "slither", "hsss"}

//====================================================
func Eat(animal Animal){
	fmt.Println(animal.food)
}

func Move(animal Animal){
	fmt.Println(animal.locomotion)
}

func Speak(animal Animal){
		fmt.Println(animal.noise)
}