// Created by: Charlotte Jhu
// Created on: May 2023
//
// This program tells a user what rating movies they can watch

package main

import (
	"fmt"
)

func main() {
	// This function tells the user what rating movies they can watch
	// variables
	var userAge int

	// input
	fmt.Println("This program tells you what rating movies you can watch.")
	fmt.Println()
	fmt.Print("Enter your age: ")
	fmt.Scanln(&userAge)

	// process
	if userAge >= 17 {
		// output
		fmt.Println("You can watch R rated movies alone.")
	} else if userAge >= 13 {
		// output
		fmt.Println("You can watch PG-13 rated movies alone.")
	} else if userAge >= 5 {
		// output
		fmt.Println("You can watch G or PG rated movies alone.")
	} else {
		// output
		fmt.Println("You are too young to watch any movies.")
	}
	fmt.Println("\nDone.")
}
