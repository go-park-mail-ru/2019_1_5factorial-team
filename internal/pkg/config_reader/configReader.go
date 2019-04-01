package config_reader

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
)

func ReadConfigFile(path string, filename string, dest interface{}) error {
	//data, err := ioutil.ReadFile("configs/" + filename)
	data, err := ioutil.ReadFile(path + filename)
	if err != nil {
		return errors.Wrap(err, "cant read config file:")
	}

	err = json.Unmarshal(data, &dest)
	if err != nil {
		return errors.Wrap(err, "cant parse config:")
	}

	return nil
}
