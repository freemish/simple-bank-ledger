package ui

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/freemish/simple-bank-ledger/adapters"
	"github.com/freemish/simple-bank-ledger/processes"
)

func AdapterQuit(adapters.IProcessHandler) {
	os.Exit(1)
}

func AdapterHelp(adapter adapters.IProcessHandler) {
	fmt.Println(HelpText(adapter.GetLoggedInCustomer() != nil))
}

func AdapterLogin(adapter adapters.IProcessHandler) {
	username := PromptForInput("Enter your username: ")
	password := PromptForInput("Enter your password: ")
	cust, err := adapter.Login(username, password)
	if err != nil || cust == nil {
		fmt.Println("There was an error logging you in. Please try again.")
		fmt.Println(HelpText(false))
	} else {
		fmt.Printf("Successfully logged in. Welcome, %s.\n", cust.Name)
		fmt.Println(HelpText(true))
	}
}

func AdapterRegister(adapter adapters.IProcessHandler) {
	name := PromptForInput("Enter your full name: ")
	username := PromptForInput("Enter your username: ")
	password := PromptForInput("Enter your password: ")
	passwordConfirm := PromptForInput("Confirm your password by typing it again: ")

	if password != passwordConfirm {
		fmt.Println("Your password confirmation didn't match. Registration cancelled.")
		return
	}

	pwValErr := runCliPasswordValidationRules(password)
	if pwValErr != nil {
		fmt.Println(pwValErr.Error())
	}

	_, err := adapter.Register(name, username, password)
	if err != nil {
		if err == processes.ErrCustomerAlreadyExists {
			fmt.Println("Please try again with another username.")
			return
		}

		fmt.Println("Error:", err.Error())
		return
	}
	fmt.Println("Successfully registered! Please log in to proceed.")
}

func runCliPasswordValidationRules(password string) error {
	return nil
}

func AdapterLogout(adapter adapters.IProcessHandler) {
	adapter.Logout("")
	fmt.Println("Successfully logged out.")
}

func AdapterBalance(adapter adapters.IProcessHandler) {
	bal, _ := adapter.GetBalance(adapter.GetLoggedInCustomer())
	fmt.Printf("Balance: $%.2f\n", bal)
}

func AdapterTransact(adapter adapters.IProcessHandler) {
	amount, parseErr := strconv.ParseFloat(
		PromptForInput("Enter the transaction amount: "), 64)
	for parseErr != nil {
		amount, parseErr = strconv.ParseFloat(
			PromptForInput("Invalid float input. Please re-enter the transaction amount: "), 64)
	}

	debitCreditStr := strings.ToLower(PromptForInput("Enter d for debit, c for credit: "))
	for debitCreditStr != "d" && debitCreditStr != "c" {
		debitCreditStr = strings.ToLower(
			PromptForInput("Sorry, please enter either d or c. Enter d for debit, c for credit: "))
	}

	voidedIDStr := PromptForInput("Enter an ID if voiding a transaction, else leave blank: ")
	if voidedIDStr == "" {
		voidedIDStr = "0"
	}
	voidedID, parseErr := strconv.Atoi(voidedIDStr)
	for parseErr != nil {
		voidedIDStr := PromptForInput("Could not parse input as a number. Please re-enter an ID if voiding a transaction, else leave blank: ")
		if voidedIDStr == "" {
			voidedIDStr = "0"
		}
		voidedID, parseErr = strconv.Atoi(voidedIDStr)
	}

	adapter.RecordTransaction(
		adapter.GetLoggedInCustomer(),
		debitCreditStr == "d",
		amount,
		voidedID,
	)
}

func AdapterViewTransactions(adapter adapters.IProcessHandler) {
	txs, _ := adapter.GetTransactions(adapter.GetLoggedInCustomer())
	for _, t := range txs {
		fmt.Println("\t", t.DebugString())
	}
}
