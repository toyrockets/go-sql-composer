package decimal_test

import (
    sql "com.toyrockets/sqlcomposer"
    "testing"
    "fmt"
)

func TestSimpleSelect(t *testing.T) {
    statement := sql.Select("a", "b", "c").From("t1", "t2", "t3").Where(map[string]interface{}{
        "foo": "bar",
        "blarg": 10,
    })
    SQL, values := statement.GenerateSQL()
    fmt.Println(SQL)
    fmt.Println(values)

    result := "select a, b, c from t1, t2, t3 where foo = $1 and blarg = $2"
    if SQL != result  {
        t.Error("Expected ", result, " got ", SQL)
    }

}


