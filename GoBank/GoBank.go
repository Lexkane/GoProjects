package main

import "fmt"

type Currency string

type Amount struct {
	Currency Currency
	Value    float32
}

type Balance map[Currency]float32

func (b *Balance) Add(amount Amount) *Balance {
	current, ok := (*b)[amount.Currency]
	if ok {
		(*b)[amount.Currency] = current + amount.Value
	} else {
		(*b)[amount.Currency] = amount.Value
	}
	return b
}

func main() {
	b := &Balance{Currency("USD"): 100.0}
	b = b.Add(Amount{Currency: Currency("USD"), Value: 5.0})

	fmt.Println("Balance: ", b)
}
