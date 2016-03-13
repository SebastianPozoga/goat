package command

import (
	"fmt"
	"github.com/goatcms/goat-core/filesystem"
	"github.com/goatcms/goat-core/scaffolding"
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
	if err = s.Write(); err != nil {
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

	w, err := workspace.NewWorkspace(args[0])
	if err != nil {
		fmt.Println("Create new workspace object fail", err)
		os.Exit(1)
	}

	if len(args) < 2 {
		// if a user didn't set source repository, build from cache
		if err := w.CreateFromCache(); err != nil {
			fmt.Println("Create new workspace (from default/cached repository) fail", err)
			os.Exit(1)
		}
	} else {
		// if a user set source repository, build from remote repository
		if err := w.CreateFromRepo(args[1]); err != nil {
			fmt.Println("Create new workspace (from remote repository) fail", err)
			os.Exit(1)
		}
	}

	//build submodules
	if scaffolding.IsScaffoldingDir(args[0]) {
		s, err := scaffolding.ReadScaffolding(w, args[0])
		if err != nil {
			fmt.Println("Read scaffolfing fail", err)
			os.Exit(1)
		}
		if err := s.BuildAllSubModules(); err != nil {
			fmt.Println("Build submodules fail", err)
			os.Exit(1)
		}
	}

	//set new workspace
	s.Workspace = w.GetAbsPath()
	if err = s.Write(); err != nil {
		fmt.Println("no write goat setting file", err)
		os.Exit(1)
	}
}

func CleanWorkspaceCacheCommand(s *settings.Settings, args []string) {
	if err := workspace.CleanWorkspaceCache(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
