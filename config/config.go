package config

import "github.com/corvuscrypto/birdnest/requests"

//Can't touch this :>c
type config struct {
	variables map[string]interface{}
	pipeline  []func(*requests.Request)
}

func NewConfig() *config {
	Config = new(config)
	Config.variables = make(map[string]interface{})
	Config.pipeline = make([]func(*requests.Request), 0)
	return Config
}

//Get returns the key value (or nil) as an interface{} type
func (c *config) Get(key string) interface{} {
	return c.variables[key]
}

//GetBool returns the key value as a bool
func (c *config) GetBool(key string) bool {
	return c.variables[key].(bool)
}

//GetInt returns the key value as a int
func (c *config) GetInt(key string) int {
	if c.variables[key] == nil {
		return 0
	}
	return c.variables[key].(int)
}

//GetFloat returns the key value as a float64
func (c *config) GetFloat(key string) float64 {
	return c.variables[key].(float64)
}

//GetString returns the key value as a string
func (c *config) GetString(key string) string {
	if c.variables[key] == nil {
		return ""
	}
	return c.variables[key].(string)
}

//Set sets a value on the config struct
func (c *config) Set(key string, value interface{}) {
	c.variables[key] = value
}
