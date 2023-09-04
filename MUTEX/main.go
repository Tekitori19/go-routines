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
	// biến lưu số dư trong bank
	var bankBalance int

	var balance sync.Mutex

	// in ra giá trị ban đầu
	fmt.Printf("Initial account balance: $%d.00\n", bankBalance)

	//thu nhập hằng tuần
	incomes := []Income{
		{Source: "Main leg job", Amount: 500},
		{Source: "Bet", Amount: 50},
		{Source: "Gift", Amount: 10},
		{Source: "Investments", Amount: 100},
	}

	wg.Add(len(incomes))

	//số tiền kiếm đc qua 52 tuần
	for i, income := range incomes {

		go func(i int, income Income) {
			defer wg.Done()

			for week := 1; week <= 52; week++ {
				balance.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balance.Unlock()

				fmt.Printf("On week %d, you earned $%d.00 from %s\n", week, income.Amount, income.Source)
			}
		}(i, income)
	}

	wg.Wait()

	//in ra số dư cuối
	fmt.Printf("Final bank balance: $%d.00", bankBalance)
}
