/*
In the code above, we first demonstrate the race condition by executing incrementCounter and decrementCounter concurrently in a loop. We create 100 goroutines that call these functions, and we let the program sleep for one second before printing the final value of counter. As expected, the final value of counter is unpredictable and may be different every time we run the program.

Next, we fix the race condition by using a mutex to ensure that only one goroutine can access and modify counter at any given time. We create a new anonymous function for each goroutine that acquires the mutex, calls the corresponding function, and releases the mutex. Again, we create 100 goroutines and let the program sleep for one second before printing the final value of counter. This time, the final value of counter is always 0, indicating that the race condition has been fixed.

Note that in this example, we use a time.Sleep call to wait for the goroutines to finish before printing the final value of counter. In practice, it is better to use synchronization mechanisms such as wait groups or channels to wait for the goroutines to finish.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

var counter int

func incrementCounter() {
	counter++
}

func decrementCounter() {
	counter--
}

func main() {
	// Demonstrate the race condition
	fmt.Println("Demonstrating the race condition...")
	for i := 0; i < 100; i++ {
		go incrementCounter()
		go decrementCounter()
	}
	time.Sleep(1 * time.Second)
	fmt.Printf("Final value of counter (race condition): %d\n", counter)

	// Fix the race condition using a mutex
	var mutex sync.Mutex
	counter = 0
	fmt.Println("Fixing the race condition using a mutex...")
	for i := 0; i < 100; i++ {
		go func() {
			mutex.Lock()
			incrementCounter()
			mutex.Unlock()
		}()
		go func() {
			mutex.Lock()
			decrementCounter()
			mutex.Unlock()
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Printf("Final value of counter (fixed with mutex): %d\n", counter)
}

