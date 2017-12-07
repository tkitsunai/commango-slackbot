package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

var Config ApplicationConfig

type ApplicationConfig struct {
	Slack Slack
}

type Slack struct {
	Channel string
	Token   string
	User    string
	Icon    string
	BotID   string
}

func (s *Slack) BotMention() string {
	return "<@" + s.BotID + ">"
}

func init() {
	if _, err := toml.DecodeFile("config.toml", &Config); err != nil {
		panic(fmt.Errorf("could not read config ", err))
	}
}
