package system

import (
	"example.com/os/helpers"
)

func CommandHandlers(subCommands helpers.SubCommands, flags helpers.Options) {
	switch subCommands.Primary {
	case "info":
		ShowOsInfo()
	case "cpu":
		ShowCPUInfo()
	case "disk":
		ShowDiskInfo(flags.Path)
	case "proc":
		ShowProcessInfo()
	}
}
