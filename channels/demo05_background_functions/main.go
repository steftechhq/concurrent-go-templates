package main

import "fmt"

//execute functions concurrently in the background and then collect their results once they finish

func findFactors(number int) []int {
	result := make([]int, 0)
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			result = append(result, i)
		}
	}
	return result
}

// If the first findFactors() call is not yet finished, the reading from the channel will block the main goroutine until we have the results.
func main() {
	resultCh := make(chan []int)

	go func() {
		resultCh <- findFactors(3419110721)
	}()
	fmt.Println(findFactors(3419110721))
	fmt.Println(<-resultCh)
}
