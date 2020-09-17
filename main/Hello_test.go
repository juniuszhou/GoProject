package main

import "testing"

func TestPrintError(t *testing.T) {
	result := PrintError()
	if result != 5 {
		t.Error("result from print error incorrect.")
	}
}
