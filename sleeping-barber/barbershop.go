package main

import (
	"github.com/fatih/color"
	"time"
)

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarbersDoneChan chan bool
	ClientsChan     chan string
	Open            bool
}

func (shop *BarberShop) addBarber(barber string) {
	shop.NumberOfBarbers++

	go func() {
		isSleeping := false
		color.Yellow("%d goes to the waiting room to check for clients ", barber)

		for {
			if len(shop.ClientsChan) == 0 {
				color.Yellow("%d goes to sleep", barber)
				isSleeping = true
			}

			client, shopOpen := <-shop.ClientsChan
			if shopOpen {
				if isSleeping {
					color.Yellow("%s wakes up", barber)
					isSleeping = false
				}

				shop.cutHair(barber, client)
			} else {
				shop.sendBarberHome(barber)
				return
			}
		}
	}()
}

func (shop *BarberShop) cutHair(barber, client string) {
	color.Green("%s is cutting %s's hair", barber, client)
	time.Sleep(shop.HairCutDuration)
	color.Green("%s is done cutting %s's hair", barber, client)
}

func (shop *BarberShop) sendBarberHome(barber string) {
	color.Cyan("%s is going home", barber)
	shop.BarbersDoneChan <- true
}

func (shop *BarberShop) closeShopForDay() {
	color.Cyan("Closing shop for the day")
	close(shop.ClientsChan)
	shop.Open = false

	for i := 0; i < shop.NumberOfBarbers; i++ {
		<-shop.BarbersDoneChan
	}

	close(shop.BarbersDoneChan)
	color.Green("Barber shop is closed")
}

func (shop *BarberShop) addClient(client string) {
	color.Green("%s enters the shop", client)

	if shop.Open {
		select {
		case shop.ClientsChan <- client:
			color.Yellow("%s takes a seat in the waiting room", client)
		default:
			color.Red("%s leaves the shop because it is full", client)
		}
	} else {
		color.Red("%s leaves the shop because it is closed", client)
	}
}
