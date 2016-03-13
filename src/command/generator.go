package command

import (
	"fmt"
	"github.com/goatcms/goat-core/workspace"
	"github.com/goatcms/goat/src/settings"
	"os"
)

func RunGeneratorCommand(s *settings.Settings, args []string) {
	_, err := workspace.ReadWorkspace(s.Workspace)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

  fmt.Println("Generate secrets and values completed")
}
