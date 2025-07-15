package main

import (
	"fmt"

	"github.com/parkertr/deck"
)

func main() {
	fmt.Println("=== Deck of Cards Library Demo ===")

	// Create a new deck
	d := deck.NewDeck()
	fmt.Printf("Created new deck with %d cards\n", d.Size())

	// Show some cards
	fmt.Println("\nFirst 5 cards in order:")
	topCards, _ := d.PeekN(5)
	for i, card := range topCards {
		fmt.Printf("%d. %s (%s)\n", i+1, card.String(), card.ShortString())
	}

	// Shuffle the deck
	fmt.Println("\nShuffling the deck...")
	d.Shuffle()

	// Show the same positions after shuffle
	fmt.Println("First 5 cards after shuffle:")
	topCards, _ = d.PeekN(5)
	for i, card := range topCards {
		fmt.Printf("%d. %s (%s)\n", i+1, card.String(), card.ShortString())
	}

	// Deal some cards
	fmt.Println("\nDealing 5 cards:")
	hand, err := d.DealN(5)
	if err != nil {
		panic(err)
	}

	for i, card := range hand {
		color := "black"
		if card.IsRed() {
			color = "red"
		}
		faceCard := ""
		if card.IsFaceCard() {
			faceCard = " (face card)"
		}
		fmt.Printf("%d. %s - %s%s\n", i+1, card.ShortString(), color, faceCard)
	}

	fmt.Printf("\nCards remaining in deck: %d\n", d.Size())

	// Filter cards
	fmt.Println("\nFiltering for red cards in remaining deck:")
	redCards := d.Filter(func(c deck.Card) bool {
		return c.IsRed()
	})
	fmt.Printf("Found %d red cards\n", redCards.Size())

	// Count by suit
	fmt.Println("\nCard count by suit in remaining deck:")
	suitCounts := d.CountBySuit()
	suits := []deck.Suit{deck.Spades, deck.Hearts, deck.Diamonds, deck.Clubs}
	for _, suit := range suits {
		fmt.Printf("%s: %d cards\n", suit.String(), suitCounts[suit])
	}

	// Create a custom hand
	fmt.Println("\nCreating a custom hand (poker royal flush):")
	royalFlush := deck.NewEmptyDeck()
	royalFlush.AddCard(deck.NewCard(deck.Spades, deck.Ten))
	royalFlush.AddCard(deck.NewCard(deck.Spades, deck.Jack))
	royalFlush.AddCard(deck.NewCard(deck.Spades, deck.Queen))
	royalFlush.AddCard(deck.NewCard(deck.Spades, deck.King))
	royalFlush.AddCard(deck.NewCard(deck.Spades, deck.Ace))

	cards := royalFlush.Cards()
	for i, card := range cards {
		fmt.Printf("%d. %s\n", i+1, card.String())
	}

	// Sort the royal flush
	fmt.Println("\nSorting the royal flush:")
	royalFlush.Sort()
	cards = royalFlush.Cards()
	for i, card := range cards {
		fmt.Printf("%d. %s\n", i+1, card.String())
	}

	// Demonstrate reproducible shuffle
	fmt.Println("\nDemonstrating reproducible shuffle with seed:")
	deck1 := deck.NewDeck()
	deck2 := deck.NewDeck()

	deck1.ShuffleWithSeed(42)
	deck2.ShuffleWithSeed(42)

	fmt.Println("Both decks shuffled with seed 42:")
	cards1, _ := deck1.PeekN(3)
	cards2, _ := deck2.PeekN(3)

	fmt.Println("Deck 1 top 3:", cards1[0].ShortString(), cards1[1].ShortString(), cards1[2].ShortString())
	fmt.Println("Deck 2 top 3:", cards2[0].ShortString(), cards2[1].ShortString(), cards2[2].ShortString())
	fmt.Println("Should be identical!")
}
