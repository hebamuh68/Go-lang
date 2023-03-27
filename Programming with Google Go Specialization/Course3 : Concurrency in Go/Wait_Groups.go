/*
In Go (also known as Golang), WaitGroups are a synchronization primitive 
used to wait for a group of goroutines to finish executing before continuing. 

A WaitGroup is a struct that contains an integer counter, which is incremented 
when a new goroutine is added to the group and decremented when a goroutine is 
finished. 

The WaitGroup's Wait() method can be used to block execution of the 
main goroutine until the counter reaches zero, indicating that all goroutines 
have finished.
*/

package main

import (
    "fmt"
    "sync"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()

    fmt.Printf("Worker %d starting\n", id)

    // Do some work here...
}

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go worker(i, &wg)
    }

    wg.Wait()
    fmt.Println("All workers finished")
}

/*
In this example, the main goroutine creates a new WaitGroup and then starts 
five worker goroutines. Each worker goroutine increments the WaitGroup's counter 
using the Add() method when it starts and decrements it using the Done() method 
when it finishes. 

Finally, the main goroutine waits for all worker goroutines to 
finish using the Wait() method before printing "All workers finished" to the 
console.

WaitGroups are useful for coordinating the execution of multiple goroutines and 
ensuring that they all finish before continuing. They can be used in a variety of 
scenarios, such as parallel processing, fan-out/fan-in concurrency patterns, and 
more.
*/