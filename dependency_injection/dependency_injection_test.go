package main_test

import (
	"bytes"
	"testing"

	. "GoLang-Showcase/dependency_injection"
)

func TestGreet(t *testing.T) {
	buffer := &bytes.Buffer{}
	Greet(buffer, "Ryan")

	got := buffer.String()
	want := "Hello, Ryan"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
