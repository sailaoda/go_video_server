package main

import (
	"testing"
)

func TestPrint(t *testing.T) {
	res := Printlto20()
	if res != 210 {
		t.Errorf("Return value not valid")
	}
}
