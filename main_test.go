package main

import "testing"

func TestGreeting(t *testing.T) {
	got := greeting()
	if got == "" {
		t.Error("greeting is empty")
	}
}
