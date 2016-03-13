package interactive

import (
	"fmt"
)

type Model struct {
	Fields map[string]Field `json:"fields"`
}

func (m *Model) Scan() error {
	for true {
		var name string
		fmt.Print("Insert field name (or empty to finish): ")
		fmt.Scanln(&name)
		if name == "" {
			break
		}
		field := Field{
			Attrs: []string{},
		}
		if err := field.Scan(); err != nil {
			return err
		}
		m.Fields[name] = field
	}
	return nil
}
