package main

import (
	"fmt"
	"sync"
)

var waitGroup sync.WaitGroup

func main() {
	CSticks := make([]*ChopStick, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopStick)
	}

	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{i + 1, CSticks[i], CSticks[(i+1)%5], 0}
	}

	for i := 0; i < 5; i++ {
		waitGroup.Add(1)
		go philosophers[i].eat()
	}
	waitGroup.Wait()
}

type ChopStick struct{ sync.Mutex }

type Philosopher struct {
	id              int
	leftCS, rightCS *ChopStick
	eatTimes        int
}

var mutex sync.Mutex

func (p Philosopher) eat() {
	defer waitGroup.Done()
	for i := 0; i < 3; i++ {
		mutex.Lock()
		if (p.eatTimes) < 3 {
			p.leftCS.Lock()
			p.rightCS.Lock()
			mutex.Unlock()
			fmt.Printf("starting to eat %d\n", p.id)
			fmt.Printf("finishing eating %d\n", p.id)
			p.leftCS.Unlock()
			p.rightCS.Unlock()

		} else {
			mutex.Unlock()
		}
	}
}
