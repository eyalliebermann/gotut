package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hello %v, welcome home!", name)
	return message
}

func main() {
	fmt.Printf(Hello("Shlomo"))
}
