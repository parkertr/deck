package deck

import "fmt"

// Suit represents a playing card suit
type Suit int

const (
	Spades Suit = iota
	Hearts
	Diamonds
	Clubs
)

// String returns the string representation of a suit
func (s Suit) String() string {
	switch s {
	case Spades:
		return "Spades"
	case Hearts:
		return "Hearts"
	case Diamonds:
		return "Diamonds"
	case Clubs:
		return "Clubs"
	default:
		return "Unknown"
	}
}

// Symbol returns the Unicode symbol for a suit
func (s Suit) Symbol() string {
	switch s {
	case Spades:
		return "♠"
	case Hearts:
		return "♥"
	case Diamonds:
		return "♦"
	case Clubs:
		return "♣"
	default:
		return "?"
	}
}

// Rank represents a playing card rank
type Rank int

const (
	Ace Rank = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

// String returns the string representation of a rank
func (r Rank) String() string {
	switch r {
	case Ace:
		return "Ace"
	case Two:
		return "Two"
	case Three:
		return "Three"
	case Four:
		return "Four"
	case Five:
		return "Five"
	case Six:
		return "Six"
	case Seven:
		return "Seven"
	case Eight:
		return "Eight"
	case Nine:
		return "Nine"
	case Ten:
		return "Ten"
	case Jack:
		return "Jack"
	case Queen:
		return "Queen"
	case King:
		return "King"
	default:
		return "Unknown"
	}
}

// Symbol returns the short symbol representation of a rank
func (r Rank) Symbol() string {
	switch r {
	case Ace:
		return "A"
	case Two:
		return "2"
	case Three:
		return "3"
	case Four:
		return "4"
	case Five:
		return "5"
	case Six:
		return "6"
	case Seven:
		return "7"
	case Eight:
		return "8"
	case Nine:
		return "9"
	case Ten:
		return "10"
	case Jack:
		return "J"
	case Queen:
		return "Q"
	case King:
		return "K"
	default:
		return "?"
	}
}

// Card represents a playing card
type Card struct {
	Suit Suit
	Rank Rank
}

// NewCard creates a new card with the given suit and rank
func NewCard(suit Suit, rank Rank) Card {
	return Card{Suit: suit, Rank: rank}
}

// String returns the full string representation of a card
func (c Card) String() string {
	return fmt.Sprintf("%s of %s", c.Rank.String(), c.Suit.String())
}

// ShortString returns the short representation of a card (e.g., "A♥")
func (c Card) ShortString() string {
	return fmt.Sprintf("%s%s", c.Rank.Symbol(), c.Suit.Symbol())
}

// IsRed returns true if the card is red (Hearts or Diamonds)
func (c Card) IsRed() bool {
	return c.Suit == Hearts || c.Suit == Diamonds
}

// IsBlack returns true if the card is black (Spades or Clubs)
func (c Card) IsBlack() bool {
	return c.Suit == Spades || c.Suit == Clubs
}

// IsFaceCard returns true if the card is a face card (Jack, Queen, King)
func (c Card) IsFaceCard() bool {
	return c.Rank == Jack || c.Rank == Queen || c.Rank == King
}
