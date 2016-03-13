package command

import (
	"fmt"
	"github.com/goatcms/goat-core/scaffolding"
	"github.com/goatcms/goat-core/workspace"
	"github.com/goatcms/goat/src/settings"
	"os"
)

func AddModuleCommand(s *settings.Settings, args []string) {
	if len(args) < 1 {
		fmt.Println("Module destination path is required")
		os.Exit(1)
	}
	if len(args) < 2 {
		fmt.Println("Module url is required")
		os.Exit(1)
	}

	dest := args[0]
	depId := args[1]

	w, err := workspace.ReadWorkspace(s.Workspace)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if _, err1 := w.Packages.Get(depId); err1 != nil {
    var err2 error
		if depId, err2 = w.Packages.CreteRecord(depId); err2 != nil {
			fmt.Println(err1)
			fmt.Println(err2)
			os.Exit(1)
		}
	  if err2 = w.Write(); err2 != nil {
			fmt.Println(err1)
			fmt.Println(err2)
			os.Exit(1)
	  }
    if err2 = w.Packages.UpdateRepo(depId); err2 != nil {
			fmt.Println(err1)
			fmt.Println(err2)
			os.Exit(1)
    }
	}

	scaff, err := scaffolding.NewScaffolding(w, s.Workspace)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	m, err := scaff.AddModule(depId, dest)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := scaff.BuildSubModule(m); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

  if err := scaff.Write(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

	fmt.Println("Module added")
}
