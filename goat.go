package main

import (
	//	"flag"
	"fmt"
	"os"
	"github.com/goatcms/goat/src/command"
)

var (
	// contain configuration
	config map[string]string
)

const (
	// The version of this tool
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
		consoleError("Unknow command")
	}

}

func consoleError(msg string) {
	fmt.Println("Error: " + msg)
	os.Exit(1)
}
