package config

type Container struct {
	BasePath string `yaml:"base_path"`
}

func GetContainerBasePath() string {
	return Cfg.Container.BasePath
}
