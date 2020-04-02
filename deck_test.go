package main

import (
	"os"
	"testing"
)

func TestNewDeckLength(t *testing.T) {
	d := newDeck()
	num := 52
	if len(d) != num {
		t.Errorf("Expected deck length of %v, got %v", num, len(d))
	}
}

func TestDeckDealing(t *testing.T) {
	d := newDeck()
	cardsDelt := 10
	hand, rDeck := deal(d, cardsDelt)

	if len(hand) != cardsDelt {
		t.Errorf("Expected Hand size of %v, got %v", cardsDelt, len(hand))
	}
	remainingDeckLength := len(d) - cardsDelt
	if len(rDeck) != remainingDeckLength {
		t.Errorf("Expected Remaing Deck length of %v, got %v", remainingDeckLength, len(rDeck))
	}

}

func TestDeckShuffle(t *testing.T) {
	ds := newDeck()
	us := newDeck()
	ds.shuffle()

	if ds[0] == us[0] && ds[1] == us[1] && ds[2] == us[2] {
		t.Errorf("Expected the decks to be more shuffled")
	}

}

func TestDeckReadAndWriteFunctionality(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := getDeckFromFile("_decktesting")

	if len(d) != len(loadedDeck) {
		t.Errorf("Expected deck lengths to match. Saved: %v, Loaded: %v", len(d), len(loadedDeck))
	}
	os.Remove("_decktesting")
}
