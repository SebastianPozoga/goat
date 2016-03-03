package command
/*
import (
	"fmt"
	"github.com/goatcms/goat-core/filesystem"
	"github.com/goatcms/goat-core/repos"
	"github.com/goatcms/goat-core/scaffolding"
	"github.com/goatcms/goat/src/settings"
	"log"
	"os"
	"path"
)

func NewCommand(s *settings.Settings, args []string) {
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
	case "go-project":
		url = "github.com/goatcms/go-project"
	}

	destPath, err := repos.GetRepoPath(args[1])
	if filesystem.FileExists(destPath) {
		log.Printf("Directory %s exists", args[1])
		os.Exit(1)
	}

	scaff, err := scaffolding.NewScaffolding(url)
	if err != nil {
		log.Printf("Scaffolding error ", args[1])
		fmt.Println(err)
		os.Exit(1)
	}
	if err = scaff.BuildSource(destPath); err != nil {
		log.Printf("Can not build a project", err)
		os.Exit(1)
	}

	log.Printf("Created success")
}
*/
