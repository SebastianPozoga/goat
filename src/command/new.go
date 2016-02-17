package command

import (
	"fmt"
	"path"
	"log"
	"os"
	"github.com/goatcms/goat-core/filesystem"
	"github.com/goatcms/goat-core/system"
)

func NewCommand(args []string) {
	if len(args) < 1 {
		log.Printf("project type or url is required\n")
		return
	}

	if len(args) < 2 {
		log.Printf("project name is required\n")
		return
	}

	projectPath := path.Join(os.Getenv("GOPATH"), "src", args[1])

	if filesystem.FileExists(projectPath) {
		log.Printf("A project already exists (at path %s)\n", projectPath)
		return
	}

	url := args[0]

	switch url {
	case "app":
		url = "github.com/goatcms/goat-app"
	}

	result, err := system.RunCommand("go", "get", url)
	if err != nil {
		log.Printf("Can not clone repositiory %s", err)
		return
	}

	log.Printf("%s", string(result))
}
