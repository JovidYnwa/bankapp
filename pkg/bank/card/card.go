package card

import (
	"bank/pkg/bank/types"
)

//Calculating tottal amount of funds in all cards

func Total(cards []types.Card) types.Money {
	sum := types.Money(0)
	for _, card := range cards {
		if !card.Active {
			continue
		}
		if card.Balance <= 0 {
			continue
		}
		sum += card.Balance
	}
	return sum
}

func PaymentSources(cards []types.Card) []types.PaymentSource {
	cardsList := []types.PaymentSource{}
	for _, card := range cards {
		if !card.Active {
			continue
		}
		if card.Balance <= 0 {
			continue
		}
		cardsList = append(cardsList, types.PaymentSource{string(card.PAN), card.Balance})
	}
	return cardsList

}
