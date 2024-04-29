package system

import (
	"fmt"
	"strconv"

	"example.com/os/helpers"
	"github.com/shirou/gopsutil/v3/cpu"
)

func ShowCPUInfo() {
	cpuInfo, err := cpu.Info()

	if err != nil {
		helpers.ExitWithErr("Can't get cpu info")
	}

	for _, v := range cpuInfo {
		fmt.Println(helpers.InfoMsg("Model name:"), v.ModelName)
		fmt.Println(helpers.InfoMsg("Mhz:"), v.Mhz)
	}

	logicalCores, err := cpu.Counts(true)
	if err != nil {
		helpers.ExitWithErr("Can't count logical cores")
	}

	physicalCores, err := cpu.Counts(false)
	if err != nil {
		helpers.ExitWithErr("Can't count physical cores")
	}

	fmt.Println(helpers.InfoMsg("Logical cores:"), strconv.Itoa(logicalCores))
	fmt.Println(helpers.InfoMsg("Physical cores:"), strconv.Itoa(physicalCores))
}
