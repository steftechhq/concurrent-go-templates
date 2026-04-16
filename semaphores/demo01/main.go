package main

import (
	"fmt"
	"sync"
)

// They allow a fixed number of permits for concurrent executions to access shared resources.
//If we ensure with a mutex that only a single goroutine has exclusive access, we ensure with a
// semaphore that at most N goroutines have exclusive access.
//In fact, a mutex gives the same functionality as a semaphore with N having a value of 1

//New semaphore function: Creates a new semaphore with X permits

//Acquire permit function: A goroutine would take one permit from the semaphore. If none are available, the goroutine will suspend and wait until one becomes available.

//Release permit function: Releases one permit so a goroutine can use it again with the acquire function.

// Another way we can think about semaphores is that they provide a similar functionality as the wait and signal of a
//
//	condition variable, with the added benefit of recording a signal even if there is no goroutine waiting.
type Semaphore struct {
	permits int
	cond    *sync.Cond
}

func newSemaphore(n int) *Semaphore {
	return &Semaphore{
		permits: n,
		cond:    sync.NewCond(&sync.Mutex{}),
	}
}

func (rw *Semaphore) Acquire() {
	rw.cond.L.Lock()
	for rw.permits <= 0 {
		rw.cond.Wait()
	}
	rw.permits--
	rw.cond.L.Unlock()
}

func (rw *Semaphore) Release() {
	rw.cond.L.Lock()

	rw.permits++
	rw.cond.Signal()

	rw.cond.L.Unlock()
}

func main() {
	semaphore := newSemaphore(0)
	for i := 0; i < 50000; i++ {
		go doWork(semaphore)
		fmt.Println("Waiting for child goroutine")
		semaphore.Acquire()
		fmt.Println("Child goroutine finished")
	}
}

func doWork(semaphore *Semaphore) {
	fmt.Println("Work started")
	fmt.Println("Work finished")
	semaphore.Release()
}
