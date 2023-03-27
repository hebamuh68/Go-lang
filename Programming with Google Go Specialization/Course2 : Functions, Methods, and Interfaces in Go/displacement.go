/*
*s = ½ a t2 + vt + s
*
*Write a program which first prompts the user
*to enter values for acceleration, initial velocity, and initial displacement.
*Then the program should prompt the user to enter a value for time and the
*program should compute the displacement after the entered time.

========================================================================

*use a function called GenDisplaceFn() which takes three float64 arguments
*return a function which computes displacement as a function of time,
assuming the given values acceleration, initial velocity, and initial
displacement. The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one
float64 argument which is the displacement travelled after time t.
*/

package main

import (
	"fmt"
)


func main() {
	var a, v, s, t float64

	fmt.Print("Enter the acceleration: ")
	fmt.Scanf("%f",&a)
	fmt.Print("Enter the velocity: ")
	fmt.Scanf("%f",&v)
	fmt.Print("Enter the displacement: ")
	fmt.Scanf("%f",&s)


    displacement := GenDisplaceFn(a, v, s)

	fmt.Print("Please enter the time: ")
	fmt.Scanf("%g",&t)
    displacement_in_time := displacement(t)

    fmt.Printf("%.2f\n",displacement_in_time)
}


//========================================================================
func GenDisplaceFn(a, v, s float64) func(float64) float64 {
    return func( t float64) float64 {
        return ((1/2)*a*(t*t)) + (v*t) + s
    }
}