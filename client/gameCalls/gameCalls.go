package gameCalls

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jchapman63/server/game"
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
