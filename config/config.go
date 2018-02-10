package config

import (
	"encoding/xml"
	"os"
)

func Read(filename string) (db, user, password string, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", "", "", err
	}

	var cfg config
	err = xml.NewDecoder(f).Decode(&cfg)
	if err != nil {
		return "", "", "", err
	}

	return cfg.Db, cfg.Name, cfg.Password, nil
}

type config struct {
	Db       string `xml:"db"`
	Name     string `xml:"user"`
	Password string `xml:"password"`
}
