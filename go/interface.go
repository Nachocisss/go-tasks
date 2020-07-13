package main

import (
	"fmt"
)

//Animal interface
type Animal interface {
	Eat()
	Move()
	Speak()
}

//----------- Types-------------
// cow scruct
type cow struct {
}

// bird scruct
type bird struct {
}

// snake scruct
type snake struct {
}

// ------------- COW----------------------
// Eat cow
func (c cow) Eat() {
	fmt.Println("grass")
}

// Move cow
func (c cow) Move() {
	fmt.Println("walk")
}

// Speak cow
func (c cow) Speak() {
	fmt.Println("moo")
}

// ------------- bird----------------------
// Eat bird
func (c bird) Eat() {
	fmt.Println("worms")
}

// Move bird
func (c bird) Move() {
	fmt.Println("fly")
}

// Speak bird
func (c bird) Speak() {
	fmt.Println("peep")
}

// ------------- snake----------------------
// Eat snake
func (c snake) Eat() {
	fmt.Println("mice")
}

// Move snake
func (c snake) Move() {
	fmt.Println("slither")
}

// Speak snake
func (c snake) Speak() {
	fmt.Println("hsss")
}

func main() {

	//var farm map[string]Animal
	farm := make(map[string]Animal)

	for {
		fmt.Print("> ")

		choice := ""
		name := ""
		species := ""
		fmt.Scan(&choice, &name, &species)

		switch choice {
		case "newanimal":
			switch species {
			case "cow":
				spe := cow{}
				farm[name] = spe
				fmt.Println("Created it!")

			case "bird":
				spe := bird{}
				farm[name] = spe
				fmt.Println("Created it!")

			case "snake":
				spe := snake{}
				farm[name] = spe
				fmt.Println("Created it!")
			}

		case "query":
			switch species {
			case "eat":
				a := farm[name]
				a.Eat()
			case "move":
				a := farm[name]
				a.Move()
			case "speak":
				a := farm[name]
				a.Speak()
			}
		}

	}
}
