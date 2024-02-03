package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/jchapman63/gokemon/pokemon"
)

func main() {
	fmt.Print("Hello World")
	poke := &pokemon.Pokemon{
		Name: "pikachu",
		Hp:   100,
	}
	// test pokemon struct, print in HTML (HTML will not be a feature of this application)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// doesn't need to place string into there
		fmt.Fprint(w, "First monster: ", html.EscapeString(poke.Name))
	})

	fmt.Println("Server is listening on port 8081")
	// why nil here? But, this will serve the app
	log.Fatal(http.ListenAndServe(":8081", nil))
}
