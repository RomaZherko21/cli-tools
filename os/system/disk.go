package system

import (
	"fmt"

	"example.com/os/helpers"
	"github.com/shirou/gopsutil/v3/disk"
)

func ShowDiskInfo(path string) {
	info, err := disk.Usage(path)

	if err != nil {
		helpers.ExitWithErr("Can't get cpu info")
	}

	fmt.Println("he", info)

	fmt.Println(helpers.InfoMsg("Path:"), info.Path)
	fmt.Println(helpers.InfoMsg("Total:"), helpers.ReadableMemory(info.Total))
	fmt.Println(helpers.InfoMsg("Free:"), helpers.ReadableMemory(info.Free))
}
