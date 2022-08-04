package main

import (
	"fmt"
	"log"

	"github.com/eyalliebermann/gotut/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Ida")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
