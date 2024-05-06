package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Pichu"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Pichu")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Pichu") = %q, %v, quiere coincidencia parra %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, Quiere "", error`, msg, err)
	}
}
