package container

import (
	"encoding/json"
	"net"
	"os"
	"path"
	"poker/daemon/internal/image"
	"poker/pkg/errno"
	"time"

	"poker/daemon/internal/pkg/common"
	"poker/pkg/config"
)

type Container struct {
	Id             string    `json:"id"`
	Name           string    `json:"name"`
	CreatedAt      time.Time `json:"created_at"`
	Path           string    `json:"path"`
	Args           []string  `json:"args"`
	State          State     `json:"state"`
	Image          string    `json:"image"`
	ResolvConfPath string    `json:"resolv_conf_path"`
	HostnamePath   string    `json:"hostname_path"`
	HostsPath      string    `json:"hosts_path"`
	LogPath        string    `json:"log_path"`
	LogFile        *os.File  `json:"-"`
	AutoRestart    bool      `json:"restart"`
	AutoRemove     bool      `json:"auto_remove"`
	RestartCount   string    `json:"restart_count"`
	Driver         string    `json:"driver"`
	Mounts         Mounts    `json:"mounts"`
	Config         Config    `json:"config"`
	Network        Network   `json:"network"`
}

type ContainerState string

const (
	Init       ContainerState = "init"
	Created    ContainerState = "created"
	Running    ContainerState = "running"
	Stopped    ContainerState = "stopped"
	Exited     ContainerState = "exited"
	Restarting ContainerState = "restarting"
	Dead       ContainerState = "dead"
)

type State struct {
	Status     ContainerState `json:"status"`
	Pid        int            `json:"pid"`
	ExitCode   string         `json:"exit_code"`
	Error      string         `json:"error"`
	StartedAt  time.Time      `json:"started_at"`
	FinishedAt time.Time      `json:"finished_at"`
}

type Mounts struct {
	Type        string `json:"type"`
	Name        string `json:"name"`
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Driver      string `json:"driver"`
	Mode        string `json:"mode"`
	RW          string `json:"rw"`
	Propagation string `json:"propagation"`
}

type Config struct {
	Hostname     string              `json:"hostname"`
	Domainname   string              `json:"domainname"`
	User         string              `json:"user"`
	ExposedPorts map[string]struct{} `json:"exposed_ports"`
	Tty          bool                `json:"tty"`
	TtySocket    string              `json:"-"`
	TtyListener  net.Listener        `json:"-"`
	Env          []string            `json:"env"`
	Volumes      map[string]struct{} `json:"volumes"`
	WorkingDir   string              `json:"working_dir"`
	Entrypoint   []string            `json:"entrypoint"`
}

type Network struct {
	// TODO
}

func NewContainer(image *image.Image, options ...CreateOption) *Container {
	c := &Container{
		Id:        common.GenerateRandomId(),
		Name:      common.GenerateRandomName(),
		CreatedAt: time.Now(),
		Path:      image.Id,
		Config:    parseImageConfig2ContainerConfig(image.ContainerConfig),
		State: State{
			Status: Init,
		},
	}

	for _, option := range options {
		option(c)
	}

	return c
}

func GetFromMetadata(src string) (*Container, error) {
	b, err := os.ReadFile(src)
	if err != nil {
		return nil, err
	}

	c := Container{}
	if err := json.Unmarshal(b, &c); err != nil {
		return nil, err
	}

	return &c, nil
}

func (c *Container) dir() string {
	return path.Join(config.GetContainerBasePath(), c.Id)
}

func (c *Container) metadataPath() string {
	return path.Join(c.dir(), "metadata.json")
}

func GetMetadataPath(id string) string {
	return path.Join(config.GetContainerBasePath(), id, "metadata.json")
}

func (c *Container) setLogPath() error {
	c.LogPath = path.Join(c.dir(), "output.log")
	f, err := os.OpenFile(c.LogPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	c.LogFile = f
	return nil
}

func (c *Container) save() error {
	if c.State.Status == Init {
		c.State.Status = Created
	}

	b, err := json.Marshal(c)
	if err != nil {
		return err
	}

	if err := os.WriteFile(c.metadataPath(), b, 0777); err != nil {
		return errno.SimpleErr(errno.ErrWriteFileFailed)
	}

	return nil
}

func (c *Container) init() error {
	return nil
}

func (c *Container) start() {

}

func (c *Container) stop() {

}

func (c *Container) restart() {
	c.stop()
	c.start()
}

func (c *Container) remove(force bool) error {
	if !c.removable() {
		if force {
			c.stop()
		} else {
			return errno.SimpleErr(errno.ErrRemoveRunningContainer)
		}
	}

	if err := os.RemoveAll(c.dir()); err != nil {
		return errno.SimpleErr(errno.ErrRemoveContainerDirFailed)
	}

	return nil
}

func (c *Container) removable() bool {
	return c.State.Status == Exited || c.State.Status == Stopped || c.State.Status == Dead
}

func (c *Container) setTty() error {
	socketPath := path.Join(c.dir(), "tty.sock")
	c.Config.Tty = true
	c.Config.TtySocket = socketPath
	l, err := common.ListenOnUnix(socketPath)
	if err != nil {
		return err
	}
	c.Config.TtyListener = l
	return nil
}

func (c *Container) logs() []byte {
	b, _ := os.ReadFile(c.LogPath)
	return b
}

func (c *Container) Rename(newName string) error {
	c.Name = newName
	return c.save()
}
