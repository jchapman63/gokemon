package player

import "github.com/jchapman63/gokemon/server/pokemon"

// this will host the player struct and its data
type Player struct {
	Name string `json:"player-name"`
	// pointer for addressing ?
	Pokemon []*pokemon.Pokemon `json:"player-pokemon"`

	Pokedex []*pokemon.Pokemon `json:"player-pokedex"`
}

func (p *Player) AddPokemon(pkmn *pokemon.Pokemon) {
	p.Pokemon = append(p.Pokemon, pkmn)
}
