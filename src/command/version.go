package command

import (
	"fmt"
	"github.com/goatcms/goat/src/settings"
	"github.com/goatcms/goat/src/consts"
)

func VersionCommand(s *settings.Settings, args []string) {
	fmt.Println("%s", consts.Version)
}
