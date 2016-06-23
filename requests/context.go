package requests

//Context is a map that has special methods added for convenience
type Context map[string]interface{}

//Set sets a key-value association on the map. If true is passed as the third argument, this method will fail to set
//the value if there is already a value set for the given key
func (c Context) Set(key string, value interface{}, noOverwrite ...bool) {

	//if third arg is true check for set value
	if len(noOverwrite) > 0 && noOverwrite[0] {
		if _, ok := c[key]; ok {
			//return to prevent overwrite
			return
		}
	}

	c[key] = value
}

//Get is just an alias for accessing the value at the given key. This is just to keep things consistent. You could just
// as easily just do Context[key]. If you specify a default value, instead of being nil, the value returned will
// be the default when the key is not present.
func (c Context) Get(key string, defaultValue ...interface{}) interface{} {
	if v, ok := c[key]; ok {
		return v
	}
	if len(defaultValue) == 0 {
		return nil
	}
	return defaultValue[0]
}
