package cli

// NOTE: Ultimately, this file will need to be cleaned up, the functionality of server interactions should be from client.go and
// creating the CLI should only be done here.
import (
	"github.com/nexidian/gocliselect"
)

// file design: function out all cli options to be called from a main loop in client.go
// Menus so far
// - MainMenu
// - AttackMenu
func MainMenu() string {
	menu := gocliselect.NewMenu("How would you like to play?")

	menu.AddItem("Host", "host")
	menu.AddItem("Connect", "connect")

	choice := menu.Display()
	return choice
}

// TODO: Add params of a pokemon's moves
func AttackMenu() string {
	actionMenu := gocliselect.NewMenu("Attack!")
	actionMenu.AddItem("Tackle", "tackle")
	actionMenu.AddItem("Quit", "quit")

	actionChoice := actionMenu.Display()

	return actionChoice
}
