package command

import (
	"fmt"
	"github.com/goatcms/goat-core/workspace"
	"github.com/goatcms/goat-core/scaffolding"
	"github.com/goatcms/goat/src/settings"
	"os"
)


func AddResourceCommand(s *settings.Settings, args []string) {
	if len(args) < 1 {
		fmt.Println("Resource type is required")
		os.Exit(1)
	}
	if len(args) < 2 {
		fmt.Println("Resource name is required")
		os.Exit(1)
	}

	w, err := workspace.ReadWorkspace(s.Workspace)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	r := scaffolding.Resource{
		Type: args[0],
		Name: args[1],
		Data: map[string]string{},
	}

	scaff, err := scaffolding.NewScaffolding(w, s.Workspace)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := scaff.BuildResource(r); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Resource " + r.Name + " builder")
}
