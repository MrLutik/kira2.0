package utils

import (
	"encoding/json"
	"strings"

	"github.com/mrlutik/kira2.0/internal/config"
	"github.com/mrlutik/kira2.0/internal/logging"
)

type JSONEditor struct {
	log *logging.Logger
}

func NewJSONEditor(logger *logging.Logger) *JSONEditor {
	return &JSONEditor{
		log: logger,
	}
}

func (j *JSONEditor) UpdateJsonValue(input []byte, config *config.JsonValue) ([]byte, error) {
	var mapJSONrepresentation map[string]any

	if err := json.Unmarshal(input, &mapJSONrepresentation); err != nil {
		return nil, err
	}

	keys := strings.Split(config.Key, ".")
	if err := j.setNested(mapJSONrepresentation, keys, config.Value); err != nil {
		return nil, err
	}

	return json.Marshal(mapJSONrepresentation)
}

func (j *JSONEditor) setNested(mapJSONrepresentation map[string]any, keys []string, value any) error {
	var exists bool
	for i := 0; i < len(keys)-1; i++ {
		if _, exists = mapJSONrepresentation[keys[i]]; !exists {
			return &TargetKeyNotFoundError{Key: keys[i]}
		}

		nestedMap, ok := mapJSONrepresentation[keys[i]].(map[string]any)
		if !ok {
			return &ExpectedMapError{Key: keys[i]}
		}

		j.log.Debugf("Found section: %s", keys[i])
		mapJSONrepresentation = nestedMap
	}

	if _, exists = mapJSONrepresentation[keys[len(keys)-1]]; !exists {
		return &TargetKeyNotFoundError{Key: keys[len(keys)-1]}
	}

	j.log.Debugf("Found key: %s\n", keys[len(keys)-1])
	j.log.Infof("Update old value '%v' -> new value '%v'", mapJSONrepresentation[keys[len(keys)-1]], value)
	mapJSONrepresentation[keys[len(keys)-1]] = value
	return nil
}
