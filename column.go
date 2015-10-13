package sqlcomposer

import (
	"fmt"
)

type Column struct {
	Name string
	Alias string
}

func (self *Column) GenerateSQL() (SQL string, values []interface{}) {
    SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: DefaultStyle})
    return
}

func (self *Column) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
	if len(self.Alias) == 0 {
		SQL = self.Name
	} else {
		SQL = fmt.Sprintf("\"%s\" as %s", self.Name, self.Alias)
	}

	values = []interface{}{}
	return
}

type ColumnList []Column

func (columns ColumnList) Len() int {
    return len(columns)
}
func (columns ColumnList) Swap(i, j int) {
    columns[i], columns[j] = columns[j], columns[i]
}
func (columns ColumnList) Less(i, j int) bool {
    return columns[i].Name < columns[j].Name
}
