package requests

import "net/http"

//context is a map that has special methods added for convenience
type context map[string]interface{}

//Set sets a key-value association on the map. If true is passed as the third argument, this method will fail to set
//the value if there is already a value set for the given key
func (c context) Set(key string, value interface{}, f ...bool) {
	if len(f) == 0 {
		f = append(f, false)
	}
	forceUnique := f[0]
	if _, ok := c[key]; ok && forceUnique {
		//return to prevent overwrite
		return
	}
	c[key] = value
}

//Get is just an alias for accessing the value at the given key. This is just to keep things consistent. You could just
// as easily just do context[key]. If you specify a default value, instead of being nil, the value returned will
// be the default when the key is not present.
func (c context) Get(key string, defaultValue ...interface{}) interface{} {
	if v, ok := c[key]; ok {
		return v
	}
	if len(defaultValue) == 0 {
		return nil
	}
	return defaultValue[0]
}

//Request is a wrapper containing the base request information along with context
type Request struct {
	*http.Request
	Response http.ResponseWriter
	Context  context
}
