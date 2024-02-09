package client

import (
	"fmt"

	"github.com/jchapman63/gokemon/client/cli"
	"github.com/jchapman63/gokemon/client/gameCalls"
	"github.com/jchapman63/gokemon/server"
)

func ClientStart() {
	var action string = cli.MainMenu()

	if action == "host" {
		// later will build a docker container
		server.Server()
	} else if action == "connect" {
		game, err := gameCalls.GameData()
		if err != nil {
			fmt.Println("Failed to connect: ", err)
		}

		// a "while" loop that goes until the game is over happens here.
		isOver, err := gameCalls.IsGameOver()
		if err != nil {
			fmt.Println("Failed to connect: ", err)
		}

		for isOver != true {
			// generate and get actions
			choice := cli.AttackMenu()

			// temporary print
			fmt.Println("json data")
			fmt.Println(game.Pokemon[0].Hp) // returns 100
			if choice == "tackle" {
				// call attack, it returns a game state -> which is the struct of interest
				game, err := gameCalls.BasicAttack()
				if err != nil {
					fmt.Println("failed attack called: ", err)
					return
				}
				fmt.Println("json data after attack")
				fmt.Println(game.Pokemon[1].Hp)
			} else if choice == "quit" {
				return
			}
		}
	}
}
