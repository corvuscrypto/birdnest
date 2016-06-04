package config

import "github.com/corvuscrypto/birdnest/requests"

//Config is the config... no tricks here.
var Config *config

func init() {
	Config = new(config)
	Config.variables = make(map[string]interface{})
	Config.pipeline = make([]func(*requests.Request), 0)
}
