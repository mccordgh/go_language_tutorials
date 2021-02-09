package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	/*
		Set properties of the predefined Logger, including the log entry prefix and a flag to disable
		printing the time, source file, and line number.
	*/
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message
	message, err := greetings.Hello("Butthead")

	// A slice of names
	names := []string{"Joey", "Deedee", "Some other Ramon"}
	messages, err := greetings.Hellos(names)

	// If an error was returned, print it to the console and exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message to the console
	fmt.Println()
	fmt.Println(message)

	fmt.Println("****************************")
	for _, greeting := range messages {
		fmt.Println(greeting)
	}

	fmt.Println()
}
