package config

import (
	"encoding/xml"
	"os"

	"github.com/go-sql-driver/mysql"
)

func Read(filename string) (conf *mysql.Config, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	var cfg config
	err = xml.NewDecoder(f).Decode(&cfg)
	if err != nil {
		return nil, err
	}

	conf = mysql.NewConfig()
	conf.User = cfg.User
	conf.Passwd = cfg.Password
	conf.DBName = cfg.Db
	conf.Net = "tcp"
	conf.Addr = "localhost:" + cfg.Port
	conf.ParseTime = true

	return conf, nil
}

type config struct {
	Db       string `xml:"db"`
	User     string `xml:"user"`
	Password string `xml:"password"`
	Port     string `xml:"port"`
}
