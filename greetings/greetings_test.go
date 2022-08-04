package greetings

import (
	"regexp"
	"testing"
)

func TestEmptyString(t *testing.T) {
	name := ""
	msg, err := Hello(name)
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q,%v, want match for "",err`, msg, err)
	}
}

func TestHelloWName(t *testing.T) {
	name := "Paula"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Paula") = %q,%v, want match for %q,nil`, msg, err, want)
	}
}

func TestHelloWWrongName(t *testing.T) {
	name1 := "Quelle"
	name2 := "Richard"
	want1 := regexp.MustCompile(`\b` + name1 + `\b`)
	want2 := regexp.MustCompile(`\b` + name2 + `\b`)

	msg1, err := Hello(name1)
	//verifying test code by trying a different regex
	if !want1.MatchString(msg1) || want2.MatchString(msg1) || err != nil || msg1 == "" {
		t.Fatalf(`Hello("Quelle") = %q,%v, want match for %q,nil`, msg1, err, want1)
	}
}
