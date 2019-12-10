package performance

import (
	"encoding/json"
	"github.com/devingen/atama-api/dto"
	"io/ioutil"
	"os"
)

func ReadFile(name string) (*dto.BuildPairsBody, error) {

	jsonFile, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var body dto.BuildPairsBody

	err = json.Unmarshal(byteValue, &body)
	if err != nil {
		return nil, err
	}

	return &body, nil
}

var files = []string{
	//"build-pairs-5-5.json",
	//"build-pairs-7-7.json",
	//"build-pairs-8-8.json",
	"build-pairs-9-9.json",
	//"build-pairs-10-10.json",
	//"build-pairs-15-15.json",
	//"build-pairs-21-21.json",
}
