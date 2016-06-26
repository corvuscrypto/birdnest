package config

import "testing"

type testStruct struct {
	ID int
}

func TestConfig(T *testing.T) {

	//Test getters with all cases (with and without defaults)
	if Config.Get("test", "default") != "default" {
		T.Errorf("Incorrect value in configuration!")
	}
	if Config.Get("test") != nil {
		T.Errorf("Incorrect value in configuration!")
	}
	if Config.GetInt("test") != 0 {
		T.Errorf("Incorrect value in configuration!")
	}
	if Config.GetInt("test", 1337) != 1337 {
		T.Errorf("Incorrect value in configuration!")
	}
	if Config.GetFloat("test") != 0.0 {
		T.Errorf("Incorrect value in configuration!")
	}
	if Config.GetFloat("test", 3.14) != 3.14 {
		T.Errorf("Incorrect value in configuration!")
	}
	if Config.GetBool("test") {
		T.Errorf("Incorrect value in configuration!")
	}
	if !Config.GetBool("test", true) {
		T.Errorf("Incorrect value in configuration!")
	}
	if Config.GetString("test") != "" {
		T.Errorf("Incorrect value in configuration!")
	}
	if Config.GetString("test", "hello") != "hello" {
		T.Errorf("Incorrect value in configuration!")
	}

	//set a complex config value
	t := new(testStruct)
	t.ID = 42
	Config.Set("test", t)

	//retrieve the complex data
	if Config.Get("test").(*testStruct).ID != 42 {
		T.Errorf("Incorrect value in configuration!")
	}
}
