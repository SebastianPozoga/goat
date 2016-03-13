package interactive

import (
	"fmt"
)

type Model struct {
	Fields map[string]Field
}

func (f *Model) Scan() error {
	model := Model{
		Fields: map[string]Field{},
	}
	for true {
		var name string
		fmt.Println("Insert field name (or empty to finish): ", correctTypes)
		fmt.Scanln(&name)
		if name == "" {
			break
		}
		field := Field{}
		field.Scan()
		model.Fields[name] = field
	}
	return nil
}
