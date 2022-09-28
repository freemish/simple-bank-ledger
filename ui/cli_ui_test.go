package ui

import (
	"fmt"
	"testing"
)

func TestHelpTextNotLoggedIn(t *testing.T) {
	InitializeUIModule()
	want := "\"g\" - log in\n\"r\" - register\n\"h\" - help\n\"q\" - quit"
	got := HelpText(false)

	if want != got {
		t.Errorf("Wanted: \n---%s but got: \n---%s", want, got)
	}
}

func TestHelpTextLoggedIn(t *testing.T) {
	InitializeUIModule()
	got := HelpText(true)
	want := "\"h\" - help\n\"t\" - start transaction\n\"v\" - view history\n\"b\" - balance\n\"x\" - log out\n\"q\" - quit"
	if want != got {
		t.Errorf("Wanted: \n---%s but got: \n---%s", want, got)
	}
}

func TestGetOptionFromInput(t *testing.T) {
	InitializeUIModule()

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
