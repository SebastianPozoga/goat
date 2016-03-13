package interactive

import (
	"fmt"
	"github.com/goatcms/goat-core/generator"
	"github.com/goatcms/goat-core/varutil"
)

type Model struct {
	Fields map[string]Field
}

type Field struct {
	Type        string
	Validations Validations
}

type Validations struct {
	Required bool
}
