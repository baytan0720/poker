package container

import "poker/daemon/internal/image"

func parseImageConfig2ContainerConfig(imageConfig image.Config) Config {
	return Config{
		Hostname:     imageConfig.Hostname,
		Domainname:   imageConfig.Domainname,
		User:         imageConfig.User,
		ExposedPorts: imageConfig.ExposedPorts,
		Tty:          imageConfig.Tty,
		Env:          imageConfig.Env,
		Volumes:      imageConfig.Volumes,
		WorkingDir:   imageConfig.WorkingDir,
		Entrypoint:   imageConfig.Entrypoint,
	}
}
