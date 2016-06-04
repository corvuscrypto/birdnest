package config

import "github.com/corvuscrypto/birdnest/requests"

//Config is the config... no tricks here.
var Config *config

var pipeline []func(*requests.Request)

func init() {
	Config = &config{
		make(map[string]interface{}),
	}
	pipeline = make([]func(*requests.Request), 1)
}
