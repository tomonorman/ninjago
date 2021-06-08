// Go will only run a program of package main. We can run this file by running go run main.go
// Uncomment things as you go to test them in the terminal.
// I recommend having the "Go" extension in vscode. This will show errors without you having to compile and run your program
// It will also automatially import packages as you use them in your code
package main

import (
	"fmt"
	"sort"
)

func main() {
	// greeting := "Hi Everyone!"
	// // The Contains method takes two arguments, the string we are searching, and then what we are searching for. The below will
	// // return true.
	// fmt.Println(strings.Contains(greeting, "Hi"))
	// // The ReplaceAll method takes three arguments, the string, what to replace, and what to replace it with. It does not alter
	// // the original string
	// fmt.Println(strings.ReplaceAll(greeting, "Hi", "Hello"))
	// // Make everything uppercase, again does not alter the original string
	// fmt.Println(strings.ToUpper(greeting))
	// // Find the index of the search term (second argument)
	// fmt.Println(strings.Index(greeting, "ry"))
	// // Returns a splice that has been split based on the second argument.
	// fmt.Println(strings.Split(greeting, " "))

	// Sorts Package
	ages := []int{45, 33, 27, 75, 60, 16, 80}
	// The ints method will sort integers in an array. It is DESTRUCTIVE and will alater the array
	sort.Ints(ages)
	fmt.Println(ages)

	// Return the position of the second argument given an array (first argument)
	index := sort.SearchInts(ages, 33)
	fmt.Println(index)

	names := []string{"Tomo", "Alex", "Snickers", "Sam", "Patrick", "Liz", "Aline"}
	// Sort an array of strings in alphabetical order.
	sort.Strings(names)
	fmt.Println(names)
	// Search for the position of a string
	fmt.Println(sort.SearchStrings(names, "Tomo"))
}
