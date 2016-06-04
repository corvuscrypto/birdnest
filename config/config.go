package config

import "github.com/corvuscrypto/birdnest/requests"

//Can't touch this :>c
type config struct {
	variables map[string]interface{}
	pipeline  []func(*requests.Request)
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
	return c.variables[key].(int)
}

//GetFloat returns the key value as a float64
func (c *config) GetFloat(key string) float64 {
	return c.variables[key].(float64)
}

//GetString returns the key value as a string
func (c *config) GetString(key string) string {
	return c.variables[key].(string)
}

//Set sets a value on the config struct
func (c *config) Set(key string, value interface{}) {
	c.variables[key] = value
}
