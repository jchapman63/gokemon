package cli

// NOTE: Ultimately, this file will need to be cleaned up, the functionality of server interactions should be from client.go and
// creating the CLI should only be done here.
import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/nexidian/gocliselect"
)

// file design: function out all cli options to be called from a main loop in client.go
// Menus so far
// - MainMenu
// - AttackMenu
// - CreatePlayer
// - ChooseMonster
func MainMenu() string {
	menu := gocliselect.NewMenu("How would you like to play?")

	menu.AddItem("Host", "host")
	menu.AddItem("Connect", "connect")

	choice := menu.Display()
	return choice
}

func CreatePlayer() string {
	fmt.Print("Name your player: ")
	reader := bufio.NewReader(os.Stdin)
	// read up to and including delimiter
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	// remove the delimeter from the string
	input = strings.TrimSuffix(input, "\n")
	return input
}

func ChooseMonster() string {
	actionMenu := gocliselect.NewMenu("Choose A Pokemon!")
	actionMenu.AddItem("Pikachu", "pikachu")
	actionMenu.AddItem("Bulbasaur", "bulbasaur")
	actionMenu.AddItem("Gibble", "gibble")
	actionMenu.AddItem("Whooper", "whooper")

	actionChoice := actionMenu.Display()

	return actionChoice
}

// TODO: Add params of a pokemon's moves
func AttackMenu() string {
	actionMenu := gocliselect.NewMenu("Attack!")
	actionMenu.AddItem("Tackle", "tackle")
	actionMenu.AddItem("Quit", "quit")

	actionChoice := actionMenu.Display()

	return actionChoice
}
