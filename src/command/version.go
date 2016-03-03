package command

import (
	"fmt"
	"github.com/goatcms/goat/src/settings"
)

func VersionCommand(s *settings.Settings, args []string) {
	fmt.Println("%s", s.Version)
}
