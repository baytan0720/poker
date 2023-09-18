package image

import (
	"time"
)

type Image struct {
	Id              string    `json:"id"`
	RepoTags        []string  `json:"repo_tags"`
	RepoDigests     []string  `json:"repo_digests"`
	Parent          string    `json:"parent"`
	CreatedAt       time.Time `json:"created_at"`
	ContainerConfig Config    `json:"container_config"`
	Author          string    `json:"author"`
	Size            int       `json:"size"`
	RootFS          RootFS    `json:"root_fs"`
}

type Config struct {
	Hostname     string              `json:"hostname"`
	Domainname   string              `json:"domainname"`
	User         string              `json:"user"`
	ExposedPorts map[string]struct{} `json:"exposed_ports"`
	Tty          bool                `json:"tty"`
	Env          []string            `json:"env"`
	Volumes      map[string]struct{} `json:"volumes"`
	WorkingDir   string              `json:"working_dir"`
	Entrypoint   []string            `json:"entrypoint"`
}

type RootFS struct {
	Type   string   `json:"type"`
	Layers []string `json:"layers"`
}
