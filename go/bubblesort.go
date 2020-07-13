package main

import (
	"fmt"
	//"strings"
	//"bufio"
	//"os"

	"strconv"
)

//Swap function to swap elements
func Swap(slice []int, index int) {
	aux := slice[index]
	slice[index] = slice[index-1]
	slice[index-1] = aux
}

//BubbleSort function to bubble sort slice
func BubbleSort(slice []int) {
	l := len(slice)
	i := l - 1
	if l > 1 {
		for i = l - 1; i > 0; i-- {
			if slice[i] < slice[i-1] {
				Swap(slice, i)
			}
		}
	}
}

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
		BubbleSort(slice)
	}
	fmt.Printf("Your Slice: %v\n", slice)

}
