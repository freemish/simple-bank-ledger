package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/freemish/simple-bank-ledger/bankledger"
)

// HandleInput hands off handling the user's input to helpers depending on whether
// user is logged in.
func HandleInput(scanner *bufio.Scanner, cust *bankledger.Customer) {
	if cust == nil {
		handleLoginAndRegistrationOptions(scanner)
		return
	}
	handleUserOptions(scanner, cust)
}

func handleLoginAndRegistrationOptions(scanner *bufio.Scanner) {
	switch strings.ToLower(scanner.Text()) {
	case "g": // Login
		currentUser = loginHandler(scanner)
	case "r": // Register
		currentUser = registrationHandler(scanner)
	case "h": // Help
		fmt.Println(HelpText(nil))
	case "q": // Quit
		os.Exit(1)
	default: // Invalid response
		fmt.Println("Did not recognize input.")
	}
}

func promptForInput(scanner *bufio.Scanner, prompt string) string {
	fmt.Print(prompt + ": ")
	scanner.Scan()
	return scanner.Text()
}

func loginHandler(scanner *bufio.Scanner) *bankledger.Customer {
	username := promptForInput(scanner, "Enter your username")
	password := promptForInput(scanner, "Enter your password")

	cust, err := bankledger.Login(username, password)
	if err != nil {
		fmt.Println(err.Error())
		return cust
	}

	fmt.Printf("Welcome back, %s! You have a balance of $%.2f.\n", cust.Name, cust.GetBalance())
	fmt.Println(HelpText(cust))
	return cust
}

func registrationHandler(scanner *bufio.Scanner) *bankledger.Customer {
	name := promptForInput(scanner, "Enter your full name")
	username := promptForInput(scanner, "Enter your username")
	password := promptForInput(scanner, "Enter your password")
	passwordConfirm := promptForInput(scanner, "Confirm your password by typing it again")

	if password != passwordConfirm {
		fmt.Println("Your password confirmation didn't match. Registration cancelled.")
		return nil
	}

	confirm := promptForInput(scanner, fmt.Sprintf("Are you sure you want to create an account with name %s and username %s?\nEnter y to confirm, any other key to cancel: ", name, username))
	if confirm != "y" {
		fmt.Println("Registration cancelled.")
		return nil
	}

	cust, err := bankledger.CreateAccount(name, username, password)
	if err != nil {
		fmt.Println(err.Error()) // handle making error text user-friendly
		return nil
	}
	fmt.Printf("Account %s created. You are now logged in.\n", username)
	fmt.Println(HelpText(cust))
	return cust
}

func handleUserOptions(scanner *bufio.Scanner, cust *bankledger.Customer) {
	switch strings.ToLower(scanner.Text()) {
	case "h": // Help
		fmt.Println(HelpText(cust))
	case "x": // Log out
		currentUser = nil
		fmt.Println("You have been logged out. See you later!")
		fmt.Println(HelpText(nil))
	case "v": // View history
		fmt.Println("(Implement view history.)")
	case "b": // View balance
		fmt.Printf("Your balance is $%.2f.\n", cust.GetBalance())
	case "r": // Record transaction
		recordTransactionHandler(scanner, cust)
	case "q": // Quit
		os.Exit(1)
	default: // Invalid response
		fmt.Println("Input not recognized. Enter \"h\" for options.")
	}
}

// HelpText returns appropriate help text depending on whether user is logged in.
func HelpText(cust *bankledger.Customer) string {
	if cust == nil {
		return "\"g\" - log in\n\"r\" - register\n\"h\" - help\n\"q\" - quit"
	}
	return "\"r\" - start a transaction\n\"v\" - view history\n\"b\" - view balance\n\"x\" - log out\n\"h\" - help\n\"q\" - log out and quit"
}

func recordTransactionHandler(scanner *bufio.Scanner, cust *bankledger.Customer) bool {
	promptForInput(scanner, "Enter \"w\" to withdraw funds from your account, or \"d\" to deposit funds to your account.")

	return true
}
