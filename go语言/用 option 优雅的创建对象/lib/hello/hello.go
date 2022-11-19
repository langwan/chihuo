package hello

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

type Instance struct {
	Message    string
	Name       string
	LatestTime time.Time
	Version    string
}

type Option func(*Instance)

func New(opts ...Option) *Instance {
	i := &Instance{}
	for _, opt := range opts {
		opt(i)
	}
	return i
}

type Config struct {
	Message string `json:"message"`
}

func WithConfig(filename string) Option {
	data, err := os.ReadFile(filename)
	if err != nil {
		return func(i *Instance) {

		}
	}
	c := &Config{}
	yaml.Unmarshal(data, c)
	return func(i *Instance) {
		i.Message = c.Message
	}
}

func WithName(name string) Option {

	return func(i *Instance) {
		i.Name = name
	}
}

func WithVersion(version string) Option {
	return func(i *Instance) {
		i.LatestTime = time.Now()
		i.Version = version
	}
}

func (i *Instance) String() string {
	return fmt.Sprintf("name is %s, message = %s, latest time = %s, version = %s", i.Name, i.Message, i.LatestTime, i.Version)
}
