package main

import (
	"fmt"
	"os"
	"time"
)

// create struct for user profile
type userprofile struct {
	name      string
	age       int
	email     string
	createdAt time.Time
}

// create a function to welcome the user to the system.
func WelcomeMessage() {
	fmt.Println("Welcome to the user profile system!")
	fmt.Println("Here you can create your user profile.")
}

func main() {

	// call the welcome message function
	WelcomeMessage()

	// determine the current year for birth year calculation
	currentYear := time.Now().Year()

	// create a variable to store the user profile information and prompt the user to enter their details.
	var user userprofile
	fmt.Print("Enter your Name: ")
	fmt.Scanln(&user.name)

	// ask once for the age and exit immediately if the user is under 18
	fmt.Print("Enter your Age: ")
	fmt.Scanln(&user.age)
	if user.age < 18 {
		fmt.Println("Sorry, you must be at least 18 years old to create a user profile.")
		os.Exit(1) // terminate the program with a non-zero status
	}

	fmt.Print("Enter your Email: ")
	fmt.Scanln(&user.email)

	// Initialize createdAt with current time
	user.createdAt = time.Now()

	// calculate the birth year of the user by subtracting their age from the current year.
	birthYear := currentYear - user.age

	// display the user profile information and the calculated birth year to the user.
	fmt.Println("\nUser Profile Created:")
	fmt.Printf("Name: %s\n", user.name)
	fmt.Printf("Age: %d\n", user.age)
	fmt.Printf("Email: %s\n", user.email)
	fmt.Printf("Hej %s, du är %d, och du föddes år %d, you created your profile at %v\n", user.name, user.age, birthYear, user.createdAt)
}
