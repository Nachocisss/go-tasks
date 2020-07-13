package main

import (
	"fmt"
	"time"
)

// EXPLANATION: Here, we are creating 2 goroutine that calls de function plusOPne, diferents executions of it
// give diferent results for the race conditions, this mean that, to be concurrent, the process execute
// line instructions of diferents goroutine at the "same" time to improved perfomance. This order is not
// predictible and changes in every execution.
//
// We use de time sleep to make this condition more visible.

func plusOne(number *int) {

	for i := 0; i < 4; i++ {
		time.Sleep(1000 * time.Millisecond)
		*number = *number + 1
		fmt.Println(*number)
	}
}
func main() {
	n := 0
	go plusOne(&n)
	plusOne(&n)
}
