package cli

// NOTE: Ultimately, this file will need to be cleaned up, the functionality of server interactions should be from client.go and
// creating the CLI should only be done here.
import (
	"fmt"

	"github.com/jchapman63/gokemon/client/gameCalls"
	"github.com/jchapman63/gokemon/server"
	"github.com/nexidian/gocliselect"
)

// file design: function out all cli options to be called from a main loop in client.go
// Menus so far
// - MainMenu
// - AttackMenu
// The menus should be agnostic to what their "funcitonality" is.  I just want to get strings back from them
func MainMenu() {
	menu := gocliselect.NewMenu("How would you like to play?")

	menu.AddItem("Host", "host")
	menu.AddItem("Connect", "connect")

	choice := menu.Display()

	if choice == "host" {
		// idealistically, choosing to host builds a docker container that hosts the application for the host player.
		// this way the CLI is free to select more options that send requests
		server.Server()
	} else if choice == "connect" {

		game, err := gameCalls.GameData()
		if err != nil {
			// TODO configure UI to handle error
			return
		}
		fmt.Println("json data")
		fmt.Println(game.Pokemon[0].Hp) // returns 100

		actionChoice := AttackMenu()

		if actionChoice == "tackle" {
			// call attack, it returns a game state -> which is the struct of interest
			game, err := gameCalls.BasicAttack()
			if err != nil {
				fmt.Println("failed attack called: ", err)
				return
			}
			fmt.Println("json data after attack")
			fmt.Println(game.Pokemon[1].Hp) // returns crazy negative num
		}
	}
}

// TODO: Add params
func AttackMenu() string {
	actionMenu := gocliselect.NewMenu("Attack!")
	actionMenu.AddItem("Tackle", "tackle")
	actionMenu.AddItem("Quit", "quit")

	actionChoice := actionMenu.Display()

	return actionChoice
}
