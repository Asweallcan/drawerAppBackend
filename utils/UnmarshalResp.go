package utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func UnmarshalResp(resp io.ReadCloser, target interface{}) error {
	body, err := ioutil.ReadAll(resp)
	if err == nil {
		err = json.Unmarshal(body, &target)
	}
	return err
}
