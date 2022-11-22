package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

type CfgStruct struct {
	Port1 string `yaml:"port1"`
	Port2 string `yaml:"port2"`
}

func NewConfigurationFromFile(configPath *string) (CfgStruct, error) {
	var config CfgStruct
	file, err := os.Open(*configPath)
	if err != nil {
		return config, fmt.Errorf("Failed open config file :%v", err)
	}
	defer file.Close()
	read, err := io.ReadAll(file)
	if err != nil {
		return config, fmt.Errorf("Failed reading file :%v", err)
	}
	err = yaml.Unmarshal(read, &config)
	if err != nil {
		return config, fmt.Errorf("Yaml unmarshalling error :%v", err)
	}
	return config, nil
}
