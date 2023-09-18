package main

import (
	"flag"

	"poker/daemon/engine"
	"poker/pkg/config"

	log "github.com/sirupsen/logrus"
)

var (
	configPath   string
	pokerVersion string
	buildTime    string
	gitRevision  string
	goVersion    string
)

func init() {
	flag.StringVar(&configPath, "config", "/etc/poker/config.yaml", "config file path")
}

func main() {
	flag.Parse()

	if err := config.ReadConfig(configPath); err != nil {
		log.Fatalf("read config fail, error: %s", err)
	}

	config.SetVersion(pokerVersion, buildTime, gitRevision, goVersion)

	engine.Run()
}
