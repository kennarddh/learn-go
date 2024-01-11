package main

import (
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
    log.SetFlags(0)

	messages, err := Hellos([]string{"test","a"})

	if err != nil {
        log.Fatal(err)
	}

	for _, message := range messages {
		fmt.Println(message)
    }
}