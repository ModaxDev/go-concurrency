package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {
	// variables for bank balance
	var bankBalance int
	var balance sync.Mutex

	// print out the initial balance
	fmt.Printf("Initial balance: $%d", bankBalance)
	fmt.Println("----------------------------------------")

	// define weekly revenue
	incomes := []Income{
		{"Main job", 500},
		{"Gifts", 10},
		{"Part time job", 50},
		{"Investments", 100},
	}

	wg.Add(len(incomes))

	// loop through 52 and print out the balance; keep a running total
	for i, income := range incomes {
		go func(i int, income Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				balance.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balance.Unlock()

				fmt.Printf("On Week %d: $%d from %s\n", week, income.Amount, income.Source)
			}
		}(i, income)
	}
	wg.Wait()
	fmt.Println("----------------------------------------")

	// print out the final balance
	fmt.Printf("Final balance: $%d", bankBalance)
}
