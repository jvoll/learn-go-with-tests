package main

// NewInMemoryPlayerStore spins up an empty player store
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

// InMemoryPlayerStore keeps palyer data in memory
type InMemoryPlayerStore struct {
	store map[string]int
}

// RecordWin updates a player's win record (or adds a record if non-existent)
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

// GetPlayerScore returns the score for a given player
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}
