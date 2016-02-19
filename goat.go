package main

import (
	"fmt"
	"github.com/goatcms/goat/src/command"
	"os"
)

const (
	version = "0.0.1"
)

func main() {
	if len(os.Args) < 2 {
		command.HelpCommand(nil, version)
		return
	}

	switch os.Args[1] {
	case "new":
		command.NewCommand(os.Args[2:])

	case "version":
		command.VersionCommand(os.Args[2:], version)

	case "help":
		command.HelpCommand(os.Args[2:], version)

	default:
		fmt.Println("Error: Unknow command")
		os.Exit(1)
	}
}
