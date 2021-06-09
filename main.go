// Go will only run a program of package main. We can run this file by running go run main.go
// Uncomment things as you go to test them in the terminal.
// I recommend having the "Go" extension in vscode. This will show errors without you having to compile and run your program
package main

import "fmt"

func main() {
	age := 45

	fmt.Println(age > 50)
	fmt.Println(age >= 27)
	fmt.Println(age == 35)
	fmt.Println(age != 42)

	// Simple example of if else statements
	if age < 30 {
		fmt.Println("Age is less than 30")
	} else if age < 40 {
		fmt.Println("Age is less than 40")
	} else {
		fmt.Println("Age is not less than 40")
	}

	// Usng continue and break
	names := []string{"Tomo", "Patrick", "Snickers", "Alex", "Yusuke", "Aline"}

	for index, value := range names {
		// the continue will break out of the current iteration, but continue with the loop
		// so it will output fcontinuing at position 1, but not go down to the "value at ...."
		// rather it will go back to start the for loop.
		if index == 1 {
			fmt.Println("Continuing at pos", index)
			continue
		}
		// break will break out of the loop completly, so the below will only go p to
		// position 2, and break out (the rest of the elements will not fire)
		if index > 2 {
			fmt.Println("Breaking at position", index)
			break
		}

		fmt.Printf("The value at pos %v is %v\n", index, value)
	}
}
