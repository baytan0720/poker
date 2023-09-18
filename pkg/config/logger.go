package config

type Logger struct {
	Level  string `yaml:"level"`
	Format string `yaml:"formatter"`
	Output string `yaml:"output"`
}

func GetLoggerLevel() string {
	return Cfg.Logger.Level
}

func GetLoggerFormat() string {
	return Cfg.Logger.Format
}

func GetLoggerOutput() string {
	return Cfg.Logger.Output
}
