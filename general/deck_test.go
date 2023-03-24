package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected first card of King of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveDeckToFileAndLoadDeckFromFile(t *testing.T) {
	_ = os.Remove("_decktesting")
	d := newDeck()
	_ = d.saveDeckToFile("_decktesting")

	loadedDeck := loadDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(loadedDeck))
	}
	_ = os.Remove("_decktesting")
}
