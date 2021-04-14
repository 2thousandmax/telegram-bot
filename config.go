// package data
package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Data     Data     `yaml:"data"`
	Messages Messages `yaml:"messages"`
}

type Data struct {
	Rings     []string                                 `yaml:"rings"`
	Lecturers []string                                 `yaml:"lecturers"`
	Classes   []string                                 `yaml:"classes"`
	Controls  []string                                 `yaml:"controls"`
	Timetable map[string]map[string][][][4]interface{} `yaml:"timetable"`
}

type Messages struct {
	Responses Responses `yaml:"responses"`
	Errors    Errors    `yaml:"errors"`
}

type Responses struct {
	Start          string `yaml:"start"`
	Setup          string `yaml:"setup"`
	UnknownCommand string `yaml:"unknown_command"`
}

type Errors struct {
	Default         string `yaml:"default"`
	InvalidGroup    string `yaml:"invalid_group"`
	GroupNotFound   string `yaml:"group_not_found"`
	MessageOutdated string `yaml:"message_outdated"`
}

func NewConfig(file string) (*Config, error) {
	cfg := Config{}
	f, err := ioutil.ReadFile(file)
	if err != nil {
		return &Config{}, err
	}

	if err := yaml.Unmarshal(f, &cfg); err != nil {
		return &Config{}, err
	}

	fmt.Println("Successfully unmarshalled")
	fmt.Println(cfg.Data.Rings)
	d, err := yaml.Marshal(&cfg)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	fmt.Printf("--- dump:\n%s\n\n", string(d))

	return &cfg, nil
}
