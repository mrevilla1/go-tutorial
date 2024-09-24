package main

import (
	"fmt"
	"log"

	"rabor.io/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	message, err := greetings.Hello("dolor")
	message2, err2 := greetings.Hellos([]string{"profundo", "testicular"})
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil || err2 != nil {
		log.Fatal(err)
		log.Fatal(err2)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
	fmt.Println(message2)
    go list -f '{{.Target}}'
    go list -f '{{.Target}}'
    /home/marco/go/bin/hello
    set PATH=%PATH%;C:\path\to\your\install\directory
    export PATH=$PATH:home/marco/go/bin/hello
go 