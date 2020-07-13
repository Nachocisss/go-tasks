package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Animal struct
type Animal struct {
	food       string
	locomotion string
	noise      string
}

// Eat function
func Eat(a Animal) {
	fmt.Println(a.food)
}

// Move function
func Move(a Animal) {
	fmt.Println(a.locomotion)
}

// Speak function
func Speak(a Animal) {
	fmt.Println(a.noise)
}

func main() {
	//get's file name
	fmt.Printf("welcome: ")

	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}
	snake := Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	for {

		fmt.Printf(">")
		reader := bufio.NewReader(os.Stdin)
		name, _ := reader.ReadString('\n')
		n := strings.TrimSpace(name)

		// get every line and separate it
		separated := strings.Split(n, " ")

		if separated[0] == "cow" {

			if separated[1] == "eat" {
				Eat(cow)
			}
			if separated[1] == "move" {
				Move(cow)
			}
			if separated[1] == "speak" {
				Speak(cow)
			}
		}
		if separated[0] == "bird" {
			if separated[1] == "eat" {
				Eat(bird)
			}
			if separated[1] == "move" {
				Move(bird)
			}
			if separated[1] == "speak" {
				Speak(bird)
			}
		}
		if separated[0] == "snake" {
			if separated[1] == "eat" {
				Eat(snake)
			}
			if separated[1] == "move" {
				Move(snake)
			}
			if separated[1] == "speak" {
				Speak(snake)
			}
		}
	}
}
