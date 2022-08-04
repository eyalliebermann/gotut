package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty Name")
	}
	message := fmt.Sprintf(getRandomFormat(), name)
	return message, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Hellos(names []string) ([]string, error) {
	empty := []string{"", ""}
	return empty, errors.New("Not Implemented")
}
func getRandomFormat() string {
	formats := []string{
		"Hello %v, welcome back home!",
		"Buenas tardes, %v. Que deurmas con angelitos :-)",
		"Guten Abend, %v!",
		"Good evening %v",
	}
	return formats[rand.Intn(len(formats))]
}
