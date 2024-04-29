package main

import (
	"example.com/os/helpers"
	"example.com/os/system"
)

func main() {
	subCommands := helpers.GetSubCommands()
	flags := helpers.ParseFlags()

	system.CommandHandlers(subCommands, flags)
}
