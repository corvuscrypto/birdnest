package config

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//ReadConfig reads a file and parses the values into the Config struct
func ReadConfig(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		panic("Couldn't find the config file at the path specified!")
	}
	scanner := bufio.NewScanner(file)
	var lineNum = 0
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		if line[0] == '#' {
			continue
		}
		var parts = strings.SplitN(line, "=", 1)
		parts = append(strings.SplitN(parts[0], ":", 1), parts[1])
		parseValues(parts, lineNum, filepath)
	}
}

func parseValues(parts []string, lineNum int, filepath string) {
	switch parts[1] {
	case "string":
		Config.Set(parts[0], parts[2])
		break
	case "int":
		value, err := strconv.Atoi(parts[2])
		if err != nil {
			panic(err)
		}
		Config.Set(parts[0], value)
		break
	case "float64":
		value, err := strconv.ParseFloat(parts[2], 64)
		if err != nil {
			panic(err)
		}
		Config.Set(parts[0], value)
		break
	case "bool":
		value, err := strconv.ParseBool(parts[2])
		if err != nil {
			panic(err)
		}
		Config.Set(parts[0], value)
		break
	default:
		panic(fmt.Sprintf("Encountered unexpected type '%s' at line %d of %s", parts[1], lineNum, filepath))
	}
}
