package command

import (
	"fmt"
	"github.com/goatcms/goat/src/settings"
	"github.com/goatcms/goat/src/consts"
)

func HelpCommand(s *settings.Settings, args []string) {
	help := fmt.Sprintf("Goat version: %s", consts.Version)
	help += "\n  goat version -> display version"
	help += "\n  goat help -> display help"
	fmt.Println(help)
}
