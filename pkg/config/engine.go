package config

type Engine struct {
	PokerVersion string `yaml:"-"`
	BuildTime    string `yaml:"-"`
	GitRevision  string `yaml:"-"`
	GoVersion    string `yaml:"-"`
	Socket       string `yaml:"socket"`
}

func GetSocketFile() string {
	return Cfg.Engine.Socket
}

func GetPokerVersion() string {
	return Cfg.Engine.PokerVersion
}

func GetBuildTime() string {
	return Cfg.Engine.BuildTime
}

func GetGitRevision() string {
	return Cfg.Engine.GitRevision
}

func GetGoVersion() string {
	return Cfg.Engine.GoVersion
}
