package common

import (
	"encoding/json"
)

func Serialize(object interface{}) ([]byte, error) {
	data, err := json.Marshal(object)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func Deserialize(serializedBytes []byte, object interface{}) error {
	if len(serializedBytes) == 0 {
		return nil
	}

	err := json.Unmarshal(serializedBytes, object)

	if err != nil {
		return err
	}

	return nil
}
