package main

import (
	"fmt"
	"log"

	"github.com/eyalliebermann/gotut/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Emil")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hello World!\n", message)

	_, oops := greetings.Hello("")
	if oops != nil {
		log.Fatal(oops)
	} else {
		fmt.Printf("Code should never get to this point")
	}
}
