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

}

//Thank Andrew Garrand :)
func TestReaderErrors(T *testing.T) {
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
	//write in bad data
	f.Write(bitties[1:])
	f.Close()

	if ReadConfig(filename) == nil {
		T.Errorf("Unexpected success encountered")
	}

	if ReadConfig("1234567890asdfghjkl") == nil {
		T.Errorf("Unexpected success encountered")
	}

}
