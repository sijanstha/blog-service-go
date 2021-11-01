package jsonutils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ShouldBindJSON(request *http.Request, object interface{}) error {
	requestBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return err
	}
	defer request.Body.Close()

	if err := json.Unmarshal(requestBody, &object); err != nil {
		return err
	}

	return nil
}
