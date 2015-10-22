package sqlcomposer_test

import (
    sql "com.toyrockets/sqlcomposer"
    "testing"
)

func TestSimpleSelect(t *testing.T) {
    statement := sql.Select("a", "b", "c").From("t1", "t2", "t3").Where(map[string]interface{}{
        "foo": "bar",
        "blarg": sql.GreaterThan(10),
    }).OrderBy("a")
    SQL, values := statement.GenerateSQL()

    expectedSQL := "select a, b, c from \"t1\", \"t2\", \"t3\" where \"foo\" = $1 and \"blarg\" = $2"
    if SQL != expectedSQL  {
        t.Error("Expected ", expectedSQL, " got ", SQL)
    }

	expectedValues := []interface{}{10, "bar"}
	
	if len(values) != len(expectedValues) {
        t.Error("Expected ", expectedValues, " got ", values)
	} else {
		for index, value := range values {
			if value != expectedValues[index] {
		        t.Error("Expected ", expectedValues, " got ", values)
				break
			}
		}
	}

}


