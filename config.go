// package data
package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Postgres PostgresDatabaseConfig `yaml:"postgres"`
	Bot      BotConfig              `yaml:"bot"`
	Replies  map[string]string      `yaml:"replies"`
	Errors   map[string]string      `yaml:"errors"`
}

type BotConfig struct {
	Host string `yaml:"host"`
}


// type Data struct {
// 	Rings     []string                                 `yaml:"rings"`
// 	Lecturers []string                                 `yaml:"lecturers"`
// 	Classes   []string                                 `yaml:"classes"`
// 	Controls  []string                                 `yaml:"controls"`
// 	Timetable map[string]map[string][][][4]interface{} `yaml:"timetable"`
// }

// type Messages struct {
// 	Responses Responses `yaml:"responses"`
// 	Errors    Errors    `yaml:"errors"`
// }

// type Responses struct {
// 	Start          string `yaml:"start"`
// 	Setup          string `yaml:"setup"`
// 	UnknownCommand string `yaml:"unknown_command"`
// }

func NewConfig(file string) (Config, error) {
	cfg := Config{}
	f, err := ioutil.ReadFile(file)
	if err != nil {
		return Config{}, err
	}

	if err := yaml.Unmarshal(f, &cfg); err != nil {
		return Config{}, err
	}

	fmt.Println("Successfully unmarshalled")
	d, err := yaml.Marshal(&cfg)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	fmt.Printf("--- dump:\n%s\n\n", string(d))

	return cfg, nil
}
