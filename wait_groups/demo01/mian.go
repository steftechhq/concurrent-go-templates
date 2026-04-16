package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//Wait groups and barriers are two synchronization abstractions that work on groups of goroutines.
// We typically use wait groups to wait for a group of tasks to complete. On the other hand, we use barriers to synchronize many goroutines at a
// common point.

// Done()
// Decrements the wait group size counter by 1
// Wait()
// Blocks until the wait group counter size is 0
// Add(delta int)
// Increments the wait group size counter by delta

func main() {
	wg := sync.WaitGroup{}
	wg.Add(4)

	for i := 0; i < 4; i++ {
		go doWork(i, &wg)
	}
	wg.Wait()
	fmt.Println("All complete")
}

func doWork(id int, wg *sync.WaitGroup) {
	i := rand.Intn(5)
	time.Sleep(time.Duration(i) * time.Second)

	fmt.Println(id, "Done working after", i, "seconds")
	wg.Done()
}
