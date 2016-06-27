package config

import (
	"encoding/json"
	"io/ioutil"
	"math"
	"os"

	"github.com/corvuscrypto/birdnest/logging"
)

//ReadConfig reads a file and parses the values into the Config struct
func ReadConfig(filepath string) {
	logger := logging.GetLogger()
	//open the file
	f, err := os.Open(filepath)
	if err != nil {
		logger.Log(logging.ERROR, err)
		os.Exit(1)
	}
	data, _ := ioutil.ReadAll(f)
	values := make(map[string]interface{})

	err = json.Unmarshal(data, &values)
	if err != nil {
		logger.Log(logging.ERROR, err)
		os.Exit(1)
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
