// Go will only run a program of package main. We can run this file by running go run main.go
// Uncomment things as you go to test them in the terminal.
// I recommend having the "Go" extension in vscode. This will show errors without you having to compile and run your program
package main

import "fmt"

func main() {

	x := 0
	// for is used for all kinds of loops
	// Below is while x is less than 5 do stuff
	for x < 5 {
		fmt.Println("value of x is:", x)
		x++
	}
	// Does the same as above, but starts i as 0, runs until i is less than 5, and each
	// time increaments i
	for i := 0; i < 5; i++ {
		fmt.Println("Value of i is:", i)
	}

	// Looping a slice of strings by its length
	names := []string{"Tomo", "Liz", "Snickers", "Patrick"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// Cycle through a particular slice, and get the index and the value
	// Similar to a for in loop
	for index, value := range names {
		fmt.Printf("The value at index: %v, is %v.\n", index, value)
	}

	// If we dont want to use the index, or the value, we need to replace it with an underscore
	for _, value := range names {
		fmt.Printf("The value at is %v.\n", value)
		value = "new string"
	}

	// Note how despite changing the value of each position to "new string" above, it does not change
	// the original array (as can be seen by running the below). value only has local scope.
	fmt.Println(names)
}
