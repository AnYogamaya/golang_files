package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)

	fmt.Println("====== Hand ====")
	hand.print()
	fmt.Println("====== remaining cards ====")
	remainingCards.print()

	cards.saveToFile("my-card.txt")

	fmt.Println("====== From file ====")
	cards = newDeckFromFile("my-card.txt")
	cards.print()

	cards = newDeck()

	fmt.Println("====== Suffle cards ====")
	cards.shuffle()
	cards.print()

}
