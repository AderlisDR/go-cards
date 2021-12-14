package main

import "fmt"

func main() {
	cards := newDeck()
	cards.shuffle()

	hand, restOfCards := deal(cards, 7)

	fmt.Println("---My Hand of Cards---")
	hand.print()
	fmt.Println("\n---Rest of Cards---")
	restOfCards.print()
}
