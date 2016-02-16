package main

import (
	"flag"
	"log"
	"os"

	"errors"
	"fmt"
	"strings"
	"time"
)

var (
	// contain configuration
	config map[string]string
)

const (
	// The version of this tool
	version = "0.0.1"

	// Used for outputting console messages
	logDivider = "\n------\n"
)

func main() {
	if len(args) < 2 {
		showHelp()
		return
	}

	command := os.Args[1]
	switch command {

	case "version", "v":
		showVersion()

	case "help", "h", "wat", "?":
		showHelp(args)

	default:
		consoleError("Unknow command")
	}

}

func showVersion() {
	msg := logDivider
	msg += fmt.Sprintf("Goat version: %s", version)
	msg += logDivider
	log.Print(msg)
}


func consoleError(msg string) {
	log.Fatal(msg)
	os.Exit(1)
}

func showHelp() {
	help := logDivider
	help += fmt.Sprintf("Goat version: %s", version)
	help += "\n  goat version -> display version"
	help += "\n  goat help -> display help"
	help += logDivider
	log.Print(help)
}
