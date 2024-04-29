package system

import (
	"fmt"
	"runtime"

	"example.com/os/helpers"
)

func ShowOsInfo() {
	fmt.Println(helpers.InfoMsg("Operating system:"), runtime.GOOS)
	fmt.Println(helpers.InfoMsg("Architecture:"), runtime.GOARCH)
	fmt.Println(helpers.InfoMsg("CPU Info:"), fmt.Sprintf("NumCPU: %d, NumCgoCall: %d", runtime.NumCPU(), runtime.NumCgoCall()))
}
