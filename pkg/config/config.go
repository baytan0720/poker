package config

import (
	"os"

	"gopkg.in/yaml.v3"

	"poker/pkg/errno"
)

var Cfg Config

type Config struct {
	Engine    Engine    `yaml:"Engine"`
	Logger    Logger    `yaml:"Logger"`
	Container Container `yaml:"Container"`
}

func ReadConfig(configPath string) error {
	b, err := os.ReadFile(configPath)
	if err != nil {
		return errno.NewErr(errno.ErrReadFileFailed, configPath)
	}

	if err := yaml.Unmarshal(b, &Cfg); err != nil {
		return err
	}

	return nil
}

func SetVersion(pokerVersion string, buildTime string, gitRevision string, goVersion string) {
	Cfg.Engine.PokerVersion = pokerVersion
	Cfg.Engine.BuildTime = buildTime
	Cfg.Engine.GitRevision = gitRevision
	Cfg.Engine.GoVersion = goVersion
}
