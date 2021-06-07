// Go will only a program of package main. We can run this file by running go run variables.go
package main

import "fmt"

func main() {

	// Strings
	// Here we are explicity naming the type
	var nameOne string = "Tomo"
	// Here Go will infer that the variable is of type string given the assignment
	var nameTwo = "Snickers"
	// Here we are declaring the variable to be of type string before hand without a value. Later on, when we assign
	// the variable must be a string. The default value will be an empty string if we dont give it a value ourselves.
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	// We can reassign a variable as long as it is of the same type.
	nameOne = "Patrick"
	// and we can assign a variable as long as we have declared it before
	nameThree = "Same"

	fmt.Println(nameOne, nameTwo, nameThree)

	// We can use the below shorthand to declare and assign a variable without having to use the var keyword. It can only be done
	// when initializing the variable. Go will infer the type. The shorthand ":=" cannot be used outside a function.
	nameFour := "Alex"
	fmt.Println(nameFour)

	// Integers (ints)

	var ageOne int = 25
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits and memory -> we can declare an int with a certain number of bits as per the below.
	// an 8 bit integer can have the range -128 through 127
	var numOne int8 = 25
	var numTwo int8 = -128

	// a uint must be a positive number, which means we can double the positive range of numbers available to us.
	// eg uint8 is available up to 255.
	var numThree uint8 = 255

	fmt.Println(numOne, numTwo, numThree)

	// Floats must be declared as either 32bit or 64bit
	var floatOne float32 = 25.98

	fmt.Println(floatOne)

}
