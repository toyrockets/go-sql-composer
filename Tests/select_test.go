package sqlcomposer_test

import (
	sql "com.toyrockets/sqlcomposer"
	"testing"
)

func TestSimpleSelect(t *testing.T) {
	statement := sql.Select("a", "b", "c").From("t1", "t2", "t3").Where(map[string]interface{}{
		"foo":   "bar",
		"blarg": sql.GreaterThan(10),
	}).OrderBy("a")
	SQL, values := statement.GenerateSQL()

	expectedSQL := `select "a", "b", "c" from "t1", "t2", "t3" where "blarg" > $1 and "foo" = $2 order by "a" asc`
	if SQL != expectedSQL {
		t.Error("Expected\n", expectedSQL, "\ngot\n", SQL)
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

func TestSelectWithAliases(t *testing.T) {
	statement := sql.Select(map[string]interface{} {
		"a" : "c1",
		"b": "c.2",
		"c": nil,
	}).From("t1");
	SQL, values := statement.GenerateSQL()

	expectedSQL := `select "a" as "c1", "b" as "c.2", "c" from "t1"`
	if SQL != expectedSQL {
		t.Error("Expected\n", expectedSQL, "\ngot\n", SQL)
	}

	expectedValues := []interface{}{}

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

func TestSelectWithJoin(t *testing.T) {
	statement := sql.Select("a", "b", "c").From("t1", "t2", "t3").Where(map[string]interface{}{
		"foo":   "bar",
		"blarg": sql.GreaterThan(10),
	}).Join("t4", map[string]interface{}{
		"t4.id": &sql.SQLIdentifier{Name: "t3.parent_id"},
	}).OrderBy("a")
	SQL, values := statement.GenerateSQL()

	expectedSQL := `select "a", "b", "c" from "t1", "t2", "t3" join "t4" on "t4"."id" = "t3"."parent_id" where "blarg" > $1 and "foo" = $2 order by "a" asc`
	if SQL != expectedSQL {
		t.Error("Expected\n", expectedSQL, "\ngot\n", SQL)
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

