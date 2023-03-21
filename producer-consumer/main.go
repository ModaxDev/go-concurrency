package main

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"time"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func pizzeria(pizzaMaker *Producer) {
	// keep track of the number of pizzas made
	var i = 0

	// run forever or until we receive a quit notification

	// try to make a pizza

	for {
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			// we tried to make a pizza ( we sent something to the data channel)
			case pizzaMaker.data <- *currentPizza:
			case quitChan := <-pizzaMaker.quit:
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}
	}
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Received order #%d \n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		total++

		fmt.Printf("Making pizza #%d. It will take %d seconds ..... \n", pizzaNumber, delay)
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("We ran out of ingredients for pizza #%d. We are sorry for the inconvenience. \n", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("The cook quit while making pizza #%d. \n", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("Pizza #%d is ready! \n", pizzaNumber)
		}

		return &PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}

	}

	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// print out a message
	color.Cyan("The pizza shop is open!")
	color.Cyan("----------------------------------------")

	// create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// run the producer in the background
	go pizzeria(pizzaJob)

	// create and run a consumer
	for i := range pizzaJob.data {
		if i.pizzaNumber <= NumberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("Order #%d is ready! \n", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("Order #%d failed! \n", i.pizzaNumber)
			}
		} else {
			color.Cyan("Done making pizzas!")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("Error: %s", err)
			}
		}
	}

	// print out the ending message
	color.Cyan("----------------------------------------")
	color.Cyan("Done for the day!")

	color.Cyan("We made %d pizzas today, but failed %d pizzas with %d attemps in total ", pizzasMade, pizzasFailed, total)

	switch {
	case pizzasFailed > 9:
		color.Red("We had a bad day!")
	case pizzasFailed >= 6:
		color.Yellow("We had not ok day!")
	case pizzasFailed >= 4:
		color.Yellow("We had a ok day!")
	case pizzasFailed >= 2:
		color.Yellow("We had pretty good day!")
	default:
		color.Green("We had a great day!")
	}
}
