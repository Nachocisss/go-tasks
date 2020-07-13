package main

import (
	"fmt"
	//"strings"
	//"bufio"
	//"os"
)

//GenDisplaceFn function
func GenDisplaceFn(a, v, s float64) func(float64) float64 {
	fn := func(t float64) float64 {
		partOne := (a * t * t) / 2
		partTwo := v * t
		return partOne + partTwo + s
	}
	return fn
}

func main() {
	var a, v, s, t float64
	fmt.Printf("acceleration: ")
	fmt.Scan(&a)

	fmt.Printf("Initial velocity: ")
	fmt.Scan(&v)

	fmt.Printf("initial displacement: ")
	fmt.Scan(&s)

	fmt.Printf("time: ")
	fmt.Scan(&t)

	fn := GenDisplaceFn(a, v, s)
	fmt.Println("Final displacement:")
	fmt.Println(fn(t))

}
