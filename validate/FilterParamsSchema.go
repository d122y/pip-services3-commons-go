package validate

import "github.com/pip-services-go/pip-services-commons-go/convert"

func NewFilterParamsSchema() *MapSchema {
	return NewMapSchema(convert.String, nil)
}
