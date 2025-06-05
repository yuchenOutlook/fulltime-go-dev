package main

import "fmt"

var (
	firstName  = "John"
	lastName = "Doe"
)

const (
	version = 1
)

func main() {
	version := 1
	fmt.Println("Hello, World!", firstName, lastName, version)
}