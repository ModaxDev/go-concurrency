package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup

	words := []string{"Hello", "World", "WorldWorld!", "WorldWorldWorldWorld"}

	wg.Add(len(words))

	for _, word := range words {
		go printSomething(word, &wg)
	}

	wg.Wait()
}
