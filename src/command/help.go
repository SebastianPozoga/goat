package command

import (
	"fmt"
)

func HelpCommand(args []string, version string) {
	help := fmt.Sprintf("Goat version: %s", version)
	help += "\n  goat version -> display version"
	help += "\n  goat help -> display help"
	fmt.Println(help)
}
