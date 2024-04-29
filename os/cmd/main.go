package main

import (
	"example.com/os/helpers"
	"example.com/os/system"
)

func main() {
	subCommands := helpers.GetSubCommands()
	flags := helpers.ParseFlags()

	if subCommands.Primary == "info" {
		system.ShowOsInfo()
	}

	if subCommands.Primary == "cpu" {
		system.ShowCPUInfo()
	}

	if subCommands.Primary == "disk" {
		system.ShowDiskInfo(flags.Path)
	}

	// fmt.Println("Name:", *name)
	// fmt.Println("Age:", *age)
	// fmt.Println("Married:", *married)
	// fmt.Println("Duration:", *duration)
}
