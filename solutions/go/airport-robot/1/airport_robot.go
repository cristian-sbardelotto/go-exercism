package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
type Greeter interface {
	LanguageName() string
	Greet(visitorName string) string
}

type German struct{}

func (greeter German) LanguageName() string {
	return "German"
}

func (greeter German) Greet(visitorName string) string {
	return fmt.Sprintf("Hallo, %s!", visitorName)
}

type Italian struct{}

func (greeter Italian) LanguageName() string {
	return "Italian"
}

func (greeter Italian) Greet(visitorName string) string {
	return fmt.Sprintf("Ciao %s!", visitorName)
}

type Portuguese struct{}

func (greeter Portuguese) LanguageName() string {
	return "Portuguese"
}

func (greeter Portuguese) Greet(visitorName string) string {
	return fmt.Sprintf("Olá %s!", visitorName)
}

func SayHello(visitorName string, greeter Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", greeter.LanguageName(), greeter.Greet(visitorName))
}
