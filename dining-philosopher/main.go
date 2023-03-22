package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type Philosopher struct {
	name      string
	rightFork int
	leftFork  int
}

var philosophers = []Philosopher{
	{name: "Plato", rightFork: 0, leftFork: 4},
	{name: "Socrates", rightFork: 1, leftFork: 0},
	{name: "Aristotle", rightFork: 2, leftFork: 1},
	{name: "Pascal", rightFork: 3, leftFork: 2},
	{name: "Locke", rightFork: 4, leftFork: 3},
}

var hunger = 3 // number of times each philosopher eats
var eatTime = 1 * time.Second
var thinkTime = 3 * time.Second
var sleepTime = 1 * time.Second

var orderMutex = &sync.Mutex{}
var orderFinished = []string{}

func main() {
	fmt.Println("Dining Philosophers Problem")
	fmt.Println("-------------------------------")
	fmt.Println("The table is empty.")

	time.Sleep(sleepTime)

	dine()

	fmt.Println("-------------------------------")
	fmt.Println("The table is empty.")
	fmt.Printf("Order finished: %s", strings.Join(orderFinished, ", "))
}

func dine() {
	wg := sync.WaitGroup{}
	wg.Add(len(philosophers))

	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	for i := 0; i < len(philosophers); i++ {
		go diningProblem(philosophers[i], forks, seated, &wg)
	}

	wg.Wait()
}

func diningProblem(philosopher Philosopher, forks map[int]*sync.Mutex, seated *sync.WaitGroup, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("%s is sitting at the table.\n", philosopher.name)
	seated.Done()

	seated.Wait()

	for i := hunger; i > 0; i-- {

		if philosopher.leftFork > philosopher.rightFork {
			forks[philosopher.rightFork].Lock()
			fmt.Printf("takes the right fork %s.\n", philosopher.name)
			forks[philosopher.leftFork].Lock()
			fmt.Printf("takes the left fork %s.\n", philosopher.name)
		} else {
			forks[philosopher.leftFork].Lock()
			fmt.Printf("takes the left fork %s.\n", philosopher.name)
			forks[philosopher.rightFork].Lock()
			fmt.Printf("takes the right fork %s.\n", philosopher.name)
		}

		fmt.Printf("%s has both fork and is eating.\n", philosopher.name)
		time.Sleep(eatTime)

		fmt.Printf("%s is thinking.\n", philosopher.name)
		time.Sleep(thinkTime)

		forks[philosopher.leftFork].Unlock()
		forks[philosopher.rightFork].Unlock()

		fmt.Printf("%s put down the forks.\n", philosopher.name)
	}

	fmt.Printf("%s is done eating.\n", philosopher.name)
	fmt.Printf("%s left the table.\n", philosopher.name)
	orderMutex.Lock()
	orderFinished = append(orderFinished, philosopher.name)
	orderMutex.Unlock()
}
