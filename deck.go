package main

// Deck represents a collection of units.
type Deck struct {
	MergedDeck []Unit          `json:"mergedDeck"`
	Base       [][]interface{} `json:"base"`
	Name       string          `json:"deckName"`
	Randomizer [][]string      `json:"randomizer"`
}

// AdvancedSet returns the set of advanced units for the given replay.
func (d *Deck) AdvancedSet() []string {
	return d.Randomizer[0]
}