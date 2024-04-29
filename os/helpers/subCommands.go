package helpers

import (
	"os"
)

type SubCommands struct {
	Primary   string
	Secondary string
}

func ShowHelp() {
	ExitWithErr("Not enough arguments. Type --help for help")
}

func SubCommandsValidation() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) == 0 {
		ShowHelp()
	}
}

func GetSubCommands() SubCommands {
	SubCommandsValidation()

	subCommands := SubCommands{}

	primary := os.Args[1]
	secondary := ""
	if len(os.Args) > 2 {
		secondary = os.Args[2]
	}

	if primary == "info" || primary == "cpu" || primary == "disk" {
		subCommands.Primary = primary
	}

	if secondary == "get" {
		subCommands.Secondary = secondary
	}

	return subCommands
}
