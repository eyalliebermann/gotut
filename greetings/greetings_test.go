package greetings

import (
	"errors"
	"regexp"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	name := "Paula"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q,%v, want match for %q,nil`, msg, err, want)
	}
}

func TestShouldFailOnErrorNotNull(t *testing.T) {
	name := "Paula"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)
	err = errors.New("Fake Error")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q,%v, want match for %q,nil`, msg, err, want)
	}
}
