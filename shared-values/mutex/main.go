package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	m := sync.Mutex{}

	no := 100
	runWithWaitgroup(&no, &wg, &m)
	runWithSleep(&no, &m)

	m.Lock()
	fmt.Println("NO IS ", no)
	m.Unlock()
}

func incr(no *int, m *sync.Mutex) {
	tmp := 1000000
	for tmp > 0 {
		m.Lock()
		*no++
		m.Unlock()
		tmp--
	}
}

func decr(no *int, m *sync.Mutex) {
	tmp := 1000000
	for tmp > 0 {
		m.Lock()
		*no--
		m.Unlock()
		tmp--
	}
}

func runWithSleep(no *int, m *sync.Mutex) {

	go incr(no, m)
	go decr(no, m)

	time.Sleep(2 * time.Second)

}

func runWithWaitgroup(no *int, wg *sync.WaitGroup, m *sync.Mutex) {

	wg.Add(2)

	go func() {
		defer wg.Done()
		incr(no, m)
	}()

	go func() {
		defer wg.Done()
		decr(no, m)
	}()

	wg.Wait()

}
