package container

type CreateOption func(*Container)

func WithName(name string) CreateOption {
	return func(c *Container) {
		c.Name = name
	}
}

func WithCommandArgs(command string, args []string) CreateOption {
	return func(c *Container) {
		c.Path = command
		c.Args = args
	}
}

func WithEnv(env []string) CreateOption {
	return func(c *Container) {
		c.Config.Env = append(c.Config.Env, env...)
	}
}

func WithExposedPorts(ports []string) CreateOption {
	return func(c *Container) {
		for _, p := range ports {
			c.Config.ExposedPorts[p] = struct{}{}
		}
	}
}

func WithVolumes(volumes []string) CreateOption {
	return func(c *Container) {
		for _, v := range volumes {
			c.Config.Volumes[v] = struct{}{}
		}
	}
}

func WithHostname(hostname string) CreateOption {
	return func(c *Container) {
		c.Config.Hostname = hostname
	}
}

func WithUser(user string) CreateOption {
	return func(c *Container) {
		c.Config.User = user
	}
}

func WithWorkingDir(workingDir string) CreateOption {
	return func(c *Container) {
		c.Config.WorkingDir = workingDir
	}
}

func WithAutoRestart() CreateOption {
	return func(c *Container) {
		c.AutoRestart = true
	}
}

func WithAutoRemove() CreateOption {
	return func(c *Container) {
		c.AutoRemove = true
	}
}

func WithTty() CreateOption {
	return func(c *Container) {
		c.Config.Tty = true
	}
}
