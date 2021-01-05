package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Hasan", "Husen", "Hasen"}

	// Request a greeting message
	message, err := greetings.Hellos(names)
	// If an eror was returned, print it to the console
	// and exit the program
	if err != nil {
		log.Fatal(err)
	}

	// If no errors was returned, than simply print the message
	fmt.Println(message)
}
