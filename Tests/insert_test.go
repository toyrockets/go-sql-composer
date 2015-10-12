package decimal_test

import (
    sql "com.toyrockets/sqlcomposer"
    "testing"
    "fmt"
)

func TestSimpleInsert(t *testing.T) {
    statement := sql.Insert("user").Values(map[interface{}]interface{}{
        "foo": "bar",
        "blarg": 10,
    })
    SQL, values := statement.GenerateSQL()
    fmt.Println(SQL)
    fmt.Println(values)

    result := "insert into user (blarg, foo) values($1, $2)"
    if SQL != result  {
        t.Error("Expected ", result, " got ", SQL)
    }

}


