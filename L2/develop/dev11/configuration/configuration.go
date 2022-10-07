package configuration

import (
	"gopkg.in/yaml.v2"
	"io"
	"log"
	"os"
)

type Config struct {
	Addr string `yaml:"port"`
}

func New() *Config {
	return &Config{}
}

func (c *Config) LoadConfig(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		log.Println("open file")
		return err
	}
	defer file.Close()
	all, err := io.ReadAll(file)
	if err != nil {
		log.Println("Read all")
		return err
	}
	err = yaml.Unmarshal(all, c)
	if err != nil {
		log.Println("Unmarshall data")
		log.Println(err)
		return err
	}
	return nil
}
