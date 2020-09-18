package main

import (
	"bank/pkg/bank/card"
	"bank/pkg/bank/types"
	"fmt"
)

func main() {
	//var val1 types.Money
	//val1 = 1
	cards := []types.Card{
		{
			PAN:     "5058 xxxx xxxx 8888",
			Balance: 5000,
			Active:  true,
		},
		{
			PAN:     "5058 xxxx xxxx 8888",
			Balance: 5000,
			Active:  false,
		},
	}
	fmt.Println(card.PaymentSources(cards))

}
