package system

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/process"
)

func ShowProcessInfo() {
	processes, _ := process.Processes()
	for _, process := range processes {
		name, _ := process.Name()
		status, _ := process.Status()
		fmt.Println(fmt.Sprintf("%v, %v, %v", name, process.Pid, status))
	}
}
