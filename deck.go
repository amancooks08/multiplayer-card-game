package main

import (
	"math/rand"
	"time"
)

// NewDeck creates a new deck of cards
func NewDeck() Deck {
	var deck Deck
	for _, rank := range ranks {
		for _, suit := range suits {
			card := Card{Rank: rank, Suit: suit}
			deck = append(deck, card)
		}
	}
	return deck
}

// Shuffle shuffles the deck
func (d *Deck) Shuffle() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range *d {
		j := rand.Intn(i + 1)
		(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
	}
}
