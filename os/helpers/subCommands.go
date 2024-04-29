package helpers

import (
	"os"
)

type SubCommands struct {
	Primary   string
	Secondary string
}

func getPrimaryCommandsList() map[string]bool {
	list := make(map[string]bool)

	list["info"] = true
	list["cpu"] = true
	list["disk"] = true
	list["proc"] = true

	return list
}

func getSecondaryCommandsList() map[string]bool {
	list := make(map[string]bool)

	list["get"] = true

	return list
}

func validation() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) == 0 {
		ExitWithErr("Not enough arguments. Type --help for help")
	}
}

func GetSubCommands() SubCommands {
	validation()

	primaryKeys := getPrimaryCommandsList()
	secondaryKeys := getSecondaryCommandsList()

	subCommands := SubCommands{}

	primary := os.Args[1]
	secondary := ""
	if len(os.Args) > 2 {
		secondary = os.Args[2]
	}

	_, ok := primaryKeys[primary]
	if ok {
		subCommands.Primary = primary
	}

	_, ok = secondaryKeys[secondary]
	if ok {
		subCommands.Secondary = secondary
	}

	return subCommands
}
