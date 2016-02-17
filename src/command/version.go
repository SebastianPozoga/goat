package command

import (
	"fmt"
)

func VersionCommand(args []string, version string) {
	msg := fmt.Sprintf("%s", version)
	fmt.Println(msg)
}
