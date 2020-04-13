package utils

import (
	"bytes"
	"encoding/json"
	"io"
)

func RequestBody(m interface{}) (io.Reader, error) {
	str, err := json.Marshal(m)
	if err == nil {
		return bytes.NewReader(str), nil
	}
	return nil, err
}
