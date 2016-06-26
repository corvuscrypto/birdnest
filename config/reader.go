package config

import (
	"encoding/json"
	"io/ioutil"
	"math"
	"os"
)

//ReadConfig reads a file and parses the values into the Config struct
func ReadConfig(filepath string) {
	//open the file
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	data, _ := ioutil.ReadAll(f)
	values := make(map[string]interface{})

	err = json.Unmarshal(data, &values)
	if err != nil {
		panic(err)
	}

	for k, v := range values {
		switch t := v.(type) {
		case map[string]interface{}:
			continue
		case []interface{}:
			continue
		case float64:
			if math.Trunc(t) == t {
				Config.Set(k, int(t))
				continue
			}
		}
		Config.Set(k, v)
	}
}
