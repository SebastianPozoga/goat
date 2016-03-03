package command

import (
	"fmt"
	"github.com/goatcms/goat/src/settings"
)

func HelpCommand(s *settings.Settings, args []string) {
	help := fmt.Sprintf("Goat version: %s", s.Version)
	help += "\n  goat version -> display version"
	help += "\n  goat help -> display help"
	fmt.Println(help)
}
