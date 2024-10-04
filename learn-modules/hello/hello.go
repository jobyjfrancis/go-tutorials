package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings:")
	log.SetFlags(0)

	message1, err1 := greetings.Hello("John")
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Println(message1)

	message2, err2 := greetings.Hello("")
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(message2)
}
