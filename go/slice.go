package main
import (
	"fmt"
	//"strings"
	//"bufio"
	//"os"
	"sort"
	"strconv"
)
func main(){
	var slice []int;
	var askedNumber string;
	x := "X"
	for {
		fmt.Printf("Put a number (X to exit): ");
		fmt.Scan(&askedNumber);
		i1, _ := strconv.Atoi(askedNumber)

		// Chech if want to exit
		if askedNumber == x{
			break
		}
		slice = append(slice,i1)
		sort.Ints(slice)
		fmt.Printf("Your Slice: %v\n", slice)

	}
	fmt.Printf("Goodbye");

}