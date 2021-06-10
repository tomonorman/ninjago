// Go will only run a program of package main. We can run this file by running go run main.go
// Uncomment things as you go to test them in the terminal.
// I recommend having the "Go" extension in vscode. This will show errors without you having to compile and run your program
package main

import (
	"fmt"
	"math"
)

// When we declare the arguments a function can take, we must also declare what type
// that function will be. It cannot accept any other type.
func sayGreeting(name string) {
	fmt.Printf("Good Morning %v \n", name)
}

func sayBye(name string) {
	fmt.Printf("Good Bye %v \n", name)
}

// A function can also take a function as an argument. Here we will give the function
// a slice of strings, and a function that will do some kind of greeting.
func cycleNames(names []string, greeting func(string)) {
	// for each name (we dont need the index) in names, we will fire the
	// greeting function we have passed in as an argument.
	for _, name := range names {
		greeting(name)
	}
}

// When we use a function to return something, we must state what type we will be returning
// between the () and {}, in our case a float64 to get the area of a circle (pi r squared)
func circleArea(radius float64) float64 {
	return math.Pi * (radius * radius)
}

// The main function fires automatically when we run main.go
func main() {
	sayGreeting("Tomo")
	sayGreeting("Snickers")
	sayBye("Tomo")

	// We call cycleNames and give it a list of names as a first argument, and then the sayGreeting function as
	// the second argument. Notice we dont use the () at the end of the function as we dont want to invoke it,
	// only give a reference to it.
	cycleNames([]string{"Patrick", "Alex", "Khor"}, sayGreeting)
	// Same again, but with the sayBye function
	cycleNames([]string{"Patrick", "Alex", "Khor"}, sayBye)

	areaOne := circleArea(10)
	areaTwo := circleArea(7.4)

	fmt.Println(areaOne, areaTwo)
	// We can do the same as above but print out to 3 decimal places
	fmt.Printf("Circle one is %0.3f and circle two %0.3f", areaOne, areaTwo)
}
