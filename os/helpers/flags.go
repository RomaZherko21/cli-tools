package helpers

import (
	"github.com/jessevdk/go-flags"
)

type Options struct {
	Path string `long:"path" description:"Folder path" default:"/"`
	// Port    int    `short:"p" long:"port" description:"Port to listen on" default:"8080"`
	// Config  string `short:"c" long:"config" description:"Path to the configuration file"`
}

func ParseFlags() Options {
	var opts Options
	_, err := flags.Parse(&opts)
	if err != nil {
		ExitWithErr(err.Error())
	}

	return opts
}
