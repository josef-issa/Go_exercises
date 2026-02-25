/*
** Real-World Scenario

You are building a simple banking system.
Each bank account has:
Owner name
Balance
Transaction history (list of transactions with type, amount, and date)
check balance */

package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
	"time"

)

// cerate Welcome message for the user and display the available options for the calculator.
func WelcomeMessage() {
	fmt.Println(`
	 Welcome to the bank account system!
	 Please choose one of the following options:
	 1. Deposit money
	 2. Withdraw money
	 3. Check balance
	 4. View transaction history
	 5. Exit
	`)
	
}

// create a struct to represent a transaction with fields for the transaction type, amount, balance, and date.

type transaction struct {
	Type   string
	Amount float64
	Balance float64
	Date   time.Time
}

// create a struct to represent a bank account with fields for the account name and balance, and methods for depositing and withdrawing money.

type BankAccount struct {
	AccountName string
	Balance       float64
	transactions    []transaction
}

// create a function to welcome the user to the system and display the available options.

func (account *BankAccount) Deposit(amount float64) {
	
    if amount <= 0{
		fmt.Println("Invalid amount! Deposit failed.")
		return
	}
	
	account.Balance += amount
	account.recordTransaction("Deposit", amount)
	fmt.Printf("✅ Deposit successful! New balance: %.2f €\n", account.Balance)
	
}

// create a function to withdraw money from the account, ensuring that the withdrawal amount does not exceed the current balance.

func (account *BankAccount) Withdraw(amount float64) {
	if amount <= 0 {
		fmt.Println("Invalid amount! Withdrawal failed.")
		return
	}
	
	if amount > account.Balance {
		fmt.Println("Insufficient balance! Withdrawal failed.")
		return
	
	} else {
		account.Balance -= amount
		account.recordTransaction("Withdrawal", amount)
		fmt.Printf("✅ Withdrawal successful! New balance: %.2f €\n", account.Balance)
	}
}
// recordTransaction adds a new transaction to the history

func (account *BankAccount) recordTransaction(txType string, amount float64) {
	tx := transaction{
		Type:      txType,
		Amount:    amount,
		Balance:   account.Balance,
		Date:      time.Now(),
	}
	account.transactions = append(account.transactions, tx)
}

// PrintHistory displays all transactions for the account

func (account *BankAccount) PrintHistory() {
	if len(account.transactions) == 0 {
		fmt.Println("No transactions found.")
		return
	}

	fmt.Printf("\n===== Transaction History for %s =====\n", account.AccountName)
	fmt.Printf("%-5s %-12s %-10s %-12s %-20s\n", "#", "Type", "Amount", "Balance", "Date & Time")
	fmt.Println(strings.Repeat("-", 65))

	for i, tx := range account.transactions {
		fmt.Printf("%-5d %-12s %-10.2f %-12.2f %s\n",
			i+1,
			tx.Type,
			tx.Amount,
			tx.Balance,
			tx.Date.Format("2006-01-02 15:04:05"),
		)
	}
	fmt.Println(strings.Repeat("-", 65))
	fmt.Printf("Total transactions: %d\n\n", len(account.transactions))

}

func main() {
	// call the welcome message function to greet the user and display the available options.
	WelcomeMessage()

	reader := bufio.NewReader(os.Stdin)

	// create a variable to store the user input and prompt the user to enter their choice of calculation.
	fmt.Print("Please enter your account name: ")
	accountName, _ := reader.ReadString('\n')
	accountName = strings.TrimSpace(accountName)

	account := BankAccount{AccountName: accountName, Balance: 0.0}
	fmt.Println()
	fmt.Printf("Welcome, %s! Your account has been created with a balance of %.2f €\n", account.AccountName, account.Balance)
	fmt.Println()
	for {
		fmt.Print("Please choose an option (1-5): ")
		fmt.Print()
		optionStr, _ := reader.ReadString('\n')
		optionStr = strings.TrimSpace(optionStr)
		option, err := strconv.Atoi(optionStr)
		if err != nil {
			fmt.Println("Invalid option! Please enter a number between 1 and 4.")
			continue
		}
		if option == 5 {
		fmt.Println("Thank you for using the bank account system! Goodbye!")
		os.Exit(0) // exit the program
		}

	
		switch option {
		case 1:
			fmt.Print("Enter the amount to deposit: ")
			fmt.Print()
			amountStr, _ := reader.ReadString('\n')
			amountStr = strings.TrimSpace(amountStr)
			amount, err := strconv.ParseFloat(amountStr, 64)
			if err != nil {
				fmt.Println("Invalid amount! Please enter a valid number.")
				continue
			}
			account.Deposit(amount)
			fmt.Println()
			//fmt.Printf("Your current balance is: %.2f € \n", account.Balance)

		case 2:
			fmt.Print("Enter the amount to withdraw: ")
			fmt.Print()
			amountStr, _ := reader.ReadString('\n')
			amountStr = strings.TrimSpace(amountStr)
			amount, err := strconv.ParseFloat(amountStr, 64)
			if err != nil {
				fmt.Println("Invalid amount! Please enter a valid number.")
				continue
			}
			account.Withdraw(amount)
			fmt.Println()
			//fmt.Printf("Your current balance is: %.2f € \n", account.Balance)

		case 3:
			fmt.Printf("Your current balance is: %.2f € \n", account.Balance)
			fmt.Println()
		case 4:
			account.PrintHistory()
            fmt.Println()
		case 5:
			fmt.Printf("Thank you for using the bank account system! %s! Goodbye! 👋\n", account.AccountName)
			return

		default:
			fmt.Println("Invalid option! Please enter a number between 1 and 5.")
		}
	}
}

