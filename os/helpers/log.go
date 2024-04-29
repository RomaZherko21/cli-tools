package helpers

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func InfoMsg(str string) string {
	return color.New(color.Bold, color.FgBlue).SprintFunc()(str)
}

func ErrorMsg(str string) string {
	return color.New(color.Bold, color.FgRed).SprintFunc()(str)
}

func ExitWithErr(msg string) {
	fmt.Println(ErrorMsg(msg))
	os.Exit(1)
}
