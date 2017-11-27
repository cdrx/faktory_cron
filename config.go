package main

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Path    string
	Faktory string `yaml:"faktory"`
	Jobs    []*Job `yaml:"jobs"`
}

func (c *Config) Update() error {
	b, err := ioutil.ReadFile(c.Path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(b, &c)
	if err != nil {
		return err
	}

	if config.Faktory != "" {
		os.Setenv("FAKTORY_URL", config.Faktory)
	}

	return nil
}

func NewConfig(path string) *Config {
	return &Config{
		Path: path,
	}
}
