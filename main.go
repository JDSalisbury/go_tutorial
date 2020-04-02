package main

func main() {
	cards := newDeck()
	cards.shuffle()
	hand, remainingDeck := deal(cards, 4)
	cards.count()
	hand.print()
	hand.count()
	remainingDeck.count()
	// remainingDeck.saveToFile("test.txt")
	// fileDeck := getDeckFromFile("test.txt")
	// fileDeck.print()
}
