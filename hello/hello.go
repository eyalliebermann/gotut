package main

import (
	"fmt"
	"log"

	"github.com/eyalliebermann/gotut/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Julius")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	names := []string{"Kaufman", "Ludwig", "Martha"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
