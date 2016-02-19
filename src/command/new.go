package command

import (
	"path"
	"log"
	"os"
	"github.com/goatcms/goat-core/filesystem"
	"github.com/goatcms/goat-core/system"
	"github.com/goatcms/goat-core/scaffolding"
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

	if !filesystem.IsDir(path.Join(os.Getenv("GOPATH"), "src", url)) {
		_, err := system.RunCommand("go", "load", url)
		if err != nil {
			log.Printf("Can not clone repositiory %s", err)
			os.Exit(1)
		}
	}

	err := scaffolding.Build(".", "./", []string{"el1", "el2"})
	if err != nil {
		log.Printf("Can not build a project", err)
		os.Exit(1)
	}

	log.Printf("some end log")
}
