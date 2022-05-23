package main

import (
	"os"
	"strings"
	"testing"
)

//Check if there are 52 cards
//Check if there are 13 cards of each shape
func TestNewDeck(t *testing.T) {
	d := newDeck()
	testDeck(d, t)
}

func testDeck(d deck, t *testing.T) {
	length := len(d)
	if len(d) != 52 {
		t.Errorf("Expected number of cards = 52. Actual number of cards = %v", length)
	}
	noOfSpades := 0
	noOfClubs := 0
	noOfDiamonds := 0
	noOfHearts := 0

	for i := range d {
		switch {
		case strings.Contains(d[i], "Spades"):
			noOfSpades++
		case strings.Contains(d[i], "Diamonds"):
			noOfDiamonds++
		case strings.Contains(d[i], "Hearts"):
			noOfHearts++
		case strings.Contains(d[i], "Clubs"):
			noOfClubs++
		default:
			t.Errorf("Unrecognized card shape/suit in deck")
		}
	}
	if noOfSpades != 13 {
		t.Errorf("Expected no. of Spades = 13. Actual number of Spade = %v", noOfSpades)
	}
	if noOfHearts != 13 {
		t.Errorf("Expected no. of Hearts = 13. Actual number of Hearts = %v", noOfHearts)
	}
	if noOfDiamonds != 13 {
		t.Errorf("Expected no. of Diamonds = 13. Actual number of Diamonds = %v", noOfDiamonds)
	}
	if noOfClubs != 13 {
		t.Errorf("Expected no. of Clubs = 13. Actual number of clubs = %v", noOfClubs)
	}
}

func TestDeal(t *testing.T) {
	d := newDeck()
	dealtCards, remainingCards := deck.deal(d, 4)
	if len(dealtCards) != 4 {
		t.Errorf("Expected number of cards = 4. Actual dealt cards = %v", len(dealtCards))
	}
	if len(remainingCards) != 48 {
		t.Errorf("Expected number of remaining cards = 47. Actual remaining cards = %v", len(remainingCards))
	}
}

func TestToString(t *testing.T) {
	d := newDeck()
	str := d.toString()
	noOfCommas := strings.Count(str, ",")
	if noOfCommas != 51 {
		t.Errorf("Expected number of commas = 51. Actual number = %v", noOfCommas)
	}
}

func TestShuffle(t *testing.T) {
	d := newDeck()
	d.shuffle()
	if len(d) != 52 {
		t.Errorf("Expected 52 cards after shuffling. Actual number of cards = %v", len(d))
	}
}

func TestSaveDeckAndReadDeckFromFile(t *testing.T) {
	testFileName := "_deckTesting"
	os.Remove(testFileName)

	deck := newDeck()
	deck.saveToFile(testFileName)

	loadedDeck := readFromFile(testFileName)
	if(len(loadedDeck) != len(deck)) {
		t.Errorf("Expected number of cards in deck = 52. Actual number of cards after loading from file = %v", len(loadedDeck))
	}
	testDeck(loadedDeck, t)
	os.Remove(testFileName)
}
