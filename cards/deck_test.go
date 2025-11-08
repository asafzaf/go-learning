package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()
	if len(deck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(deck))
	}

	if deck[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be 'Ace of Spades', but got %v", deck[0])
	}

	if deck[len(deck)-1] != "Four of Clubs" {
		t.Errorf("Expected last card to be 'Four of Clubs', but got %v", deck[len(deck)-1])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	err := deck.saveToFile("_decktesting")
	if err != nil {
		t.Errorf("Error saving deck to file: %v", err)
	}

	newDeck := newDeckFromFile("_decktesting")
	if len(newDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(newDeck))
	}

	os.Remove("_decktesting")
}
