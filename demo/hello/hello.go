package main

import (
	"fmt"
	// "rsc.io/quote"
	"example.com/greetings"
	"log"
)

func main(){
	// fmt.Println("Hello, world!")
	// fmt.Println(quote.Go())
	// message := greetings.Hello("Jake")
	// fmt.Println(message)
	
	// Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)
	
	// A slice of names.
    names := []string{"Gladys", "Samantha", "Darrin"}
	

    // Request a greeting message.
    // message, err := greetings.Hello("Jake")
    // If an error was returned, print it to the console and
    // exit the program.
	
	// Request greeting messages for the names.
    messages, err := greetings.Hellos(names)
	
    if err != nil {
        log.Fatal(err)
    }

    // If no error was returned, print the returned message
    // to the console.
    // fmt.Println(message)
	
	// If no error was returned, print the returned map of
    // messages to the console.
    fmt.Println(messages)
	
}