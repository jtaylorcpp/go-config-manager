package types

import (
	"encoding/json"
)

type Action = string

const (
	DELETE Action = "delete"
	UPDATE Action = "update"
)

type Transaction struct {
	Version       uint64
	Action        Action
	PartialConfig Config
}

type KeyValue struct {
	Key   string
	Value interface{}
}

func (t Transaction) GetKeyValueSet() ([]KeyValue, error) {
	workingSet := []KeyValue{}

	finishedSet := []KeyValue{}

	var pConfig map[string]interface{}
	jsonErr := json.Unmarshal(t.PartialConfig, &pConfig)
	if jsonErr != nil {
		return []KeyValue{}, jsonErr
	}

	logger.Debugf("partial config to be applied: %#v", pConfig)
	for k, v := range pConfig {
		workingSet = append(workingSet, KeyValue{k, v})
	}

	for len(workingSet) > 0 {
		workingPair := workingSet[0]
		workingSet = workingSet[1:]
		logger.Debugf("working on transaction kv: %#v", workingPair)

		switch workingPair.Value.(type) {
		case map[interface{}]interface{}, map[string]interface{}:
			ValueMap := workingPair.Value.(map[string]interface{})
			for k, v := range ValueMap {
				workingSet = append(workingSet, KeyValue{workingPair.Key + "." + k, v})
			}
		default:
			finishedSet = append(finishedSet, workingPair)
		}
	}

	return finishedSet, nil
}
