package main

import "fmt"

func main() {
	cards := []string{"Ace of Hearts", newCard()}

	cards = append(cards, "Six of Spades")

	for _, card := range cards {
		fmt.Println(card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
