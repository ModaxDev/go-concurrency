package main

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"time"
)

var seatingCapacity = 10
var arrivalRate = 100
var cutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

func main() {
	rand.Seed(time.Now().UnixNano())

	color.Yellow("The sleeping barber problem")
	color.Yellow("----------------------------")

	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	shop := BarberShop{
		ShopCapacity:    seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		ClientsChan:     clientChan,
		BarbersDoneChan: doneChan,
		Open:            true,
	}

	color.Green("Barber shop is open")
	color.Green("-------------------")

	shop.addBarber("Franck")
	shop.addBarber("Clement")
	shop.addBarber("Susan")
	shop.addBarber("Theo")
	shop.addBarber("Louis")
	shop.addBarber("Maman")

	shopClosing := make(chan bool)
	closed := make(chan bool)

	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		shop.closeShopForDay()
		closed <- true
	}()

	i := 1

	go func() {
		for {
			randomMillisecond := rand.Int() % (arrivalRate * 2)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randomMillisecond)):
				shop.addClient(fmt.Sprintf("Client %d", i))
				i++
			}
		}
	}()

	<-closed
}
