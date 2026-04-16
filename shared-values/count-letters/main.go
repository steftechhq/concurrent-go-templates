package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

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

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func countLetters(url string, frequency []int) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	for _, b := range body {
		c := strings.ToLower(string(b))
		cIndex := strings.Index(allLetters, c)
		if cIndex >= 0 {
			frequency[cIndex] += 1
		}
	}
	fmt.Println("Completed:", url)
}

func countLettersConcurrent(url string, frequency []int) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	for _, b := range body {
		c := strings.ToLower(string(b))
		cIndex := strings.Index(allLetters, c)
		if cIndex >= 0 {
			frequency[cIndex] += 1
		}
	}
	fmt.Println("Completed:", url)
}

func main() {

	no := 100
	m := sync.Mutex{}
	go incr(&no, &m)
	go decr(&no, &m)

	time.Sleep(2 * time.Second)

	m.Lock()
	fmt.Println("NO IS ", no)
	m.Unlock()
	//loadSequential()
}

func loadSequential() {
	start := time.Now()
	var frequency = make([]int, 26)
	for i := 1000; i <= 1020; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		fmt.Println("Get url ", i)
		countLetters(url, frequency)
	}
	for i, c := range allLetters {
		fmt.Printf("%c-%d ", c, frequency[i])
	}
	elapsed := time.Since(start)
	fmt.Printf("\n\nElapsed time: %.3f seconds\n\n", elapsed.Seconds())

}
func loadConcurrently() {
	start := time.Now()
	var frequency = make([]int, 26)
	for i := 1000; i <= 1020; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		fmt.Println("Get url ", i)
		go countLetters(url, frequency)
	}
	for i, c := range allLetters {
		fmt.Printf("%c-%d ", c, frequency[i])
	}
	elapsed := time.Since(start)
	fmt.Printf("\n\nElapsed time: %.3f seconds\n\n", elapsed.Seconds())

}
