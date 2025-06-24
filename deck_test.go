package deck

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := NewDeck()

	if deck.Size() != 52 {
		t.Errorf("Expected deck size to be 52, got %d", deck.Size())
	}

	if deck.IsEmpty() {
		t.Error("New deck should not be empty")
	}
}

func TestCard(t *testing.T) {
	card := NewCard(Hearts, Ace)

	if card.Suit != Hearts {
		t.Errorf("Expected suit to be Hearts, got %v", card.Suit)
	}

	if card.Rank != Ace {
		t.Errorf("Expected rank to be Ace, got %v", card.Rank)
	}

	if !card.IsRed() {
		t.Error("Hearts should be red")
	}

	if card.IsBlack() {
		t.Error("Hearts should not be black")
	}

	if card.IsFaceCard() {
		t.Error("Ace should not be a face card")
	}
}

func TestFaceCard(t *testing.T) {
	jack := NewCard(Spades, Jack)
	queen := NewCard(Hearts, Queen)
	king := NewCard(Diamonds, King)

	if !jack.IsFaceCard() {
		t.Error("Jack should be a face card")
	}

	if !queen.IsFaceCard() {
		t.Error("Queen should be a face card")
	}

	if !king.IsFaceCard() {
		t.Error("King should be a face card")
	}
}

func TestCardColors(t *testing.T) {
	spadeCard := NewCard(Spades, Ace)
	heartCard := NewCard(Hearts, Ace)
	diamondCard := NewCard(Diamonds, Ace)
	clubCard := NewCard(Clubs, Ace)

	if !spadeCard.IsBlack() || spadeCard.IsRed() {
		t.Error("Spades should be black")
	}

	if !heartCard.IsRed() || heartCard.IsBlack() {
		t.Error("Hearts should be red")
	}

	if !diamondCard.IsRed() || diamondCard.IsBlack() {
		t.Error("Diamonds should be red")
	}

	if !clubCard.IsBlack() || clubCard.IsRed() {
		t.Error("Clubs should be black")
	}
}

func TestDeal(t *testing.T) {
	deck := NewDeck()
	originalSize := deck.Size()

	_, err := deck.Deal()
	if err != nil {
		t.Errorf("Unexpected error dealing card: %v", err)
	}

	if deck.Size() != originalSize-1 {
		t.Errorf("Expected deck size to be %d after dealing, got %d", originalSize-1, deck.Size())
	}

	// Test dealing from empty deck
	emptyDeck := NewEmptyDeck()
	_, err = emptyDeck.Deal()
	if err == nil {
		t.Error("Expected error when dealing from empty deck")
	}
}

func TestDealN(t *testing.T) {
	deck := NewDeck()
	originalSize := deck.Size()

	cards, err := deck.DealN(5)
	if err != nil {
		t.Errorf("Unexpected error dealing 5 cards: %v", err)
	}

	if len(cards) != 5 {
		t.Errorf("Expected 5 cards, got %d", len(cards))
	}

	if deck.Size() != originalSize-5 {
		t.Errorf("Expected deck size to be %d after dealing 5 cards, got %d", originalSize-5, deck.Size())
	}

	// Test dealing more cards than available
	_, err = deck.DealN(100)
	if err == nil {
		t.Error("Expected error when dealing more cards than available")
	}
}

func TestShuffle(t *testing.T) {
	deck1 := NewDeck()
	deck2 := NewDeck()

	// Get initial order
	cards1 := deck1.Cards()
	cards2 := deck2.Cards()

	// Verify initial decks are identical
	for i := range cards1 {
		if cards1[i].Suit != cards2[i].Suit || cards1[i].Rank != cards2[i].Rank {
			t.Error("Initial decks should be identical")
		}
	}

	// Shuffle one deck
	deck1.Shuffle()
	shuffledCards := deck1.Cards()

	// Check that shuffle changed the order (very likely but not guaranteed)
	different := false
	for i := range cards1 {
		if shuffledCards[i].Suit != cards1[i].Suit || shuffledCards[i].Rank != cards1[i].Rank {
			different = true
			break
		}
	}

	// While not guaranteed, it's extremely unlikely that shuffle produces same order
	if !different {
		t.Log("Warning: Shuffle didn't change order (very unlikely but possible)")
	}

	// Verify deck still has same number of cards
	if deck1.Size() != 52 {
		t.Errorf("Shuffled deck should still have 52 cards, got %d", deck1.Size())
	}
}

