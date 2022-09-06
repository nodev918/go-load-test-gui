package main

import (
	"testing"

	"fyne.io/fyne/v2/test"
)

func TestGreeting(t *testing.T) {
	out, in := makeUI()

	if out.Text != "Hello world!" {
		t.Error("Incorrect initial greeting")
	}

	test.Type(in, "Yale")
	if out.Text != "Hello Yale!" {
		t.Error("Incorrect initial greeting")
	}

}
