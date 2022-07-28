package ui

import (
	"fmt"
	"testing"
)

// TODO: add assertions to these tests

func TestHelpTextNotLoggedIn(t *testing.T) {
	fmt.Println(HelpText(false))
}

func TestHelpTextLoggedIn(t *testing.T) {
	fmt.Println(HelpText(true))
}

func TestGetOptionFromInput(t *testing.T) {
	populateOptionsMap()

	fmt.Println("Testing not logged in:")
	for _, inputOption := range optionsList {
		fmt.Printf("Testing with key: %s\n", inputOption.Key)
		GetOptionFromInput(inputOption.Key, false)
	}
	fmt.Println("Testing with key: whatever")
	GetOptionFromInput("whatever", false)

	fmt.Println("Testing logged in:")
	for _, inputOption := range optionsList {
		fmt.Printf("Testing with key: %s\n", inputOption.Key)
		GetOptionFromInput(inputOption.Key, true)
	}
	fmt.Println("Testing with key: whatever")
	GetOptionFromInput("whatever", true)
}
