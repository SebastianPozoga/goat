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
	Version   string `json:"-"`
	Workspace string `json:"workspace"`
}

func NewSettings(version string) (*Settings, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}
	return &Settings{
		path:      usr.HomeDir + settingsPath,
		Workspace: os.Getenv("GOPATH"),
		Version:   version,
	}, nil
}

func (s *Settings) Read() error {
	return varutil.ReadJson(s.path, s)
}

func (s *Settings) Write() error {
	return varutil.WriteJson(s.path, s)
}
