package main

import (
	"fmt"
	"sync"
	"time"
)

//When a group of goroutines are suspended on Wait() and we call Signal(), we only wake up one of the goroutines.
// We have no control over which goroutine the system will resume, and we should assume that it can be any goroutine blocked
// on the condition variable’s Wait(). Using Broadcast(), we ensure that all suspended goroutines on the condition variable are resumed.

// game that has players waiting for everyone to join before the game begins.

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	playersInGame := 4
	for playerId := 0; playerId < 4; playerId++ {
		go playerHandler(cond, &playersInGame, playerId)
		time.Sleep(1 * time.Second)
	}

}

func playerHandler(cond *sync.Cond, playersRemaining *int, playerId int) {
	cond.L.Lock()
	fmt.Println(playerId, ": Connected")
	*playersRemaining--
	if *playersRemaining == 0 {
		cond.Broadcast()
	}
	for *playersRemaining > 0 {
		fmt.Println(playerId, ": Waiting for more players")
		cond.Wait() //release this mutex atomically when we call Wait()
	}
	cond.L.Unlock()
	//When all the other goroutines unblock from the Wait(), as a result of the Broadcast(),
	// they exit the condition checking loop and release the mutex.
	fmt.Println("All players connected. Ready player", playerId)
	//Game started
}
