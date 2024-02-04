package player

import "github.com/jchapman63/server/pokemon"

// this will host the player struct and its data
type player struct {
	Name string
	// pointer for addressing ?
	pokemon []*pokemon.Pokemon
}
