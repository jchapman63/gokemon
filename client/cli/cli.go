package cli

// NOTE: Ultimately, this file will need to be cleaned up, the functionality of server interactions should be from client.go and
// creating the CLI should only be done here.
import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jchapman63/server"
	"github.com/nexidian/gocliselect"
)

var baseUrl = "http://localhost:8081"

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
		// JSON TESTING FROM SERVER ENDPOINT
		// Currently is just a bunch of bytes that get converted to string, I am sure I can marshal this into a usable struct
		respJSON, err := http.Get(baseUrl + "/state")
		if err != nil {
			fmt.Println("server not found")
			return
		}
		defer respJSON.Body.Close() // close resp body before function ends

		// TODO TEST JSON UNPACKING
		var data map[string]interface{}

		// read resp body
		bodyJSON, err := io.ReadAll(respJSON.Body)
		if err != nil {
			fmt.Println("Error reading response body: ", err)
			return
		}
		if err := json.Unmarshal(bodyJSON, &data); err != nil {
			panic(err)
		}
		fmt.Println("json data")
		fmt.Println(data)
		// END JSON TESTING FROM SERVER ENDPOINT

		// connect to the match
		resp, err := http.Get(baseUrl + "/")
		if err != nil {
			fmt.Println("server not found")
			return
		}
		defer resp.Body.Close() // close resp body before function ends

		// read resp body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body: ", err)
			return
		}
		fmt.Println("\n", string(body), "\n")

		actionMenu := gocliselect.NewMenu("Attack!")
		actionMenu.AddItem("Tackle", "tackle")
		actionMenu.AddItem("Quit", "quit")

		actionChoice := actionMenu.Display()

		if actionChoice == "tackle" {
			_, err := http.Get(baseUrl + "/damage")
			if err != nil {
				fmt.Println("error, disconnecting: ", err)
				return
			}

			// see tackle results
			resp, err := http.Get(baseUrl + "/")
			if err != nil {
				fmt.Println("server not found")
				return
			}
			defer resp.Body.Close() // close resp body before function ends
			// read resp body
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Error reading response body: ", err)
				return
			}
			fmt.Println("\n", string(body), "\n")
		}
	}
}
