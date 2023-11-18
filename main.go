package main

import "fmt"

// main function prints "Alice"
// Note: main_test.go looks for "Bob"
// so testing displays an error
func main() {
	msg := sayHello("Alice")
	fmt.Println(msg)
}


// formatted sayHello function
func sayHello(name string) string {
	return fmt.Sprintf("Hi %s", name)
}
