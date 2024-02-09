package game

import "github.com/jchapman63/gokemon/server/pokemon"

type Game struct {
	Pokemon []*pokemon.Pokemon `json:"pokemon"`
}

func (g *Game) IsGameOver() bool {
	for i := range g.Pokemon {
		if g.Pokemon[i].Hp <= 0 {
			return true
		}
	}
	return false
}
