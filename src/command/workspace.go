package command

import (
	"fmt"
	"github.com/goatcms/goat-core/filesystem"
	"github.com/goatcms/goat-core/workspace"
	"github.com/goatcms/goat/src/settings"
	"os"
	"path/filepath"
)

func SetWorkspaceCommand(s *settings.Settings, args []string) {
	var (
		workspacePath string
		err           error
	)
	if len(args) < 1 {
		workspacePath, err = filepath.Abs(".")
		if err != nil {
			fmt.Println("get current path fail", err)
			os.Exit(1)
		}
	} else {
		workspacePath = args[0]
	}

	s.Workspace = workspacePath
	if err = s.Write(); err!=nil {
		fmt.Println("no write goat setting file", err)
		os.Exit(1)
	}

	fmt.Println(s.Workspace)
}

func GetWorkspaceCommand(s *settings.Settings, args []string) {
	fmt.Println(s.Workspace)
}

func NewWorkspaceCommand(s *settings.Settings, args []string) {
	if len(args) < 1 {
		fmt.Println("Workspace path is required")
		return
	}

	if filesystem.IsExist(args[0]) {
		fmt.Println("File or directory", args[0], "exists")
		os.Exit(1)
	}

	w := workspace.NewWorkspace(args[0])
	if err := w.Create(); err != nil {
		fmt.Println("Create new workspace fail", err)
		os.Exit(1)
	}
}
