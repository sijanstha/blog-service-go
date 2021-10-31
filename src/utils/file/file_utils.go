package fileutils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	ErrFileNotFound = errors.New("file not found")
)

func LoadResourceAsString(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", ErrFileNotFound
	}

	return strings.TrimSpace(string(data)), nil
}

func GetTableInformation(path string) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	columnNameList := make([]string, 0)
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, ErrFileNotFound
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var tableInfo map[string]interface{}
	json.Unmarshal([]byte(byteValue), &tableInfo)

	columnsInfoRaw := tableInfo["columns"]
	columns := columnsInfoRaw.([]interface{})
	for _, column := range columns {
		columnInfo := column.(map[string]interface{})
		columnNameList = append(columnNameList, fmt.Sprintf("%v", columnInfo["name"]))
	}
	result["columns"] = columnNameList
	result["tableName"] = tableInfo["tableName"]
	return result, nil
}
