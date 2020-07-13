package main

import (
	"fmt"
	//"strings"
	"bufio"
	"encoding/json"
	"os"
)

func main() {
	fmt.Printf("Enter a name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')

	fmt.Printf("Enter an Adress: ")
	address, _ := reader.ReadString('\n')
	m := map[string]string{name: address}
	barr, _ := json.Marshal(m)
	fmt.Println(string(barr))
}
