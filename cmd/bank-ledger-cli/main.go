package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/freemish/simple-bank-ledger/bankledger"
	//"github.com/freemish/errgo"
)

var currentUser *bankledger.Customer

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to World's Best Bank!")
	fmt.Println(HelpText(currentUser))
	for {
		scanner.Scan()
		HandleInput(scanner, currentUser)
	}
}

func test() {
	// test backend stuff:

	// log in to make sure it fails
	fmt.Println("Test 1 - Attempting login with no matching registered users. Should error.")
	_, err := bankledger.Login("molly", "1234")
	if err != nil {
		//fmt.Println(errgo.Wrap(err).StackTrace())
		fmt.Println("Test 1 - Success!")
	} else {
		fmt.Println("Test 1 - Expected error not encountered. Failed.")
		return
	}

	// create account
	fmt.Println("Test 2 - Creating an account. Should succeed.")
	_, err = bankledger.CreateAccount("Molly", "molly", "1234")
	if err == nil {
		fmt.Println("Test 2 - No error encountered. Success!")
	} else {
		//fmt.Println(errgo.Wrap(err).StackTrace())
		fmt.Println("Test 2 - Unexpected error encountered. Failed.")
		return
	}

	// try to log in with wrong password
	fmt.Println("Test 3 - Trying to log in with wrong password. Should error.")
	_, err = bankledger.Login("molly", "123")
	if err != nil {
		//fmt.Println(errgo.Wrap(err).StackTrace())
		fmt.Println("Test 3 - Success!")
	} else {
		fmt.Println("Test 3 - Expected error not encountered. Failed.")
		return
	}

	// try to log in with right password
	fmt.Println("Test 4 - Trying to log in with correct password. Should succeed.")
	cust, err := bankledger.Login("molly", "1234")
	if err != nil {
		//fmt.Println(errgo.Wrap(err).StackTrace())
		fmt.Println("Test 4 - Unexpected error encountered. Failed.")
		return
	}
	fmt.Println("Test 4 - No error encountered. Success!")

	// record a deposit
	fmt.Println("Test 5 - Record a deposit. Should succeed.")
	err = cust.RecordTransaction("Test Transaction", "", 30.0)
	if err != nil {
		//fmt.Println(errgo.Wrap(err).StackTrace())
		fmt.Println("Test 5 - Unexpected error encountered. Failed.")
		return
	}
	fmt.Println("Test 5 - No error encountered. Success!")

	// get balance
	fmt.Println("Test 6 - Get customer's balance.")
	balance := cust.GetBalance()
	if balance != 30.0 {
		fmt.Printf("Test 6 - Balance is %.2f rather than expected value, 30. Failed.\n", balance)
		return
	}
	fmt.Printf("Test 6 - Balance is %.2f. Success!\n", balance)

	// attempt to record too-large withdrawal
	fmt.Println("Test 7 - Attempt to record a too-large withdrawal. Should error.")
	err = cust.RecordTransaction("Test Transaction 2", "", -30.01)
	if err == nil {
		fmt.Println("Test 7 - No error encountered. Failed.")
		return
	}
	//fmt.Println(errgo.Wrap(err).StackTrace())
	fmt.Println("Test 7 - Expected error encountered. Success!")

	// attempt to record valid withdrawal
	fmt.Println("Test 8 - Record valid transaction. Should succeed.")
	err = cust.RecordTransaction("Test Transaction 3", "", -30.00)
	if err != nil {
		//fmt.Println(errgo.Wrap(err).StackTrace())
		fmt.Println("Test 8 - Unexpected error encountered. Failed.")
		return
	}
	fmt.Println("Test 8 - Success!")

	// get history
	fmt.Println("Getting customer's history.")
	history := cust.GetAllHistory()
	for _, t := range history {
		fmt.Println(t)
	}
	fmt.Printf("Final balance: %.2f\n", cust.GetBalance())

	// log out
	cust = nil
	balance = cust.GetBalance()
	if balance != 0 {
		fmt.Println("Customer failed to log out?")
	}
}
