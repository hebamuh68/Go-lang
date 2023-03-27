package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Chops : Represents a Chopstick
type Chops struct {
	sync.Mutex
}

// Philo : Represents a philosopher
type Philo struct {
	id            int
	timesHasEaten int
	leftChop      *Chops
	rightChop     *Chops
}

// Host : Represents the host
type Host struct {
	sync.Mutex
	numOfPhiloEating int
}

// Eat : Simulates the eat action of a philosopher
// p: is the philosopher that wants to eat
// h : is the host to whom the philosopher must ask permission to eat
// wg: Uses a WaitGroup to indicate when the philosopher has eaten three times
func (p *Philo) Eat(h *Host, wg *sync.WaitGroup) {

	philoAteChannel := make(chan int)
	eatPermissionChannel := make(chan bool)
	//var timesEaten sync.Mutex

	// Loop three times for each philosopher
	// representing the three times that a philosopher
	// is allowed to eat
	for p.timesHasEaten < 3 {

		// Ask host permission to eat
		go h.CanPhiloEat(eatPermissionChannel)
		canEat := <-eatPermissionChannel

		if canEat {

			// Chooses which chopstick to pick up first randomly
			chopstickToPickup := rand.Intn(2)
			if chopstickToPickup == 0 {
				p.leftChop.Lock()
				p.rightChop.Lock()
			} else {
				p.rightChop.Lock()
				p.leftChop.Lock()
			}

			fmt.Printf(">> Eat(): Philosopher %d starting to eat... \n", p.id)
			time.Sleep(time.Second)

			fmt.Printf(">> Eat(): Philosopher %d finishing eating \n", p.id)

			p.timesHasEaten = p.timesHasEaten + 1
			fmt.Printf(">> Eat(): Philosopher %d has eaten %d times \n", p.id, p.timesHasEaten)

			p.leftChop.Unlock()
			p.rightChop.Unlock()

			// Tells the host that the philosopher finished eating
			// so that the host updates its internal counter
			// that tracks the number of philosophers eating
			go h.PhiloAte(philoAteChannel)

			// Waits for the host to indicate he's done updating
			// his internal state
			<-philoAteChannel

			time.Sleep(time.Second)

			if p.timesHasEaten == 3 {
				// Philosopher has ate three times, so tells
				// the main thread that is done
				fmt.Printf(">> Eat(): Philosopher %d is DONE ! \n", p.id)
				wg.Done()
			}

		}
	}
}

// CanPhiloEat :
func (h *Host) CanPhiloEat(c chan bool) {

	if h.numOfPhiloEating >= 2 {
		// There are two philosophers already eating
		// reject request from philosopher
		c <- false
	}

	// Zero or just one philosopher currently eating
	// so let's allow another philosopher to eat
	h.Lock()
	h.numOfPhiloEating = h.numOfPhiloEating + 1
	h.Unlock()

	// tells the philosopher he can start eating
	c <- true
}

// PhiloAte : Processes a message that indicates a philosopher has finished eating
// c: a channel to communicate with the philosopher
func (h *Host) PhiloAte(c chan int) {

	// One philosopher finished eating
	// so let's decrement counter to allow another philosopher start eating
	h.Lock()
	h.numOfPhiloEating = h.numOfPhiloEating - 1
	h.Unlock()
	c <- 1
}

func main() {

	// Initializes chop sticks
	csticks := make([]*Chops, 5)
	for i := 0; i < 5; i++ {
		csticks[i] = new(Chops)
	}

	// Initializes philosophers
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{i + 1, 0, csticks[i], csticks[(i+1)%5]}
	}

	// WaitGroup that indicates to the main thread when all the philosophers have finished eating
	var wg sync.WaitGroup

	// Initializes the Host
	h := new(Host)
	h.numOfPhiloEating = 0

	rand.Seed(2)

	// Starts a goroutine to perform the eat action for each philosopher
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philos[i].Eat(h, &wg)
	}

	// Waits until all philosophers have eaten three times
	wg.Wait()
	fmt.Println(">> Main(): All Philosophers ate three times, exiting...")
}