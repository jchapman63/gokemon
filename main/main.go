package main

import (
	"fmt"

	"github.com/jchapman63/client"
)

// / Dev Doc: starting out, I will just ask for simple input to get the server started.  Later, a more refined CLI would be nice.
func main() {
	fmt.Printf("Welcome to Gokemon!\n\n")
	client.ClientStart()
}
