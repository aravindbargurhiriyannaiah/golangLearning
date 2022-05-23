package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//Create a new type, deck, which is a slice of strings

type deck []string

//Receiver for a function
//Convention - Receiver name is the first letter or 2 letters of the name of the type
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//This function does not have a receiver because it is creating a new deck. There isn't anything to work with yet
func newDeck() deck {
	suits := []string{"Clubs", "Spades", "Hearts", "Diamonds"}
	values := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	cards := deck{}
	for _, suit := range suits {
		for _, value := range values {
			card := fmt.Sprintf("%v_of_%v", value, suit)
			cards = append(cards, card)
		}
	}
	return cards
}

func (d deck) deal(noOfCards int) (deck, deck) {
	return d[:noOfCards], d[noOfCards:]
}

func (d deck) toString() string {
	cards := []string(d)
	return strings.Join(cards, ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func readFromFile(fileName string) deck {
	byteSlice, error := ioutil.ReadFile(fileName)
	if error != nil {
		fmt.Println("Error: ", error)
		os.Exit(-1)
	}
	s := strings.Split(string(byteSlice), ",")
	return deck(s)
}

func (d deck) shuffle() {
	for i := range d {
		rand.Seed(time.Now().UnixNano())
		randomIndex := rand.Intn(len(d) - 1)
		d[i], d[randomIndex] = d[randomIndex], d[i]
	}
}
