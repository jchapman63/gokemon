package gameCalls

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jchapman63/gokemon/server/game"
)

var baseUrl = "localhost:8081"

func GameData() (*game.Game, error) {
	// get JSON game data
	respJSON, err := http.Get(baseUrl + "/state")
	if err != nil {
		fmt.Println("server not found")
		return nil, err
	}
	defer respJSON.Body.Close() // close resp body before function ends

	// object to unpack
	var game game.Game

	// read JSON resp body
	bodyJSON, err := io.ReadAll(respJSON.Body)
	if err != nil {
		fmt.Println("Error reading response body: ", err)
		return nil, err
	}
	if err := json.Unmarshal(bodyJSON, &game); err != nil {
		panic(err)
	}

	return &game, nil
}

// will later have parameters for the pokemon attacking (the client's mon)
// I could either pass game as parameter, or I could get a new game object... new is more up to date with the server...
// there is more abstraction I can do between these functions...
// if I am always going to be getting a new game struct, then I can have one function that takes the url as a parameter ?
// I may actually want to pass other params too though, so I will abstract atleast the RESP and JSON to struct ?
func BasicAttack() (*game.Game, error) {
	// question, will passing game here update the current game object????
	var game *game.Game
	game, err := jsonResponseToGameStruct(game, baseUrl+"/damage")
	if err != nil {
		return nil, err
	}

	return game, nil
}

// s, the struct to unpack into
// endpoint, the full api url
func jsonResponseToGameStruct(g *game.Game, endpoint string) (*game.Game, error) {
	respJSON, err := http.Get(endpoint)
	if err != nil {
		fmt.Print("Data Request Failed: ", err)
		return nil, err
	}

	bodyJSON, err := io.ReadAll(respJSON.Body)
	if err != nil {
		fmt.Println("Error reading json: ", err)
		return nil, err
	}
	if err := json.Unmarshal(bodyJSON, &g); err != nil {
		return nil, err
	}

	return g, nil
}
