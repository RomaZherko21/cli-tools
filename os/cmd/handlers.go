package main

import (
	"example.com/os/helpers"
	"example.com/os/system"
)

func commandHandlers(subCommands helpers.SubCommands, flags helpers.Options) {
	switch subCommands.Primary {
	case "info":
		system.ShowOsInfo()
	case "cpu":
		system.ShowCPUInfo()
	case "disk":
		system.ShowDiskInfo(flags.Path)
	}
}
