package config

import (
	"encoding/xml"
	"os"
)

func Read(filename string) (conf *Config, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = xml.NewDecoder(f).Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

type Config struct {
	Mysql struct {
		Db       string `xml:"db"`
		User     string `xml:"user"`
		Password string `xml:"password"`
		Port     string `xml:"port"`
	} `xml:"mysql"`
	Path struct {
		Image     string `xml:"image"`
		Thumbnail string `xml:"thumbnail"`
	} `xml:"path"`
	Port string `xml:"port"`
}
