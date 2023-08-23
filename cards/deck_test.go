package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	initialCard := d[0]
	if initialCard != "Ace of Spades" {
		t.Errorf("Expected first card 'Ace of Spades' but got %v", initialCard)
	}

	finalCard := d[len(d)-1]
	if finalCard != "Four of Clubs" {
		t.Errorf("Expected last card 'Four of Clubs' but got %v", finalCard)
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	d := newDeck()

	filename := "_decktesting"
	os.Remove(filename)

	d.saveToFile(filename)

	loadedDeck := newDeckFromFile(filename)

	expectedDeckSize := 16
	if len(loadedDeck) != expectedDeckSize {
		t.Errorf("Expected deck length of %v, but got %v", expectedDeckSize, len(loadedDeck))
	}

	os.Remove(filename)
}
