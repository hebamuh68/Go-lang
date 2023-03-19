/**
* Every request from the user must be a single line containing 2 strings. 
* The first string is the name of an animal, either “cow”, “bird”, or “snake”. 
* The second string is the name of the information requested about the animal, either “eat”, “move”, or “speak”. 
*/

package main

import (
	"fmt"
)

func main(){
	
	var animal, info string

	for{

		fmt.Printf("> ")
		fmt.Scanln(&animal, &info)

		if info == "eat"{
			switch animal {
			case "cow": Eat(Cow)
			case "bird": Eat(Bird)
			case "snack": Eat(Snake)
			}

		} else if info == "move" {
			switch animal {
			case "cow": Move(Cow)
			case "bird": Move(Bird)
			case "snack": Move(Snake)
			}

		} else {
			switch animal {
			case "cow": Speak(Cow)
			case "bird": Speak(Bird)
			case "snack": Speak(Snake)
			}

		}	
	}
}