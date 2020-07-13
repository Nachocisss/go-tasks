package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main(){
	var found bool
    fmt.Printf("Put a String: ")
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')

	if (strings.Index(strings.ToLower(str), "i") == 0) && (strings.Index(strings.ToLower(str), "a") != -1) && (strings.Index(strings.ToLower(str), "n") == len(str)-3){
		found = true
	} else{
		found = false
	}
	if (found){
		fmt.Println("Found!")
	} else{
		fmt.Println("Not Found!")
	}	
}