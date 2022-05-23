package main

// import "fmt"

// "fmt"

func main() {
	cards := newDeck()
	// dealtCards, remainingCards := cards.deal(9)
	// noOfCards := len(dealtCards)
	// noOfCardsRemaining := len(remainingCards)
	// fmt.Printf("No. of cards dealt = %v. No. of cards remaining in deck = %v\n", noOfCards, noOfCardsRemaining)
	// fmt.Printf("Dealt cards = %v\n", dealtCards)
	// fmt.Printf("Remaining cards = %v\n", remainingCards)

	// cards.saveToFile("cards-app-file")
	// deckFromFile := readFromFile("cards-app-file")
	// fmt.Println(deckFromFile)
	cards.shuffle()
	cards.print()
}
