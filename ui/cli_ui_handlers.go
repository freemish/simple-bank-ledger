package ui

import (
	"fmt"
	"os"

	"github.com/freemish/simple-bank-ledger/adapters"
)

func GetCliHandlerFromCliOption(cliOpt cliOption) func(adapters.IProcessHandler) {
	switch cliOpt.Key {
	case MessageCliHelpKey:
		return AdapterHelp
	case MessageCliQuitKey:
		return AdapterQuit
	case MessageCliLoginKey:
		return AdapterLogin
	case MessageCliRegisterKey:
		return AdapterRegister
	default:
		return AdapterHelp
	}
}

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
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("customer was nil")
		}
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
	_, err := adapter.Register(name, username, password)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Successfully registered! Please log in to proceed.")
	}
}
