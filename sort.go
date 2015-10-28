package sqlcomposer

import (
    "strings"
)

type SortOrder string

const (
    SortAscending SortOrder = "asc"
    SortDescending SortOrder = "desc"
)

type SortDescriptor struct {
    value SQLExpression
    order SortOrder
}

func (self *SortDescriptor) GenerateSQL() (SQL string, values []interface{}) {
    SQL, values = self.GenerateSQLWithContext(&SQLGenerationContext{Style: DefaultStyle})
    return
}

func (self *SortDescriptor) GenerateSQLWithContext(context *SQLGenerationContext) (SQL string, values []interface{}) {
    SQL, values = self.value.GenerateSQLWithContext(context)
    SQL = strings.Join([]string{SQL, string(self.order)}, " ")
    return
}


func ParseSortDescriptors(values...interface{}) []*SortDescriptor {
    var descriptors []*SortDescriptor
	for _, value := range values {
        var sortExpression SQLExpression

        var descriptor *SortDescriptor

        if stringValue, ok := value.(string); ok {
            sortExpression = &SQLIdentifier{Name: stringValue}
            descriptor = &SortDescriptor{value: sortExpression, order: SortAscending}
        } else if descriptorValue, ok := value.(*SortDescriptor); ok {
            descriptor = descriptorValue
        } else if descriptorValue, ok := value.(SortDescriptor); ok {
            descriptor = &descriptorValue
        } else {
            sortExpression = SQLVariable(value)
            descriptor = &SortDescriptor{value: sortExpression, order: SortAscending}
        }

        descriptors = append(descriptors, descriptor)
	}
    
    return descriptors
}

func Ascending(value string) *SortDescriptor {
    sortExpression := &SQLIdentifier{Name: value}
    return &SortDescriptor{value: sortExpression, order: SortAscending}
}

func Descending(value string) *SortDescriptor {
    sortExpression := &SQLIdentifier{Name: value}
    return &SortDescriptor{value: sortExpression, order: SortDescending}
}
