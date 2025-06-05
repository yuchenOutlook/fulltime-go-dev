package main

import (
	"fmt"
	"unsafe"
)

var (
	floatVar32 float32 = 3.14
	flootVar64 float64 = 3.14
	complexVar complex128 = 1 + 2i
	runeVar   rune = 'A' // rune is an alias for int32
	intVar int32 = 42
	uintVar uint32 = 42
	boolVar bool = true
	stringVar string = "Hello, World!"
	byteVar byte = 'A' // byte is an alias for uint8
	unsafePointerVar unsafe.Pointer = nil // unsafe.Pointer is a special type for low-level programming
	safePointerVar *int = nil // pointer to an int, can be nil or point to an int value
)

type Player struct {
	name string
	health int
	attackPower float64
}

func (player Player) getHealth() int {
	return player.health
}

func getHealth(player Player) int {
	return player.health
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	otherNumbers := make([]int, 4)

	fmt.Printf("Numbers: %v\n", numbers)
	fmt.Printf("other numbers: %v\n", otherNumbers)

	numbers = append(numbers, 6, 7, 8)
	fmt.Printf("Updated Numbers: %v\n", numbers)

	users := make(map[string]int)
	users["Alice"] = 30
	users["Bob"] = 25
	fmt.Printf("Users: %+v\n", users)

	// Example usage of the custom types
	player := Player{
		name: "Hero",
		attackPower: 10.5,
		health: 100,
	}

	// using the method to get health, the method is defined on the Player type
	fmt.Printf("health: %d\n", player.getHealth())

	// Using the function to get health, the function is defined outside the Player type
	fmt.Printf("health: %d\n", getHealth(player))

	delete(users, "Alice")
	users["Charlie"] = 28
	users["Dave"] = 22

	Alice_age, ok := users["Alice"]

	if !ok {
		fmt.Println("Alice not found in the map")
	} else {
		fmt.Printf("Alice's age: %d\n", Alice_age)
	}

	for name, age := range users {
		fmt.Printf("Name: %s, Age: %d\n", name, age)
	}

}