package fileutils

import "testing"

func TestGetTableInformation(t *testing.T) {
	expected := map[string]interface{}{
		"tableName": "tbl_users",
		"columns":   []string{"id", "name"},
	}
	got, err := GetTableInformation("UserTableDefinitionTest.json")

	if err != nil {
		t.Error("got error: ", err)
	}

	if len(got) != len(expected) {
		t.Errorf("expected len %q, got %q", len(expected), len(got))
	}

	if expected["tableName"] != got["tableName"] {
		t.Errorf("expected  %q, got %q", expected["tableName"], got["tableName"])
	}

	if !ArrayEquals(expected["columns"], got["columns"]) {
		t.Errorf("expected %q, got %q", expected, got)
	}

}

func ArrayEquals(a interface{}, b interface{}) bool {
	x := a.([]string)
	y := b.([]string)
	if len(x) != len(y) {
		return false
	}
	for i, v := range x {
		if v != y[i] {
			return false
		}
	}
	return true
}
