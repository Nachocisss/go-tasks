package main

import (
	"fmt"
	"sync"
)

type Chopstick struct {
	sync.Mutex
}

type Philosopher struct {
	id                            int
	leftChopStick, rightChopStick *Chopstick
}

func (philosopher Philosopher) eat(host *Host, wait *sync.WaitGroup) {
	if host.canEat() {
		philosopher.leftChopStick.Lock()
		philosopher.rightChopStick.Lock()
		fmt.Printf("starting to eat %d\n", philosopher.id)

		fmt.Printf("finishing eating %d\n", philosopher.id)
		philosopher.leftChopStick.Unlock()
		philosopher.rightChopStick.Unlock()
		host.numberOfEatingPhilosophers--
		wait.Done()
	}
}

type Host struct {
	chopsticks []*Chopstick
	numberOfEatingPhilosophers int
}

func (host Host) canEat() bool {
	if host.numberOfEatingPhilosophers < 2 {
		host.numberOfEatingPhilosophers++
		return true
	} else {
		return false
	}
}

func main() {
	wait := sync.WaitGroup{}
	chopsticks := make([]*Chopstick, 5)

	for i := 0; i < 5; i++ {
		chopsticks[i] = new(Chopstick)
	}

	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{i + 1, chopsticks[i], chopsticks[(i+1)%5]}
	}

	host := Host{chopsticks: chopsticks}

	for _, philosopher := range philosophers {
		for i := 0; i < 3; i++ {
			wait.Add(1)
			go philosopher.eat(&host, &wait)
		}
	}
	wait.Wait()
}
