// Go will only run a program of package main. We can run this file by running go run main.go
// Uncomment things as you go to test them in the terminal.
// I recommend having the "Go" extension in vscode. This will show errors without you having to compile and run your program
package main

import (
	"fmt"
	"strings"
)

// We can have multiple returns by adding a set of () between the argument () nd the {}
// The below will return two strings
func getInitials(fullName string) (string, string) {

	capitalizedFullName := strings.ToUpper(fullName)
	names := strings.Split(capitalizedFullName, " ")

	var initials []string

	for _, value := range names {
		// append only the first letter of each name (:1 will count up to but not including position 1)
		initials = append(initials, value[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func main() {
	// We can accept two variables as the return by using the below (variable, variable) syntax
	firstName, lastName := getInitials("tomo norman")
	fmt.Println(firstName, lastName)

	firstName2, lastName2 := getInitials("patrick campbell")
	fmt.Println(firstName2, lastName2)

	firstName3, lastName3 := getInitials("bono")
	fmt.Println(firstName3, lastName3)
}
