package main

import (
	"fmt"
	"log"

	"goStudy.com/hello/greeting"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	fmt.Println("Hello, World!")

	message, err := greeting.Hello("Chen")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	names := []string{"Chen", "Liu"}
	messages, err := greeting.Hellos(names)

	if nil != err {
		log.Fatal(err)
	}
	fmt.Println(messages)

}
