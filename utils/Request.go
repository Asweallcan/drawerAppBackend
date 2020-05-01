package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func Request(method, url string, data map[string]interface{}) (map[string]interface{}, error) {
	var (
		ret      map[string]interface{} = nil
		retError error                  = nil
		isPost                          = false
	)

	if method == "post" || method == "POST" || method == "Post" {
		isPost = true
	}

	if isPost {
		bytedData, err := json.Marshal(data)
		retError = err

		if err == nil {
			requestBody := bytes.NewReader(bytedData)
			res, err := http.Post(url, "application/json", requestBody)
			retError = err

			if err == nil {
				ret, err := unmarshalResp(res.Body)
				return ret, err
			}
		}
	} else {
		res, err := http.Get(fmt.Sprintf(url, mapValues(data)...))
		retError = err

		if err == nil {
			ret, err := unmarshalResp(res.Body)
			return ret, err
		}
	}

	return ret, retError
}

func mapValues(m map[string]interface{}) []interface{} {
	var ret []interface{}
	for _, v := range m {
		ret = append(ret, v)
	}
	return ret
}

func unmarshalResp(resp io.ReadCloser) (map[string]interface{}, error) {
	var ret map[string]interface{}
	body, err := ioutil.ReadAll(resp)

	if err == nil {
		err = json.Unmarshal(body, &ret)
	}

	return ret, err
}
