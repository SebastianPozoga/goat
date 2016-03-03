package main

import (
	"fmt"
	"github.com/goatcms/goat/src/command"
	"github.com/goatcms/goat/src/settings"
	"github.com/goatcms/goat/src/consts"
	"os"
	"strings"
)

func main() {
	settings, err := settings.NewSettings(consts.Version)
	if err!=nil {
		fmt.Println(err)
		os.Exit(1)
	}
	settings.Read()


	if len(os.Args) < 2 {
		command.HelpCommand(settings, os.Args)
		return
	}

	switch strings.ToLower(os.Args[1]) {
	case "workspace:new":
		command.NewWorkspaceCommand(settings, os.Args[2:])

	case "workspace:set":
		command.SetWorkspaceCommand(settings, os.Args[2:])

	case "workspace:get":
		command.GetWorkspaceCommand(settings, os.Args[2:])

	/*case "new":
	command.NewCommand(settings, os.Args[2:])*/

	case "version":
		command.VersionCommand(settings, os.Args[2:])

	case "help":
		command.HelpCommand(settings, os.Args[2:])

	default:
		fmt.Println("Error: Unknow command")
		os.Exit(1)
	}
}
