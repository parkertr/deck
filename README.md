# Deck 🃏

A comprehensive Go library for working with standard playing card decks. Perfect for card games, simulations, and educational projects.

[![CI](https://github.com/parkertr/deck/actions/workflows/ci.yml/badge.svg)](https://github.com/parkertr/deck/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/parkertr/deck)](https://goreportcard.com/report/github.com/parkertr/deck)
[![GoDoc](https://godoc.org/github.com/parkertr/deck?status.svg)](https://godoc.org/github.com/parkertr/deck)

## Features

- 🃏 **Standard 52-card deck** with all suits and ranks
- 🔀 **Shuffling** with optional seed for reproducible results
- 🎯 **Deal cards** individually or in batches
- 👁️ **Peek** at cards without removing them
- ➕ **Add/Remove** cards dynamically
- 🔍 **Filter** cards by custom criteria
- 📊 **Sort** cards by suit and rank
- 📈 **Count** cards by suit or rank
- 🎨 **Rich formatting** with full names and Unicode symbols
- 🔴 **Color detection** (red/black cards)
- 👑 **Face card detection** (Jack, Queen, King)

## Installation

```bash
go get github.com/parkertr/deck
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/parkertr/deck"
)

func main() {
    // Create a new deck
    d := deck.NewDeck()
    fmt.Printf("New deck has %d cards\n", d.Size()) // 52

    // Shuffle the deck
    d.Shuffle()

    // Deal 5 cards
    hand, err := d.DealN(5)
    if err != nil {
        panic(err)
    }

    fmt.Println("Your hand:")
    for i, card := range hand {
        fmt.Printf("%d. %s (%s)\n", i+1, card.String(), card.ShortString())
    }
    // Output: 1. Ace of Hearts (A♥)
}
```

## API Documentation

### Creating Decks

```go
// Create a standard 52-card deck
deck := deck.NewDeck()

// Create an empty deck
emptyDeck := deck.NewEmptyDeck()

// Create deck from existing cards
cards := []deck.Card{deck.NewCard(deck.Hearts, deck.Ace)}
customDeck := deck.NewDeckFromCards(cards)
```

### Basic Operations

```go
// Check deck properties
size := deck.Size()        // Get number of cards
empty := deck.IsEmpty()    // Check if deck is empty

// Deal cards
card, err := deck.Deal()           // Deal one card
cards, err := deck.DealN(5)       // Deal multiple cards

// Peek at cards (without removing)
topCard, err := deck.Peek()        // Peek at top card
topCards, err := deck.PeekN(3)     // Peek at top 3 cards

// Add/Remove cards
deck.AddCard(card)                 // Add a card to the deck
removed := deck.RemoveCard(card)   // Remove specific card (returns bool)

// Deck management
deck.Reset()                       // Reset to full 52-card deck
deck.Clear()                       // Remove all cards
```

### Shuffling

```go
// Random shuffle
deck.Shuffle()

// Reproducible shuffle with seed
deck.ShuffleWithSeed(12345)

// Two decks with same seed will have identical order
deck1.ShuffleWithSeed(42)
deck2.ShuffleWithSeed(42)
// deck1 and deck2 now have identical card order
```

### Card Operations

```go
// Create cards
card := deck.NewCard(deck.Hearts, deck.Ace)

// Card properties
isRed := card.IsRed()          // true for Hearts/Diamonds
isBlack := card.IsBlack()      // true for Spades/Clubs
isFace := card.IsFaceCard()    // true for Jack/Queen/King

// Card representation
full := card.String()          // "Ace of Hearts"
short := card.ShortString()    // "A♥"

// Check if deck contains a card
exists := deck.Contains(card)
```

### Advanced Operations

```go
// Filter cards by criteria
redCards := deck.Filter(func(c deck.Card) bool {
    return c.IsRed()
})

faceCards := deck.Filter(func(c deck.Card) bool {
    return c.IsFaceCard()
})

// Sort deck (by suit, then rank)
deck.Sort()

// Count cards
suitCounts := deck.CountBySuit()    // map[Suit]int
rankCounts := deck.CountByRank()    // map[Rank]int

fmt.Printf("Hearts: %d\n", suitCounts[deck.Hearts])
fmt.Printf("Aces: %d\n", rankCounts[deck.Ace])
```

## Suits and Ranks

### Suits
- `deck.Spades` ♠
- `deck.Hearts` ♥
- `deck.Diamonds` ♦
- `deck.Clubs` ♣

### Ranks
- `deck.Ace` through `deck.King`
- Standard order: Ace, Two, Three, ..., Ten, Jack, Queen, King

## Complete Example

```go
package main

import (
    "fmt"
    "github.com/parkertr/deck"
)

func main() {
    // Create and shuffle a deck
    d := deck.NewDeck()
    d.Shuffle()

    // Deal a poker hand
    hand, _ := d.DealN(5)
    fmt.Println("Poker hand:")
    for _, card := range hand {
        color := "red"
        if card.IsBlack() {
            color = "black"
        }
        fmt.Printf("  %s (%s)\n", card.String(), color)
    }

    // Count remaining cards by suit
    fmt.Printf("\nRemaining cards: %d\n", d.Size())
    counts := d.CountBySuit()
    for _, suit := range []deck.Suit{deck.Spades, deck.Hearts, deck.Diamonds, deck.Clubs} {
        fmt.Printf("  %s: %d\n", suit.String(), counts[suit])
    }

    // Create a royal flush
    royalFlush := deck.NewEmptyDeck()
    royalFlush.AddCard(deck.NewCard(deck.Spades, deck.Ten))
    royalFlush.AddCard(deck.NewCard(deck.Spades, deck.Jack))
    royalFlush.AddCard(deck.NewCard(deck.Spades, deck.Queen))
    royalFlush.AddCard(deck.NewCard(deck.Spades, deck.King))
    royalFlush.AddCard(deck.NewCard(deck.Spades, deck.Ace))

    fmt.Println("\nRoyal Flush:")
    for _, card := range royalFlush.Cards() {
        fmt.Printf("  %s\n", card.ShortString())
    }
}
```

## Testing

Run the test suite:

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with race detection
go test -race ./...

# Or use the Makefile
make test
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Write tests for your changes
4. Ensure all tests pass (`make test`)
5. Commit your changes (`git commit -am 'Add amazing feature'`)
6. Push to the branch (`git push origin feature/amazing-feature`)
7. Open a Pull Request

## Development

This project uses:
- **Go 1.23+** for development
- **GitHub Actions** for CI/CD
- **golangci-lint** for code quality
- **Make** for build automation

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Example Projects

Check out the [example](./example/) directory for a complete demonstration of the library's capabilities.

---

Made with ❤️ for the Go community
