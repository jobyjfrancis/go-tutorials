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
	log.SetPrefix("greetings:")
	log.SetFlags(0)

	/*message1, err1 := greetings.Hello("John")
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Println(message1)

	message2, err2 := greetings.Hello("")
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(message2)*/

	// A slice of names.
	names := []string{"Gladys", "John", "Mary"}

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned map of
	// messages to the console.
	for _, message := range messages {
		fmt.Println(message)
	}
}
