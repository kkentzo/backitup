package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Actions  []*Action
	Notifier *Notifier `yaml:"notify"`
}

func NewConfigFromFile(file string) (*Config, error) {
	contents, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	c := &Config{}
	err = yaml.Unmarshal(contents, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
