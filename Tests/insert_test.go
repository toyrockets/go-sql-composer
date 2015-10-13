package sqlcomposer_test

import (
    sql "com.toyrockets/sqlcomposer"
    "testing"
)

func TestSimpleInsert(t *testing.T) {
    statement := sql.Insert("user").Values(map[interface{}]interface{}{
        "foo": "bar",
        "blarg": 10,
    })
    SQL, values := statement.GenerateSQL()

    result := "insert into user (blarg, foo) values($1, $2)"
    if SQL != result  {
        t.Error("Expected ", result, " got ", SQL)
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

func TestInsertWithSubSelect(t *testing.T) {
    selectStatement := sql.Select("id").From("table1").Where(map[interface{}]interface{}{
        "external_id" : 10,
    });

    statement := sql.Insert("user").Values(map[interface{}]interface{}{
        "foo": "bar",
        "blarg": 10,
        "id": selectStatement,
    })
    SQL, values := statement.GenerateSQL()

    expectedSQL := "insert into user (blarg, foo, id) values ($1, $2, select id from table1 where external_id = $3)"
    if SQL != expectedSQL  {
        t.Error("Expected ", expectedSQL, " got ", SQL)
    }

	expectedValues := []interface{}{10, "bar", 10}

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


