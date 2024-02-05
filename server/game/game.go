package game

import "github.com/jchapman63/server/pokemon"

type Game struct {
	Pokemon []*pokemon.Pokemon `json:"pokemon"`
}

func (g *Game) isGameOver() bool {
	for i := range g.Pokemon {
		if g.Pokemon[i].Hp <= 0 {
			return true
		}
	}
	return false
}
