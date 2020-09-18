package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleTotal() {
	cards := []types.Card{
		{
			PAN:     "5058 xxxx xxxx 8888",
			Balance: 10_000_00,
			Active:  true,
		},
		{
			PAN:     "5058 xxxx xxxx 0000",
			Balance: -10_000_00,
			Active:  true,
		},
		{
			PAN:     "5058 xxxx xxxx 1111",
			Balance: 15_000_00,
			Active:  false,
		},
	}
	result := PaymentSources(cards)
	for _, v := range result {
		fmt.Println(v.Number)
	}

	// Output:
	//5058 xxxx xxxx 8888
}
