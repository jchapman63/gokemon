package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"

	"github.com/jchapman63/gokemon/pokemon"
)

func main() {
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

	// monster info
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// doesn't need to place string into there
		fmt.Fprint(w, "First monster: ", html.EscapeString(pika.Name), "Health: ", html.EscapeString(strconv.Itoa(pika.Hp)))
		fmt.Fprint(w, "\n")
		fmt.Fprint(w, "Second monster: ", html.EscapeString(bulbasaur.Name), "Health: ", html.EscapeString(strconv.Itoa(bulbasaur.Hp)))
	})

	// a simple attack as a demo
	http.HandleFunc("/damage", func(w http.ResponseWriter, r *http.Request) {
		pika.Attack(bulbasaur, tackle)
	})

	fmt.Println("Server is listening localhost:8081")
	// why nil here? But, this will serve the app
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func localAttackTest(pika *pokemon.Pokemon, bulb *pokemon.Pokemon) {
	move := pika.Moves[0]
	pika.Attack(bulb, move)
	fmt.Print("bubla health: ", bulb.Hp)
}

// TIP: web page not refreshed? use shift + cmd + R
