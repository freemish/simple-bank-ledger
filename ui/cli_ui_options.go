package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

type optionToFunction struct {
	Key                 string
	Desc                string
	Function            func(args ...interface{})
	FunctionArgs        []interface{}
	ShowWhenLoggedIn    bool
	ShowWhenNotLoggedIn bool
}

var optionsList = []optionToFunction{
	{
		Key:                 MessageCliLoginKey,
		Desc:                MessageCliLoginDesc,
		ShowWhenLoggedIn:    false,
		ShowWhenNotLoggedIn: true,
	},
	{
		Key:                 MessageCliRegisterKey,
		Desc:                MessageCliRegisterDesc,
		ShowWhenLoggedIn:    false,
		ShowWhenNotLoggedIn: true,
	},
	{
		Key:                 MessageCliHelpKey,
		Desc:                MessageCliHelpDesc,
		ShowWhenLoggedIn:    true,
		ShowWhenNotLoggedIn: true,
	},
	{
		Key:                 MessageCliTransactKey,
		Desc:                MessageCliTransactDesc,
		ShowWhenLoggedIn:    true,
		ShowWhenNotLoggedIn: false,
	},
	{
		Key:                 MessageCliViewHistoryKey,
		Desc:                MessageCliViewHistoryDesc,
		ShowWhenLoggedIn:    true,
		ShowWhenNotLoggedIn: false,
	},
	{
		Key:                 MessageCliBalanceKey,
		Desc:                MessageCliBalanceDesc,
		ShowWhenLoggedIn:    true,
		ShowWhenNotLoggedIn: false,
	},
	{
		Key:                 MessageCliLogOutKey,
		Desc:                MessageCliLogOutDesc,
		ShowWhenLoggedIn:    true,
		ShowWhenNotLoggedIn: false,
	},
	{
		Key:                 MessageCliQuitKey,
		Desc:                MessageCliQuitDesc,
		ShowWhenLoggedIn:    true,
		ShowWhenNotLoggedIn: true,
	},
}

var optionsMap = make(map[string]optionToFunction)

func populateOptionsMap() {
	if len(optionsMap) > 0 {
		return
	}
	for _, val := range optionsList {
		optionsMap[val.Key] = val
	}
}

func StartInteraction() {
	populateOptionsMap()
	fmt.Println(MessageCliWelcomeMessage)
	fmt.Println(HelpText(false))
	for {
		PromptForInput("")
	}
}

func PromptForInput(prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return strings.ToLower(scanner.Text())
}

func HelpText(loggedIn bool) string {
	var promptsToJoin = []string{}
	for _, val := range optionsList {
		if (loggedIn && val.ShowWhenLoggedIn) || (!loggedIn && val.ShowWhenNotLoggedIn) {
			promptsToJoin = append(promptsToJoin, fmt.Sprintf("\"%s\" - %s", val.Key, val.Desc))
		}
	}
	return strings.Join(promptsToJoin, "\n")
}

func GetOptionFromInput(input string, loggedIn bool) *optionToFunction {
	optionVal, exists := optionsMap[input]
	fmt.Println(exists)
	if !exists || (loggedIn && !optionVal.ShowWhenLoggedIn) || (!loggedIn && !optionVal.ShowWhenNotLoggedIn) {
		fmt.Println()
		fmt.Println(HelpText(loggedIn))
		return nil
	}
	return &optionVal
}