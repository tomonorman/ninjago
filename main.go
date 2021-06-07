// Go will only a program of package main. We can run this file by running go run main.go
// Uncomment things as you go to test them in the terminal.
package main

import "fmt"

func main() {

	// Variables

	// Strings
	// Here we are explicity naming the type
	// var nameOne string = "Tomo"
	// // Here Go will infer that the variable is of type string given the assignment
	// var nameTwo = "Snickers"
	// // Here we are declaring the variable to be of type string before hand without a value. Later on, when we assign
	// // the variable must be a string. The default value will be an empty string if we dont give it a value ourselves.
	// var nameThree string

	// fmt.Println(nameOne, nameTwo, nameThree)

	// // We can reassign a variable as long as it is of the same type.
	// nameOne = "Patrick"
	// // and we can assign a variable as long as we have declared it before
	// nameThree = "Same"

	// fmt.Println(nameOne, nameTwo, nameThree)

	// // We can use the below shorthand to declare and assign a variable without having to use the var keyword. It can only be done
	// // when initializing the variable. Go will infer the type. The shorthand ":=" cannot be used outside a function.
	// nameFour := "Alex"
	// fmt.Println(nameFour)

	// // Integers (ints)

	// var ageOne int = 25
	// var ageTwo = 30
	// ageThree := 40

	// fmt.Println(ageOne, ageTwo, ageThree)

	// // bits and memory -> we can declare an int with a certain number of bits as per the below.
	// // an 8 bit integer can have the range -128 through 127
	// var numOne int8 = 25
	// var numTwo int8 = -128

	// // a uint must be a positive number, which means we can double the positive range of numbers available to us.
	// // eg uint8 is available up to 255.
	// var numThree uint8 = 255

	// fmt.Println(numOne, numTwo, numThree)

	// // Floats must be declared as either 32bit or 64bit
	// var floatOne float32 = 25.98

	// fmt.Println(floatOne)

	// Using fmt and formatting strings

	//  Print
	age := 35
	name := "Tomo"
	// // Print does not print a new line
	// fmt.Print("Hello, ")
	// // and Println does
	// fmt.Println("World")
	// fmt.Println("This is on a new line")
	// // Printing Variables
	// fmt.Println("My age is", age, "and my name is", name)

	// // Formatted Strings using Printf.
	// // We use %v (known as a format specifier) which will output in the default format in the order
	// // of the arguments given (think sql)
	// fmt.Printf("My age is %v and my name is %v \n", age, name)
	// // %q will wrap the variable in quotes. %q will not work with integers. Notice it outputs a hash
	// // instead of the number.
	// fmt.Printf("My age is %q and my name is %q \n", age, name)
	// // %T will give us the type of variable
	// fmt.Printf("age is of type %T name is of type %T \n", age, name)
	// // %f will output a float, but we should declare the number of decimal points we want (0.2 for two, 0.1 for 1 etc.)
	// fmt.Printf("You scored %0.1f points! \n", 22.55)
	// // Google format specifiers for more details

	// // Sprintf -> save the formatted strings. It does not print, bur rather returns the formatted string so you can store it.
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("The saved string is", str)
}
