package container

func (c *Container) Init() error {
	return c.init()
}

func (c *Container) Start() {
	c.start()
}

func (c *Container) Stop() {
	c.stop()
}

func (c *Container) Restart() {
	c.restart()
}

func (c *Container) Remove(force bool) error {
	return c.remove(force)
}

func (c *Container) Logs() []byte {
	return c.logs()
}
