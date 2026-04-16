package main

import (
	"fmt"
	"sync"
	"time"
)

//When thinking about barriers, we can visualize our goroutines being in one of two possible states: either executing their task or suspended and
// waiting for others to catch up.

//A barrier that can be reused is sometimes called a cyclic barrier.

type Barrier struct {
	size      int
	waitCount int
	cond      *sync.Cond
}

func newBarrier(size int) *Barrier {
	condVar := sync.NewCond(&sync.Mutex{})
	return &Barrier{size, 0, condVar}
}
func (b *Barrier) Wait() {
	b.cond.L.Lock()
	b.waitCount += 1

	if b.waitCount == b.size {
		b.waitCount = 0
		b.cond.Broadcast()
	} else {
		b.cond.Wait()
	}

	b.cond.L.Unlock()
}

func workAndWait(name string, timeToWork int, barrier *Barrier) {
	start := time.Now()
	for i := 0; i < 5; i++ {
		fmt.Println(time.Since(start), name, "is running")
		time.Sleep(time.Duration(timeToWork) * time.Second)
		fmt.Println(time.Since(start), name, "is waiting on barrier")
		barrier.Wait()
	}
}

func main() {

	barrier := newBarrier(2)
	go workAndWait("Red", 4, barrier)
	go workAndWait("Blue", 10, barrier)
	time.Sleep(time.Duration(100) * time.Second)

}
