package main 

import(
		"fmt"
		"bufio"
		"os"
		"strings"
)

type Animal interface{
	Eat()
	Move()
	Speak()
}

type cow struct {
	food,locomotion,spoken string 
}

func (c cow) Eat() {
	fmt.Println(c.food)
}
func (c cow) Move() {
	fmt.Println(c.locomotion)
}
func (c cow) Speak() {
	fmt.Println(c.spoken)
}

type bird struct {
	food,locomotion,spoken string
}

func (b bird) Eat() {
	fmt.Println(b.food)
}
func (b bird) Move() {
	fmt.Println(b.locomotion)
}
func (b bird) Speak() {
	fmt.Println(b.spoken)
}

type snake struct {
	food,locomotion,spoken string 
}

func (s snake) Eat() {
	fmt.Println(s.food)
}
func (s snake) Move() {
	fmt.Println(s.locomotion)
}
func (s snake) Speak() {
	fmt.Println(s.spoken)
}

func main() {
	animal_map := make(map [string]Animal)

	c := cow{food: "grass",locomotion: "walk",spoken: "moo"}
	b := bird{food: "worms",locomotion: "fly",spoken: "peep"}
	s := cow{food: "mice",locomotion: "slither",spoken: "hsss"}

	fmt.Println("Enter X to exit")
	fmt.Println("To enter a new animal: >newanimal <animal name> <cow/bird/snake>")
	fmt.Println("To query about an animal: >query <animal name> <eat/move/speak>")

	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf(">")
		scanner.Scan()
		user_string := scanner.Text()

		if strings.Contains(user_string,"X") {
			break
		}

		if strings.Contains(user_string,"newanimal") {
			temp := strings.TrimPrefix(user_string,"newanimal ")
			if strings.Contains(temp,"cow") {
				name := strings.TrimSuffix(temp," cow")
				animal_map[name] = c
			}else if strings.Contains(temp,"bird") {
				name := strings.TrimSuffix(temp," bird")
				animal_map[name] = b
			}else if strings.Contains(temp,"snake") {
				name := strings.TrimSuffix(temp," snake")
				animal_map[name] = s
			}
		}else if strings.Contains(user_string,"query") {
			temp := strings.TrimPrefix(user_string,"query ")
			if strings.Contains(temp,"eat") {
				name := strings.TrimSuffix(temp," eat")
				animal_map[name].Eat()
			}else if strings.Contains(temp,"move") {
				name := strings.TrimSuffix(temp," move")
				animal_map[name].Move()
			}else if strings.Contains(temp,"speak") {
				name := strings.TrimSuffix(temp," speak")
				animal_map[name].Speak()
			}
		}
	}
}