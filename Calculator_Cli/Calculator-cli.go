/* A program running from the terminal dose the following:
1_The user is asked to enter tow numbers.
2- He is asked to choose the process (addition, subtraction, mutiplication, division)
3- Performs the operation and displays the resultat.
*/

package main

import (
	"bufio"
	"fmt"
	"strconv"
	"os"
	"strings"
	
)

// create a function to welcome the user to the system.
func Mymessage(){
	fmt.Println(`
	🧮 Welcome to the fan calculator?

		please choese on of the process:
		*- choose (+) for Addition
		*- choose (-) forsubtraction
		*- choose (*) forMutiplication
		*- choose (/) forDivision 
		*- choose (q) to exit the program
	
	`)
}

func main(){
	// call the welcome message function
	Mymessage()

	reader := bufio.NewReader(os.Stdin)

	// create a variable to store the user input and prompt the user to enter their choice of calculation.
	
	fmt.Print("Please choose your operator, Enter your chooce: ")
    operator, _ := reader.ReadString('\n')
	operator = strings.TrimSpace(operator)
	if operator == "q" || operator == "Q" {
		fmt.Println("Thank you for using the calculator. Goodbye!")
		os.Exit(0) // exit the program
	}


	// create a variable to store the user input and prompt the user to enter the numbers for calculation.
	var num1, num2 float64
	fmt.Print("Enter the first number: ")
	num1Str, _ := reader.ReadString('\n')
	num1Str = strings.TrimSpace(num1Str)
	num1Str = strings.ReplaceAll(num1Str, ",", ".") // Replace comma with dot for decimal numbers
	num1, err1 := strconv.ParseFloat(num1Str, 64)
	if err1 != nil {
		fmt.Println("⚠️ Error: Invalid input for the first number. Please enter a valid number.")
		return
	}

	
	fmt.Print("Enter the second number: ")
	num2Str, _ := reader.ReadString('\n')
	num2Str = strings.TrimSpace(num2Str)
	num2Str = strings.ReplaceAll(num2Str, ",", ".") // Replace comma with dot for decimal numbers
	num2, err2 := strconv.ParseFloat(num2Str, 64)
	if err2 != nil {
		fmt.Println("⚠️ Error: Invalid input for the second number. Please enter a valid number.")
		return
	}

	// perform the calculation based on the user's choice and display the result to the user.
	var result float64
	valid := true
	switch operator {
	case "+":
		result = num1 + num2
		
	case "-":
		result = num1 - num2
		
	case "*":
		result = num1 * num2
		
	case "/":
		if num2 != 0 {
			result = num1 / num2
			
		} else {
			fmt.Println("⚠️ Error: Division by zero is not allowed.")
		}
	
	default:
		fmt.Println("❌ Invalid choice. Please select a valid option; please choise just one of the following: (+), (-), (*), (/)")
		
	}

	if valid {
		fmt.Printf("✅ Result is: %.2f %s %.2f = %.2f\n", num1, operator, num2, result)
	}

	
	// ask the user if they want to perform another calculation and repeat the process if they choose to do so.
	
	fmt.Print("Do you want to perform another calculation? (yes/no): ")
	again, _ := reader.ReadString('\n')
	again = strings.TrimSpace(again)
	
	if strings.ToLower(again) == "yes" {
		main() // call main function again to repeat the process
	} else {
		fmt.Println("Thank you for using the calculator. Goodbye!")
		os.Exit(0) // exit the program
	}

	

}
