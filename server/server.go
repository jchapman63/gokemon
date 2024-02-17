package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jchapman63/gokemon/server/game"
	"github.com/jchapman63/gokemon/server/player"
	"github.com/jchapman63/gokemon/server/pokemon"
)

func Server() {
	server := &http.Server{
		Addr: ":8080",
	}
	// initialize a simple game environment
	tackle := pokemon.DamageMove{
		Name:  "tackle",
		Power: 10,
	}
	pika := &pokemon.Pokemon{
		Name: "pikachu",
		Hp:   100,
		Moves: []pokemon.DamageMove{
			tackle,
		},
	}
	bulbasaur := &pokemon.Pokemon{
		Name: "bulbasaur",
		Hp:   100,
		Moves: []pokemon.DamageMove{
			tackle,
		},
	}

	game := game.Game{
		Pokemon: []*pokemon.Pokemon{
			pika,
			bulbasaur,
		},
	}

	http.HandleFunc("/join", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var player player.Player

		err := decoder.Decode(&player)
		if err != nil {
			panic(err)
		}

		game.AddPlayerToMatch(&player)
	})

	http.HandleFunc("/addPokemonToPlayer", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var adder player.MonsterAdder
		err := decoder.Decode(&adder)
		if err != nil {
			panic(err)
		}
	})

	// allow players to choose an available monster
	http.HandleFunc("/getMonsters", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		pokemon := []string{
			pokemon.Bulbasaur.Name,
			pokemon.Gibble.Name,
			pokemon.Pika.Name,
			pokemon.Whooper.Name,
		}
		json.NewEncoder(w).Encode(pokemon)
	})

	// a simple attack as a demo
	// TODO: Handle arguments for which pokemon attacks and which pokemon gets attacked
	http.HandleFunc("/damage", func(w http.ResponseWriter, r *http.Request) {
		pika.Attack(bulbasaur, tackle)
	})

	// return digestable game state
	http.HandleFunc("/state", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(game)
	})

	// calls Game's method to check if game over
	http.HandleFunc("/isOver", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(game.IsGameOver())
	})

	fmt.Println("Server is listening localhost:8080")
	// why nil here? But, this will serve the app
	log.Fatal(server.ListenAndServe())

	if game.IsGameOver() {
		shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
		defer shutdownRelease()
		server.Shutdown(shutdownCtx)
	}
}

func localAttackTest(pika *pokemon.Pokemon, bulb *pokemon.Pokemon) {
	move := pika.Moves[0]
	pika.Attack(bulb, move)
	fmt.Print("bubla health: ", bulb.Hp)
}
