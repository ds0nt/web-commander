package main

import (
	log "github.com/Sirupsen/logrus"

	"github.com/BurntSushi/toml"
	"github.com/k0kubun/pp"
)

type Config struct {
	Consumer struct {
		Key    string `toml:"key"`
		Secret string `toml:"secret"`
	} `toml:"twitter_consumer_key"`
	Access struct {
		Token  string `toml:"token"`
		Secret string `toml:"secret"`
	} `toml:"twitter_access_token"`
	Redis struct {
		Host     string `toml:"host"`
		Password string `toml:"password"`
	} `toml:"redis"`
}

var config Config

func loadConfig() {
	_, err := toml.DecodeFile(*conf, &config)
	if err != nil {
		log.Fatal(err)
	}
	pp.Print(config)
}
