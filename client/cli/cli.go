package cli

import (
	"fmt"

	"github.com/jchapman63/server"
	"github.com/nexidian/gocliselect"
)

func MainMenu() {
	menu := gocliselect.NewMenu("How would you like to play?")

	menu.AddItem("Host", "host")
	menu.AddItem("Connect", "connect")

	choice := menu.Display()

	if choice == "host" {
		server.Server()
	} else if choice == "connect" {
		fmt.Print("Not yet implemented")
	}
}
