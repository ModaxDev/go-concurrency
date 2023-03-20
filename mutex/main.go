package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string) {
	defer wg.Done()

	msg = s

}

func main() {
	msg = "Hello, universe!"

	wg.Add(2)
	go updateMessage("Hello, Go!")
	go updateMessage("Hello, world!")
	wg.Wait()

	fmt.Println(msg)

}

/*
func updateMessage(s string, m *sync.Mutex) {
	defer wg.Done()
	m.Lock()
	msg = s
	m.Unlock()
}

func main() {
	msg = "Hello, universe!"

	var mutex sync.Mutex

	wg.Add(2)
	go updateMessage("Hello, Go!", &mutex)
	go updateMessage("Hello, world!", &mutex)
	wg.Wait()

	fmt.Println(msg)

}
*/
