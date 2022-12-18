package jsonutils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	dateutils "github.com/blog-service/src/utils/date"
)

type apiResponse struct {
	Data      interface{} `json:"data"`
	Timestamp string      `json:"timestamp"`
}

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

func OkWithJSONObject(object interface{}) *apiResponse {
	return &apiResponse{
		Data:      object,
		Timestamp: dateutils.GetTodayDateInString(),
	}
}