func TestShuffleWithSeed(t *testing.T) {
	deck1 := NewDeck()
	deck2 := NewDeck()

	// Shuffle both with same seed
	deck1.ShuffleWithSeed(12345)
	deck2.ShuffleWithSeed(12345)

	cards1 := deck1.Cards()
	cards2 := deck2.Cards()

	// Should produce identical results
	for i := range cards1 {
		if cards1[i].Suit != cards2[i].Suit || cards1[i].Rank != cards2[i].Rank {
			t.Error("Decks shuffled with same seed should be identical")
		}
	}
}

func TestAddCard(t *testing.T) {
	deck := NewEmptyDeck()
	card := NewCard(Hearts, Ace)

	deck.AddCard(card)

	if deck.Size() != 1 {
		t.Errorf("Expected deck size to be 1, got %d", deck.Size())
	}

	if !deck.Contains(card) {
		t.Error("Deck should contain the added card")
	}
}

func TestRemoveCard(t *testing.T) {
	deck := NewDeck()
	originalSize := deck.Size()

	// Get a card from the deck
	cards := deck.Cards()
	cardToRemove := cards[0]

	removed := deck.RemoveCard(cardToRemove)
	if !removed {
		t.Error("Expected card to be removed")
	}

	if deck.Size() != originalSize-1 {
		t.Errorf("Expected deck size to be %d after removing card, got %d", originalSize-1, deck.Size())
	}

	if deck.Contains(cardToRemove) {
		t.Error("Deck should not contain removed card")
	}

	// Try to remove a card that doesn't exist
	nonExistentCard := NewCard(Hearts, Ace)
	deck.RemoveCard(nonExistentCard) // Remove it first
	removed = deck.RemoveCard(nonExistentCard)
	if removed {
		t.Error("Should not be able to remove non-existent card")
	}
}

func TestPeek(t *testing.T) {
	deck := NewDeck()
	originalSize := deck.Size()

	_, err := deck.Peek()
	if err != nil {
		t.Errorf("Unexpected error peeking at card: %v", err)
	}

	// Size should not change
	if deck.Size() != originalSize {
		t.Errorf("Deck size should not change after peeking, expected %d, got %d", originalSize, deck.Size())
	}

	// Peek at empty deck
	emptyDeck := NewEmptyDeck()
	_, err = emptyDeck.Peek()
	if err == nil {
		t.Error("Expected error when peeking at empty deck")
	}
}

func TestFilter(t *testing.T) {
	deck := NewDeck()

	// Filter for red cards
	redDeck := deck.Filter(func(c Card) bool {
		return c.IsRed()
	})

	if redDeck.Size() != 26 {
		t.Errorf("Expected 26 red cards, got %d", redDeck.Size())
	}

	// Verify all cards in filtered deck are red
	for _, card := range redDeck.Cards() {
		if !card.IsRed() {
			t.Error("Filtered deck should only contain red cards")
		}
	}
}

func TestSort(t *testing.T) {
	deck := NewDeck()
	deck.Shuffle() // Shuffle first to ensure it's not already sorted
	deck.Sort()

	cards := deck.Cards()

	// Verify cards are sorted by suit first, then by rank
	for i := 1; i < len(cards); i++ {
		prev, curr := cards[i-1], cards[i]

		if prev.Suit > curr.Suit {
			t.Error("Cards should be sorted by suit")
		}

		if prev.Suit == curr.Suit && prev.Rank > curr.Rank {
			t.Error("Cards of same suit should be sorted by rank")
		}
	}
}

func TestCountBySuit(t *testing.T) {
	deck := NewDeck()
	counts := deck.CountBySuit()

	expectedSuits := []Suit{Spades, Hearts, Diamonds, Clubs}
	for _, suit := range expectedSuits {
		if counts[suit] != 13 {
			t.Errorf("Expected 13 cards of suit %v, got %d", suit, counts[suit])
		}
	}
}

func TestCountByRank(t *testing.T) {
	deck := NewDeck()
	counts := deck.CountByRank()

	expectedRanks := []Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}
	for _, rank := range expectedRanks {
		if counts[rank] != 4 {
			t.Errorf("Expected 4 cards of rank %v, got %d", rank, counts[rank])
		}
	}
}

func TestCardString(t *testing.T) {
	card := NewCard(Hearts, Ace)
	expected := "Ace of Hearts"
	if card.String() != expected {
		t.Errorf("Expected %s, got %s", expected, card.String())
	}
}

func TestCardShortString(t *testing.T) {
	card := NewCard(Hearts, Ace)
	expected := "A♥"
	if card.ShortString() != expected {
		t.Errorf("Expected %s, got %s", expected, card.ShortString())
	}

	card = NewCard(Spades, King)
	expected = "K♠"
	if card.ShortString() != expected {
		t.Errorf("Expected %s, got %s", expected, card.ShortString())
	}
}
