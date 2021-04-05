package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"gopkg.in/yaml.v2"
)

type Response struct {
	TimeStamp time.Time                                `yaml:"timestamp"`
	Lecturers []string                                 `yaml:"lecturers"`
	Classes   []string                                 `yaml:"classes"`
	Timetable map[string]map[string][][][4]interface{} `yaml:"timetable"`
}

func main() {
	r := Response{}
 
	f, err := ioutil.ReadFile("data.yaml")
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(f, &r); err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("%+v\n", r)
}