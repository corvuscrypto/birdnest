package config

//Config is the config... no tricks here.
var Config *config

//Can't touch this :>c
type config struct {
	data map[string]interface{}
}

//Get returns the key value (or nil) as an interface{} type
func (c *config) Get(key string) interface{} {
	return c.data[key]
}

//GetBool returns the key value as a bool
func (c *config) GetBool(key string) bool {
	return c.data[key].(bool)
}

//GetInt returns the key value as a int
func (c *config) GetInt(key string) int {
	return c.data[key].(int)
}

//GetFloat returns the key value as a float64
func (c *config) GetFloat(key string) float64 {
	return c.data[key].(float64)
}

//GetString returns the key value as a string
func (c *config) GetString(key string) string {
	return c.data[key].(string)
}

//Set sets a value on the config struct
func (c *config) Set(key string, value interface{}) {
	c.data[key] = value
}
