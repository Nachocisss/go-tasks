package main

import (
	"fmt"
	"sync"
	"time"
)

// ChopS struct
type ChopS struct{ sync.Mutex }

//Philo struct
type Philo struct {
	leftCS, rightCS *ChopS
}

// host function
func host(c1, c2, c3, c4, c5 chan int, philos []*Philo, wg sync.WaitGroup) int {
	for {
		time.Sleep(10 * time.Millisecond)
		select {
		case <-c1:
			//				wg.Add(1)
			fmt.Println("c1")
			c1 <- 1

		case <-c2:
			//				wg.Add(1)
			fmt.Println("c2")
			c2 <- 1

		case <-c3:
			//				wg.Add(1)
			fmt.Println("c3")
			c3 <- 1

		case <-c4:
			//				wg.Add(1)
			fmt.Println("c4")
			c4 <- 1

		case <-c5:
			//				wg.Add(1)
			fmt.Println("c5")
			c5 <- 1

		default:
			return 1
		}
	}
}

// eat function
func eat(p Philo, number int, channel chan int, wg sync.WaitGroup) {
	for i := 0; i < 3; i++ {

		channel <- 1
		<-channel

		p.leftCS.Lock()
		p.rightCS.Lock()
		fmt.Println("starting to eat %d", number)
		p.rightCS.Unlock()
		p.leftCS.Unlock()
		fmt.Println("finishing eating %d", number)
	}
}

//main function
func main() {

	//create 5 ChopS Struct inside slice "CStrick"
	Csticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		Csticks[i] = new(ChopS)
	}

	//create 5 philos Struct inside slice "philos" and assign the corresponding chops

	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{Csticks[i], Csticks[(i+1)%5]}
	}
	var wg sync.WaitGroup
	// host goroutine
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)
	c4 := make(chan int)
	c5 := make(chan int)
	go host(c1, c2, c3, c4, c5, philos, wg)

	// 3 session's of eating for each philo
	go eat(*philos[0], 1, c1, wg)
	eat(*philos[1], 2, c2, wg)
	eat(*philos[2], 3, c3, wg)
	eat(*philos[3], 4, c4, wg)
	eat(*philos[4], 5, c5, wg)
}
