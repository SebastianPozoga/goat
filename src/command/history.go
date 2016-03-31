package command

import (
	"fmt"
	"github.com/goatcms/goat-core/scaffolding"
	"github.com/goatcms/goat-core/workspace"
	"github.com/goatcms/goat/src/settings"
  "os"
)

func BuildHistoryCommand(s *settings.Settings, args []string) {
	w, err := workspace.ReadWorkspace(s.Workspace)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
  scaff, err := scaffolding.ReadScaffolding(w, s.Workspace)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := scaff.BuildHistory(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("All rebuilded")
}
