/*
In Go (also known as Golang), channels are a powerful feature used for 
communicating and synchronizing between goroutines. A channel is a typed 
conduit through which values can be sent and received between goroutines. 

Channels have a type that specifies the type of data that can be sent over 
the channel.
*/


package main

import "fmt"

func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, j)
        // Do some work here...
        results <- j * 2
    }
}

func main() {
    numJobs := 5
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)

    for i := 1; i <= 3; i++ {
        go worker(i, jobs, results)
    }

    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    close(jobs)

    for r := 1; r <= numJobs; r++ {
        fmt.Printf("Result %d: %d\n", r, <-results)
    }
}

/*
In this example, the main goroutine creates two channels: a jobs channel and 
a results channel. The main goroutine also starts three worker goroutines using 
the go keyword. Each worker goroutine receives jobs from the jobs channel using 
the <-chan syntax and sends results to the results channel using the 
chan<- syntax.

The main goroutine then sends five jobs to the jobs channel and closes the 
channel to indicate that no more jobs will be sent. Finally, the main goroutine 
receives the results from the results channel and prints them to the console.

========================================================================
Channels are a powerful feature of Go that can be used for various 
synchronization and communication tasks. They provide a safe and efficient way 
of passing data between goroutines without the need for locks or other 
synchronization primitives. By using channels, Go programs can take advantage 
of concurrency and parallelism to achieve higher performance and scalability.
*/