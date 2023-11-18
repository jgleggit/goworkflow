package main

import "fmt"

// main function prints "Alice"
// Note: main_test.go looks for "Bob"
// So running main_test.go displays an error.
//
func main() {
	msg := sayHello("Bob")
	fmt.Println(msg)
}


// formatted sayHello function
func sayHello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}
