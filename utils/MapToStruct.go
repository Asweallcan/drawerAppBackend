package utils

import "encoding/json"

func MapToStruct(m map[string]interface{}, target interface{}) error {
	str, err := json.Marshal(m)
	if err == nil {
		err = json.Unmarshal(str, &target)
	}
	return err
}
