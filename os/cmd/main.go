package main

import (
	"example.com/os/helpers"
)

func main() {
	subCommands := helpers.GetSubCommands()
	flags := helpers.ParseFlags()

	commandHandlers(subCommands, flags)
}
