package command

import (
	"encoding/json"
	"fmt"
	"github.com/goatcms/goat-core/scaffolding"
	"github.com/goatcms/goat-core/workspace"
	"github.com/goatcms/goat/src/interactive"
	"github.com/goatcms/goat/src/settings"
	"os"
	"strings"
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
	if len(args) < 3 {
		fmt.Println("Resource data (json) is required")
		os.Exit(1)
	}

	var data interface{}
	jsonDecoder := json.NewDecoder(strings.NewReader(args[2]))
	if err := jsonDecoder.Decode(data); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	addResource(s, args[0], args[1], data)
	fmt.Println(args[0] + " resource " + args[1] + " builder")
}

func AddModelResourceCommand(s *settings.Settings, args []string) {
	if len(args) < 1 {
		fmt.Println("Resource name is required")
		os.Exit(1)
	}

	modelData := interactive.Model{}
	if err := modelData.Scan(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	addResource(s, "model", args[0], modelData)
	fmt.Println("Model " + args[0] + " builded")
}

func addResource(s *settings.Settings, resourceType, name string, data interface{}) {
	w, err := workspace.ReadWorkspace(s.Workspace)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	r := scaffolding.Resource{
		Type: resourceType,
		Name: name,
		Data: data,
	}

	scaff, err := scaffolding.NewScaffolding(w, s.Workspace)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := scaff.AddResource(r); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
