package deck

import (
	"errors"
	"math/rand"
	"time"
)

// Deck represents a deck of playing cards
type Deck struct {
	cards []Card
}

// NewDeck creates a new standard 52-card deck
func NewDeck() *Deck {
	cards := make([]Card, 0, 52)
	suits := []Suit{Spades, Hearts, Diamonds, Clubs}
	ranks := []Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}

	for _, suit := range suits {
		for _, rank := range ranks {
			cards = append(cards, NewCard(suit, rank))
		}
	}

	return &Deck{cards: cards}
}

// NewEmptyDeck creates a new empty deck
func NewEmptyDeck() *Deck {
	return &Deck{cards: make([]Card, 0)}
}

// NewDeckFromCards creates a new deck from a slice of cards
func NewDeckFromCards(cards []Card) *Deck {
	deckCards := make([]Card, len(cards))
	copy(deckCards, cards)
	return &Deck{cards: deckCards}
}

// Size returns the number of cards in the deck
func (d *Deck) Size() int {
	return len(d.cards)
}

// IsEmpty returns true if the deck has no cards
func (d *Deck) IsEmpty() bool {
	return len(d.cards) == 0
}

// Cards returns a copy of the cards in the deck
func (d *Deck) Cards() []Card {
	cards := make([]Card, len(d.cards))
	copy(cards, d.cards)
	return cards
}

// Shuffle shuffles the deck using Fisher-Yates algorithm
func (d *Deck) Shuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := len(d.cards) - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
}

// ShuffleWithSeed shuffles the deck with a specific seed for reproducible results
func (d *Deck) ShuffleWithSeed(seed int64) {
	r := rand.New(rand.NewSource(seed))
	for i := len(d.cards) - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
}

// Deal deals one card from the top of the deck
func (d *Deck) Deal() (Card, error) {
	if d.IsEmpty() {
		return Card{}, errors.New("cannot deal from empty deck")
	}

	card := d.cards[0]
	d.cards = d.cards[1:]
	return card, nil
}

// DealN deals n cards from the top of the deck
func (d *Deck) DealN(n int) ([]Card, error) {
	if n < 0 {
		return nil, errors.New("cannot deal negative number of cards")
	}
	if n > len(d.cards) {
		return nil, errors.New("not enough cards in deck")
	}

	cards := make([]Card, n)
	copy(cards, d.cards[:n])
	d.cards = d.cards[n:]
	return cards, nil
}

// AddCard adds a card to the bottom of the deck
func (d *Deck) AddCard(card Card) {
	d.cards = append(d.cards, card)
}

// AddCards adds multiple cards to the bottom of the deck
func (d *Deck) AddCards(cards []Card) {
	d.cards = append(d.cards, cards...)
}

// InsertCard inserts a card at the specified position (0 = top)
func (d *Deck) InsertCard(card Card, position int) error {
	if position < 0 || position > len(d.cards) {
		return errors.New("invalid position")
	}

	d.cards = append(d.cards, Card{})
	copy(d.cards[position+1:], d.cards[position:])
	d.cards[position] = card
	return nil
}

// RemoveCard removes the first occurrence of the specified card
func (d *Deck) RemoveCard(card Card) bool {
	for i, c := range d.cards {
		if c.Suit == card.Suit && c.Rank == card.Rank {
			d.cards = append(d.cards[:i], d.cards[i+1:]...)
			return true
		}
	}
	return false
}

// Peek returns the top card without removing it from the deck
func (d *Deck) Peek() (Card, error) {
	if d.IsEmpty() {
		return Card{}, errors.New("cannot peek at empty deck")
	}
	return d.cards[0], nil
}

// PeekN returns the top n cards without removing them from the deck
func (d *Deck) PeekN(n int) ([]Card, error) {
	if n < 0 {
		return nil, errors.New("cannot peek at negative number of cards")
	}
	if n > len(d.cards) {
		return nil, errors.New("not enough cards in deck")
	}

	cards := make([]Card, n)
	copy(cards, d.cards[:n])
	return cards, nil
}

// Reset resets the deck to a full 52-card deck
func (d *Deck) Reset() {
	newDeck := NewDeck()
	d.cards = newDeck.cards
}

// Clear removes all cards from the deck
func (d *Deck) Clear() {
	d.cards = d.cards[:0]
}

// Contains checks if the deck contains a specific card
func (d *Deck) Contains(card Card) bool {
	for _, c := range d.cards {
		if c.Suit == card.Suit && c.Rank == card.Rank {
			return true
		}
	}
	return false
}

// CountBySuit returns the number of cards of each suit in the deck
func (d *Deck) CountBySuit() map[Suit]int {
	counts := make(map[Suit]int)
	for _, card := range d.cards {
		counts[card.Suit]++
	}
	return counts
}

// CountByRank returns the number of cards of each rank in the deck
func (d *Deck) CountByRank() map[Rank]int {
	counts := make(map[Rank]int)
	for _, card := range d.cards {
		counts[card.Rank]++
	}
	return counts
}

// Filter returns a new deck containing only cards that match the predicate
func (d *Deck) Filter(predicate func(Card) bool) *Deck {
	var filtered []Card
	for _, card := range d.cards {
		if predicate(card) {
			filtered = append(filtered, card)
		}
	}
	return NewDeckFromCards(filtered)
}

// Sort sorts the deck by suit first, then by rank
func (d *Deck) Sort() {
	// Simple bubble sort for demonstration - could use more efficient algorithm
	for i := 0; i < len(d.cards)-1; i++ {
		for j := 0; j < len(d.cards)-i-1; j++ {
			card1, card2 := d.cards[j], d.cards[j+1]

			// Sort by suit first, then by rank
			if card1.Suit > card2.Suit ||
				(card1.Suit == card2.Suit && card1.Rank > card2.Rank) {
				d.cards[j], d.cards[j+1] = d.cards[j+1], d.cards[j]
			}
		}
	}
}
