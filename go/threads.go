package main

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
)

//order function
func order(wg *sync.WaitGroup, slice []int) {
	fmt.Println("this goroutine will sort: %v", slice)
	sort.Ints(slice)
	wg.Done()
}

//main function
func main() {

	var slice []int
	var askedNumber string
	x := "X"
	for {
		fmt.Printf("Put a number (Up to 10 integers, X to finish): ")
		fmt.Scan(&askedNumber)
		i1, _ := strconv.Atoi(askedNumber)

		// Chech if want to exit
		if askedNumber == x {
			break
		}
		slice = append(slice, i1)
	}
	//divide slice
	quarter := int((len(slice) / 4) + 1)
	part1 := slice[0:quarter]
	part2 := slice[quarter : 2*quarter]
	part3 := slice[2*quarter : 3*quarter]
	part4 := slice[3*quarter:]

	// Create goroutine and sync
	var wg sync.WaitGroup
	wg.Add(4)
	go order(&wg, part1)
	order(&wg, part2)
	order(&wg, part3)
	order(&wg, part4)
	wg.Wait()
	fmt.Println("final: %v", slice)

}
