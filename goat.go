package main

import (
	"fmt"
	"github.com/goatcms/goat/src/command"
	"github.com/goatcms/goat/src/settings"
	"os"
	"strings"
)

func main() {
	settings, err := settings.NewSettings()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	settings.Read()

	if len(os.Args) < 2 {
		command.HelpCommand(settings, os.Args)
		return
	}

	switch strings.ToLower(os.Args[1]) {
	//workspace
	case "workspace:new":
		command.NewWorkspaceCommand(settings, os.Args[2:])

	case "workspace:set":
		command.SetWorkspaceCommand(settings, os.Args[2:])

	case "workspace:get":
		command.GetWorkspaceCommand(settings, os.Args[2:])

	case "workspace:cache:clean":
		command.CleanWorkspaceCacheCommand(settings, os.Args[2:])

	//generator
	case "generator:run":
		command.RunGeneratorCommand(settings, os.Args[2:])

	//packages
	case "packages:add":
		command.AddPackagCommand(settings, os.Args[2:])

	case "packages:update":
		command.UpdatePackagesCommand(settings, os.Args[2:])

	//resources
	case "resource:add":
		command.AddResourceCommand(settings, os.Args[2:])

	case "resource:model:add":
		command.AddModelResourceCommand(settings, os.Args[2:])

	//history
	case "history:build":
		command.BuildHistoryCommand(settings, os.Args[2:])

	//module
	case "module:add":
		command.AddModuleCommand(settings, os.Args[2:])

	//console
	case "version":
		command.VersionCommand(settings, os.Args[2:])

	case "help":
		command.HelpCommand(settings, os.Args[2:])

	default:
		fmt.Println("Error: Unknow command")
		os.Exit(1)
	}
}
