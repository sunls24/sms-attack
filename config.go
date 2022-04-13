package main

import (
	"github.com/pelletier/go-toml/v2"
	"github.com/sunls24/x"
	"io/ioutil"
)

type Config struct {
	Base    BaseCfg
	ApiList []ApiCfg `toml:"api"`
}

type BaseCfg struct {
	DefaultInterval int
}

type ApiCfg struct {
	Address  string
	Interval int
}

func ParseConfig() *Config {
	var file, err = ioutil.ReadFile("./config.toml")
	x.CheckPanic(err)
	var cfg Config
	err = toml.Unmarshal(file, &cfg)
	x.CheckPanic(err)
	return &cfg
}
