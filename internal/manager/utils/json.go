package utils

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/mrlutik/kira2.0/internal/config"
	"github.com/mrlutik/kira2.0/internal/logging"
)

func UpdateJsonValue(input []byte, config *config.JsonValue) ([]byte, error) {
	var m map[string]any

	if err := json.Unmarshal(input, &m); err != nil {
		return nil, err
	}

	keys := strings.Split(config.Key, ".")
	if err := setNested(m, keys, config.Value); err != nil {
		return nil, err
	}

	return json.Marshal(m)
}

func setNested(m map[string]any, keys []string, value any) error {
	log := logging.Log

	var exists bool
	for i := 0; i < len(keys)-1; i++ {
		if _, exists = m[keys[i]]; !exists {
			return errors.New("key does not exist: " + keys[i])
		}

		nestedMap, ok := m[keys[i]].(map[string]any)
		if !ok {
			return errors.New("expected map for key: " + keys[i])
		}

		log.Debugf("Found section: %s", keys[i])
		m = nestedMap
	}

	if _, exists = m[keys[len(keys)-1]]; !exists {
		return errors.New("final key does not exist: " + keys[len(keys)-1])
	}

	log.Debugf("Found key: %s\n", keys[len(keys)-1])
	log.Printf("Update old value '%v' -> new value '%v'\n", m[keys[len(keys)-1]], value)
	m[keys[len(keys)-1]] = value
	return nil
}
