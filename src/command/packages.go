package command

import (
	"fmt"
	"github.com/goatcms/goat-core/workspace"
	"github.com/goatcms/goat/src/settings"
	"os"
)

func AddPackagCommand(s *settings.Settings, args []string) {
	w, err := workspace.ReadWorkspace(s.Workspace)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if len(args) < 1 {
		fmt.Println("You must defined dependencie id like 'github.com/goatcms/go-project' or 'github.com/goatcms/go-project.version' ")
		os.Exit(1)
	}
	for _, depId := range args {
		if _, err := w.Packages.CreteRecord(depId); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	if err := w.Write(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func UpdatePackagesCommand(s *settings.Settings, args []string) {
	w, err := workspace.ReadWorkspace(s.Workspace)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := w.Packages.Update(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
