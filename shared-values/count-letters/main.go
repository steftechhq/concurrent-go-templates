package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func countLetters(url string, frequency []int, mutex *sync.Mutex) {
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
			mutex.Lock()
			frequency[cIndex] += 1
			mutex.Unlock()
		}
	}
	fmt.Println("Completed:", url)
}

func countLettersSequental(url string, frequency []int) {
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

	// loadSequential()
	loadConcurrently()
}

func loadSequential() {
	start := time.Now()
	var frequency = make([]int, 26)
	for i := 1000; i <= 1220; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		fmt.Println("Get url ", i)
		countLettersSequental(url, frequency)
	}
	for i, c := range allLetters {
		fmt.Printf("%c-%d ", c, frequency[i])
	}
	elapsed := time.Since(start)
	fmt.Printf("\n\nElapsed time: %.3f seconds\n\n", elapsed.Seconds())

}
func loadConcurrently() {
	start := time.Now()
	mutex := sync.Mutex{}
	var frequency = make([]int, 26)
	for i := 1000; i <= 1220; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		fmt.Println("Get url ", i)
		go countLetters(url, frequency, &mutex)
	}

	for i := 0; i < 100; i++ {
		time.Sleep(100 * time.Millisecond)
		if mutex.TryLock() {
			for i, c := range allLetters {
				fmt.Printf("%c-%d ", c, frequency[i])
			}
			mutex.Unlock()
		} else {
			fmt.Println("Mutex already being used")
		}
	}

	// mutex.Lock()
	// for i, c := range allLetters {
	// 	fmt.Printf("%c-%d ", c, frequency[i])
	// }
	// mutex.Unlock()
	elapsed := time.Since(start)
	fmt.Printf("\n\nElapsed time: %.3f seconds\n\n", elapsed.Seconds())

	time.Sleep(10 * time.Second)

}
