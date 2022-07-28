package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/freemish/simple-bank-ledger/adapters"
)

var scanner = bufio.NewScanner(os.Stdin)
var loggedOutHelpText = ""
var loggedInHelpText = ""
var optionsMap = make(map[string]cliOption)

type cliOption struct {
	Key                 string
	Desc                string
	ShowWhenLoggedIn    bool
	ShowWhenNotLoggedIn bool
	CliFunc             func(adapters.IProcessHandler)
}

var optionsList = []cliOption{
	{
		Key:                 MessageCliLoginKey,
		Desc:                MessageCliLoginDesc,
		ShowWhenLoggedIn:    false,
		ShowWhenNotLoggedIn: true,
		CliFunc:             AdapterLogin,
	},
	{
		Key:                 MessageCliRegisterKey,
		Desc:                MessageCliRegisterDesc,
		ShowWhenLoggedIn:    false,
		ShowWhenNotLoggedIn: true,
		CliFunc:             AdapterRegister,
	},
	{
		Key:                 MessageCliHelpKey,
		Desc:                MessageCliHelpDesc,
		ShowWhenLoggedIn:    true,
		ShowWhenNotLoggedIn: true,
		CliFunc:             AdapterHelp,
	},
	{
		Key:                 MessageCliTransactKey,
		Desc:                MessageCliTransactDesc,
		ShowWhenLoggedIn:    true,
		ShowWhenNotLoggedIn: false,
		CliFunc:             AdapterTransact,
	},
	{
		Key:                 MessageCliViewHistoryKey,
		Desc:                MessageCliViewHistoryDesc,
		ShowWhenLoggedIn:    true,
		ShowWhenNotLoggedIn: false,
		CliFunc:             AdapterViewTransactions,
	},
	{
		Key:                 MessageCliBalanceKey,
		Desc:                MessageCliBalanceDesc,
		ShowWhenLoggedIn:    true,
		ShowWhenNotLoggedIn: false,
		CliFunc:             AdapterBalance,
	},
	{
		Key:                 MessageCliLogOutKey,
		Desc:                MessageCliLogOutDesc,
		ShowWhenLoggedIn:    true,
		ShowWhenNotLoggedIn: false,
		CliFunc:             AdapterLogout,
	},
	{
		Key:                 MessageCliQuitKey,
		Desc:                MessageCliQuitDesc,
		ShowWhenLoggedIn:    true,
		ShowWhenNotLoggedIn: true,
		CliFunc:             AdapterQuit,
	},
}

func InitializeUIModule() {
	populateOptionsMap()
	loggedOutHelpText = calculateHelpText(false)
	loggedInHelpText = calculateHelpText(true)
}

func populateOptionsMap() {
	if len(optionsMap) > 0 {
		return
	}
	for _, val := range optionsList {
		optionsMap[val.Key] = val
	}
}

func calculateHelpText(loggedIn bool) string {
	var promptsToJoin = []string{}
	for _, val := range optionsList {
		if (loggedIn && val.ShowWhenLoggedIn) || (!loggedIn && val.ShowWhenNotLoggedIn) {
			promptsToJoin = append(promptsToJoin, fmt.Sprintf("\"%s\" - %s", val.Key, val.Desc))
		}
	}
	return strings.Join(promptsToJoin, "\n")
}

func StartInteraction() {
	imh := adapters.NewInMemoryHandler()
	InitializeUIModule()

	fmt.Println(MessageCliWelcomeMessage)
	fmt.Println(HelpText(imh.GetLoggedInCustomer() != nil))
	for {
		inputStr := PromptForInput(">>> ")
		funcOpt := GetOptionFromInput(inputStr, imh.GetLoggedInCustomer() != nil)
		if funcOpt != nil {
			funcOpt.CliFunc(imh)
		} else {
			fmt.Println(MessageCliDidNotRecognizeInput)
			fmt.Println(HelpText(imh.GetLoggedInCustomer() != nil))
		}
	}
}

func PromptForInput(prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return strings.ToLower(scanner.Text())
}

func HelpText(loggedIn bool) string {
	if loggedIn {
		return loggedInHelpText
	}
	return loggedOutHelpText
}

func GetOptionFromInput(input string, loggedIn bool) *cliOption {
	optionVal, exists := optionsMap[input]
	if !exists || (loggedIn && !optionVal.ShowWhenLoggedIn) || (!loggedIn && !optionVal.ShowWhenNotLoggedIn) {
		return nil
	}
	return &optionVal
}
