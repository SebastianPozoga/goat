package command

import (
	"fmt"
	"github.com/goatcms/goat/src/consts"
	"github.com/goatcms/goat/src/settings"
)

func HelpCommand(s *settings.Settings, args []string) {
	help := "\n Goat v" + consts.Version
	help += "\n  workspace:new [path, [gitrepo]] -> create new workspace"
	help += "\n  workspace:set [[path]] -> set path as current workspace"
	help += "\n                            (default current path)"
	help += "\n  workspace:get          -> display workspace"
	help += "\n  workspace:cache:clean  -> clean workpsace cache (You can use it"
	help += "\n                            to remove cached default repository)"
	help += "\n  generator:run -> generate default values (run by defult "
	help += "\n                   hen create or use workspace)"
	help += "\n  packages:add [depId]   -> add dependencie to workspace. Example:"
	help += "\n                            'github.com/goatcms/go-project' or"
	help += "\n                            'github.com/goatcms/go-project.version'"
	help += "\n  packages:update        -> download & update all packages"
	help += "\n  resource:add [type, name, jsonData] -> add and build resource"
	help += "\n  resource:model:add [name]           -> add model resource (*i)"
	help += "\n  history:build          -> rebuild hole history"
	help += "\n  module:add [destPath, depId] -> add and build module"
	help += "\n  version -> display version"
	help += "\n  help    -> display help"
	help += "\n"
	help += "\n  *i - interactive mode"
	fmt.Println(help)
}
