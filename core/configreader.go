package scv

import (
	"encoding/json"
	"errors"
	"fmt"
	//"io/ioutil"
)

type TaskConfig struct {
	Trigger TriggerT `json:"trigger"`
	Action  ActionT  `json:"action"`
}

type TriggerT struct {
	Type  string `json:"type"`
	Param string `json:"param"`
}

type ActionT struct {
	Type  string `json:"type"`
	Param string `json:"param"`
}

func ReadConfig(raw []byte) ([]TaskConfig, error) {
	// raw, err := ioutil.ReadFile(path)
	// if err != nil {
	//   return nil, errors.New(fmt.Sprintf("Error opening TaskConfig file: %+v\n", err.Error()))
	// }
	var config []TaskConfig
	if err := json.Unmarshal(raw, &config); err != nil {
		return nil, errors.New(fmt.Sprintf("Error: %+v\n", err.Error()))
	}

	return config, nil
}
