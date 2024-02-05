package player

import "github.com/jchapman63/server/pokemon"

// this will host the player struct and its data
type Player struct {
	Name string `json:"player-name"`
	// pointer for addressing ?
	Pokemon []*pokemon.Pokemon `json:"player-pokemon"`
}
