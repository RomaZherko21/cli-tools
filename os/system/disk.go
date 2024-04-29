package system

import (
	"fmt"
	"math"

	"example.com/os/helpers"
	"github.com/shirou/gopsutil/v3/disk"
)

func ShowDiskInfo(path string) {
	info, err := disk.Usage(path)

	if err != nil {
		helpers.ExitWithErr("Can't get disk info")
	}

	fmt.Println(helpers.InfoMsg("Path:"), info.Path)

	fmt.Println(helpers.InfoMsg("\nMemory:"))
	fmt.Println(helpers.InfoMsg(" Total:"), helpers.ReadableMemory(info.Total))
	fmt.Println(helpers.InfoMsg(" Free:"), helpers.ReadableMemory(info.Free))
	fmt.Println(helpers.InfoMsg(" Used:"), fmt.Sprintf("%v (%v%%)", helpers.ReadableMemory(info.Used), math.Ceil(info.UsedPercent*100)/100))

	fmt.Println(helpers.InfoMsg("\nInodes:"))
	fmt.Println(helpers.InfoMsg(" Total:"), info.InodesTotal)
	fmt.Println(helpers.InfoMsg(" Free:"), info.InodesFree)
	fmt.Println(helpers.InfoMsg(" Used:"), fmt.Sprintf("%v (%v%%)", info.InodesUsed, math.Ceil(info.InodesUsedPercent*100)/100))

}
