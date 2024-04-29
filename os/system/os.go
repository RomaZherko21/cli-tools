package system

import (
	"fmt"
	"runtime"

	"example.com/os/helpers"
	"github.com/shirou/gopsutil/v3/host"
)

func ShowOsInfo() {
	fmt.Println(helpers.InfoMsg("Operating system:"), runtime.GOOS)

	hostInfo, err := host.Info()
	if err != nil {
		helpers.ExitWithErr("Can't get host info")
	}
	fmt.Println(helpers.InfoMsg("Host ID:"), hostInfo.HostID)
	fmt.Println(helpers.InfoMsg("Hostname:"), hostInfo.Hostname)
	fmt.Println(helpers.InfoMsg("Architecture:"), runtime.GOARCH)
	fmt.Println(helpers.InfoMsg("Architecture version:"), hostInfo.KernelVersion)
	fmt.Println(helpers.InfoMsg("Procs:"), hostInfo.Procs)

	fmt.Println(helpers.InfoMsg("CPU Info:"), fmt.Sprintf("NumCPU: %d, NumCgoCall: %d", runtime.NumCPU(), runtime.NumCgoCall()))
}
