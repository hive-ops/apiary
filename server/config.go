package server

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Config struct {
	ClusterName string   `yaml:"cluster_name"`
	IP          string   `yaml:"ip"`
	Port        int      `yaml:"port"`
	SeedNodes   []string `yaml:"seed_nodes"`
	Datacenter  string   `yaml:"datacenter"`
	Rack        string   `yaml:"rack"`
}

func LoadConfig(path string) *Config {
	var config Config

	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("error reading YAML file: %v", err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error parsing YAML file: %v", err)
	}

	return &config
}
