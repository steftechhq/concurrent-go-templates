package main

import (
	"sync"
	"testing"
)

func TestWaitGroup(t *testing.T) {
	no := 100

	var wg sync.WaitGroup
	m := sync.Mutex{}

	runWithWaitgroup(&no, &wg, &m)

	m.Lock()
	if no != 100 {
		t.Errorf("expected no to be 100, got %d", no)
	}
	m.Unlock()

}
