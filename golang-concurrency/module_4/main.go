package main

import (
	"fmt"
	"sync"
	"time"
)

type ChopS struct{ sync.Mutex }

type Philo struct {
	leftCS, rightCS *ChopS
	id              int
}

var (
	eatWgroup sync.WaitGroup
	// 5. The host allows no more than 2 philosophers to eat concurrently.
	host = make(chan int, 2)
)

func main() {
	// 1. There should be 5 philosophers sharing chopsticks,
	// with one chopstick between each adjacent pair of philosophers.
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	philosophers := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		// 3. The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
		// 6. Each philosopher is numbered, 1 through 5.
		philosophers[i] = &Philo{CSticks[i], CSticks[(i+1)%5], i + 1}
	}
	for i := 0; i < 5; i++ {
		eatWgroup.Add(1)
		go philosophers[i].eat()
	}
	eatWgroup.Wait()
}

func (p Philo) eat() {
	defer eatWgroup.Done()

	// 2. Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
	for i := 0; i < 3; i++ {
		// 4. In order to eat, a philosopher must get permission from a host which
		// executes in its own goroutine.
		host <- 1
		p.leftCS.Lock()
		p.rightCS.Lock()

		// When a philosopher starts eating (after it has obtained necessary locks)
		// it prints “starting to eat <number>” on a line by itself, where <number>
		// is the number of the philosopher.
		fmt.Printf("starting to eat %v \n", p.id)

		time.Sleep(time.Second)
		// When a philosopher finishes eating (before it has released its locks)
		// it prints “finishing eating <number>” on a line by itself, where <number>
		// is the number of the philosopher.
		fmt.Printf("finishing eating %v \n", p.id)

		p.rightCS.Unlock()
		p.leftCS.Unlock()

		<-host
	}
}
