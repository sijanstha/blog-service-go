package sqlutils

import "strings"

// returns TableName, Columns and Value decorator
// EX: tbl_post :: id,name,active :: ?,?,?
func GetTableMetadata(src map[string]interface{}) (string, string, string) {
	columns := src["columns"].([]string)

	values := ""
	columnNames := ""
	for _, col := range columns {
		values += "?,"
		columnNames += col + ","
	}

	return GetTableName(src), strings.TrimSuffix(columnNames, ","), strings.TrimSuffix(values, ",")
}

func GetTableName(src map[string]interface{}) string {
	tblName := src["tableName"]
	return tblName.(string)
}
