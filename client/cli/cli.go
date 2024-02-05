package cli

import (
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
