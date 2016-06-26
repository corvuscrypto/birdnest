package config

//Can't touch this :>c
type config struct {
	variables map[string]interface{}
}

//Get returns the key value (or nil) as an interface{} type
func (c *config) Get(key string, fallback ...interface{}) interface{} {
	v, ok := c.variables[key]
	if !ok && fallback != nil {
		return fallback[0]
	}
	return v
}

//GetBool returns the key value as a bool
func (c *config) GetBool(key string, fallback ...bool) bool {
	v, ok := c.variables[key].(bool)
	if !ok && fallback != nil {
		return fallback[0]
	}
	return v
}

//GetInt returns the key value as a int
func (c *config) GetInt(key string, fallback ...int) int {
	v, ok := c.variables[key].(int)
	if !ok && fallback != nil {
		return fallback[0]
	}
	return v
}

//GetFloat returns the key value as a float64
func (c *config) GetFloat(key string, fallback ...float64) float64 {
	v, ok := c.variables[key].(float64)
	if !ok && fallback != nil {
		return fallback[0]
	}
	return v
}

//GetString returns the key value as a string
func (c *config) GetString(key string, fallback ...string) string {
	v, ok := c.variables[key].(string)
	if !ok && fallback != nil {
		return fallback[0]
	}
	return v
}

//Set sets a value on the config struct
func (c *config) Set(key string, value interface{}) {
	c.variables[key] = value
}

//Config is the config... no tricks here.
var Config *config

func init() {
	Config = new(config)
	Config.variables = make(map[string]interface{})
}
