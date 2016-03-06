package settings

import (
	"github.com/goatcms/goat-core/varutil"
	"os"
	"os/user"
)

const (
	settingsPath = "/.goat/settings.json"
)

type Settings struct {
	path      string
	Workspace string `json:"workspace"`
}

func NewSettings() (*Settings, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}
	return &Settings{
		path:      usr.HomeDir + settingsPath,
		Workspace: os.Getenv("GOPATH"),
	}, nil
}

func (s *Settings) Read() error {
	return varutil.ReadJson(s.path, s)
}

func (s *Settings) Write() error {
	return varutil.WriteJson(s.path, s)
}
