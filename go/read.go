package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type person struct {
	fname string
	lname string
}

func main() {
	//get's file name
	var slice []person
	fmt.Printf("Enter file name (include .txt): ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	n := strings.TrimSpace(name)

	//read the whole file
	f, _ := os.Open(n)
	rd := bufio.NewReader(f)
	for {
		// get every line and separate it
		fullname, err := rd.ReadString('\n')
		separated := strings.Split(fullname, " ")

		//make struct and append to slice
		guy := person{fname: separated[0], lname: separated[1]}
		slice = append(slice, guy)

		//condition of exit (End of File)
		if err == io.EOF {
			break
		}
	}
	fmt.Printf("Your Slice: %v", slice)
}
