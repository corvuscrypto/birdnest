package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

func TestReadConfig(T *testing.T) {
	testData := map[string]interface{}{
		"testInt":    21,
		"testString": "Hello",
		"testBool":   true,
		"testFloat":  1.23,
		"testInvalid": map[string]interface{}{
			"ignored": 1,
		},
		"testInvalid2": []interface{}{1, 2},
	}

	//transform testData into byte slice
	bitties, _ := json.Marshal(testData)

	//create a temp file and write the data to it json encoded
	f, _ := ioutil.TempFile("", "configTest")
	filename := f.Name()
	defer os.Remove(filename)
	f.Write(bitties)
	f.Close()

	//now read the data into the Config
	ReadConfig(filename)

	//do checks for proper read-in values
	if Config.GetInt("testInt") != 21 {
		T.Errorf("Incorrect value in configuration!")
	}
	if Config.GetString("testString") != "Hello" {
		T.Errorf("Incorrect value in configuration!")
	}
	if Config.GetBool("testBool") != true {
		T.Errorf("Incorrect value in configuration!")
	}
	if Config.GetFloat("testFloat") != 1.23 {
		T.Errorf("Incorrect value in configuration!")
	}

	//now munge the data and try again. We should get a panic
	func() {
		defer func() {
			if recover() == nil {
				T.Errorf("Failed to panic on bad data")
			}
		}()
		bitties = bitties[1:]
		//create a temp file and write the data to it json encoded
		f, _ := ioutil.TempFile("", "configTest")
		filename := f.Name()
		defer os.Remove(filename)
		f.Write(bitties)
		f.Close()

		ReadConfig(filename)
	}()

	//now attempt to read a file that doesn't exist. should panic
	func() {
		defer func() {
			if recover() == nil {
				T.Errorf("Failed to panic on bad data")
			}
		}()
		ReadConfig("1234567890asdfghjkl")
	}()
}
