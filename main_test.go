package main

import "testing"


// Test_sayHello function fails if name != "Bob"
func Test_sayHello(t *testing.T) {
	name := "Bob"
	want := "Hello Bob"

	if got := sayHello(name); got != want {
		t.Errorf("hello() = %q, want %q", got, want)
	}
}