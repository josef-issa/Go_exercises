////Variables and data types (int, string, bool).
//Conditional statements (if, else, if, else).

//Job: Create a small program that collects everything (name + age + condition)
//Function: Create a program that asks the user for his name and age, and prints them back in a sentence.
//Function: A program that checks if the user is over 18 and prints an appropriate message.

package main

import "fmt"

func main() {
	fmt.Print("Welcome to the youtube channel name generatour!\n")
	var name string
	var age int
	var branch string
	var live string

	fmt.Print("Please, enter your name?\n")
	fmt.Scanln(&name)
	fmt.Println("Hello", name)

	fmt.Print("What is your channel about?\n")
	fmt.Scanln(&branch)
	fmt.Println("You could name your channel " + "(" + branch + " with " + name + ")\n")

	fmt.Print("Where do you live?\n")
	fmt.Scanln(&live)

	fmt.Print("Hou old are you?\n")
	fmt.Scanln(&age)
	if age >= 18 {
		fmt.Printf("Hello %s your old is %d and your channel name is %s with %s\n", name, age, branch, name)
		fmt.Println("Welcom to your Youtube channel.\n")
		fmt.Print("Congratulations.")

	} else {
		fmt.Printf("Hello %s your old is %d.\n", name, age)
		fmt.Println("Soory, You can't cerate account.\n")

	}

}
