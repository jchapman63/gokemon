package game

import (
	"github.com/jchapman63/gokemon/server/player"
	"github.com/jchapman63/gokemon/server/pokemon"
)

// THINKPOINT: check that there is not a memory problem here
type Game struct {
	Players []*player.Player   `json:"players"`
	Pokemon []*pokemon.Pokemon `json:"pokemon"`
}

// all pokemon in one player's party have fainted
func (g *Game) IsGameOver() bool {
	for i := range g.Pokemon {
		if g.Pokemon[i].Hp <= 0 {
			return true
		}
	}
	return false
}

func (g *Game) AddPlayerToMatch(p *player.Player) {
	g.Players = append(g.Players, p)
}
