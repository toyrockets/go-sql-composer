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
    fmt.Println(SQL)
    fmt.Println(values)

    result := "insert into user (blarg, foo, id) values($1, $2, select id from table1 where external_id = $3)"
    if SQL != result  {
        t.Error("Expected ", result, " got ", SQL)
    }

}


