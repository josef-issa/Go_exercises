/*Function (Practice)

Scenario: You want to print a simple "ID card" in the terminal.

Requirements: Write a Go program:

Define the following variables: firstName (string), lastName (string), birthYear (int), and hasDrivingLicense (bool).

Calculate the approximate age using the currentYear variable (choose it as 2025) and then age := currentYear - birthYear.
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("Welcome to ID card cerate!\n")

	var firstName, lastName string
	var birthYear int

	fmt.Print("Please,enter your firstname and your lastnme?\n")
	fmt.Scanln(&firstName, &lastName)

	fmt.Print("Please enter your birtheyear?\n")
	fmt.Scanln(&birthYear)

	//firstName := "Alix"
	//lastName := "Tom"
	//birthYear := 1985

	var answer string
	fmt.Print("Do you have a car license?\n (yes/no):")
	fmt.Scanln(&answer)

	//We convert the answer to lowercase letters to avoid typing problems.
	answer = strings.ToLower(answer)

	hasDrivingLicense := false
	if answer == "yes" {
		hasDrivingLicense = true
	}

	//Calculate the approximate age using the currentYear variable (choose it as 2025)
	const currentYear = 2025

	age := currentYear - birthYear

	if hasDrivingLicense {
		fmt.Printf("Hello, %s %s, You born on year %d,so your age is: %d, and you allowed to drivinin: yes 🚗\n", firstName, lastName, birthYear, age)
	} else {
		fmt.Printf("Hello, %s %s, You born on year %d,so your age is: %d, and you allowed to drivinin: no ❌\n", firstName, lastName, birthYear, age)
	}

}
