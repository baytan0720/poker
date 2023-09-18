package manager

import (
	"os"
	"poker/daemon/internal/container"
	"poker/daemon/internal/image"
	"poker/daemon/internal/pkg/common"
	"poker/pkg/config"
	"poker/pkg/errno"
	"sort"
)

var m *manager

type manager struct {
	containers     map[string]*container.Container
	name2Container map[string]*container.Container
	images         map[string]*image.Image
}

func newManager() *manager {
	return &manager{
		containers:     make(map[string]*container.Container),
		name2Container: make(map[string]*container.Container),
		images:         make(map[string]*image.Image),
	}
}

func (m *manager) addContainer(c *container.Container) {
	m.containers[c.Id] = c
	m.name2Container[c.Name] = c
}

func (m *manager) getContainer(idOrName string) (*container.Container, error) {
	if len(idOrName) == common.ContainerIDLength {
		if c, ok := m.containers[idOrName]; ok {
			return c, nil
		}
	}

	if c, ok := m.name2Container[idOrName]; ok {
		return c, nil
	}

	return nil, errno.SimpleErr(errno.ErrContainerNotFound)
}

func (m *manager) removeContainer(c *container.Container) {
	delete(m.containers, c.Id)
	delete(m.name2Container, c.Name)
}

func (m *manager) addImage(i *image.Image) {
	m.images[i.Id] = i
}

func (m *manager) getImage(id string) (*image.Image, error) {
	if id == "" {
		return nil, errno.SimpleErr(errno.ErrImageEmpty)
	}

	if i, ok := m.images[id]; ok {
		return i, nil
	}

	return nil, errno.SimpleErr(errno.ErrImageNotFound)
}

func (m *manager) removeImage(i *image.Image) {
	delete(m.images, i.Id)
}

func (m *manager) listContainers() []*container.Container {
	containers := make(container.ContainerSlice, 0, len(m.containers))
	for _, v := range m.containers {
		containers = append(containers, v)
	}
	sort.Sort(containers)
	return containers
}

func (m *manager) readAllContainers() error {
	basePath := config.GetContainerBasePath()
	entry, err := os.ReadDir(basePath)
	if err != nil {
		return err
	}
	for _, v := range entry {
		c, err := container.GetFromMetadata(container.GetMetadataPath(v.Name()))
		if err != nil {
			continue
		}
		if c.AutoRestart {
			c.Start()
		}
		m.addContainer(c)
	}
	return nil
}

func (m *manager) readAllImages() error {
	// TODO
	return nil
}

func (m *manager) validateName(name string) error {
	if name == "" {
		return nil
	}

	if len(name) >= common.ContainerIDLength {
		return errno.SimpleErr(errno.ErrContainerNameTooLong)
	}

	if _, ok := m.name2Container[name]; ok {
		return errno.SimpleErr(errno.ErrContainerNameConflict)
	}

	return nil
}

func InitManager() error {
	m = newManager()

	if err := m.readAllContainers(); err != nil {
		return err
	}

	if err := m.readAllImages(); err != nil {
		return err
	}

	return nil
}

func AddContainer(c *container.Container) {
	m.addContainer(c)
}

func GetContainer(idOrName string) (*container.Container, error) {
	return m.getContainer(idOrName)
}

func RemoveContainer(c *container.Container) {
	m.removeContainer(c)
}

func AddImage(i *image.Image) {
	m.addImage(i)
}

func GetImage(id string) (*image.Image, error) {
	return m.getImage(id)
}

func RemoveImage(i *image.Image) {
	m.removeImage(i)
}

func ValidateName(name string) error {
	return m.validateName(name)
}

func ListContainers() []*container.Container {
	return m.listContainers()
}
