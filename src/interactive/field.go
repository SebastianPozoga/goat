package interactive

import (
	"fmt"
	"github.com/goatcms/goat-core/varutil"
)

var (
	correctTypes = []string{"string", "int", "datetime", "blob"}
)

type Field struct {
	Type     string
	Required bool
	Min      int
	Max      int
	Attrs    []string
}

func (f *Field) Scan() error {
	var attrs string
	var err error
	fmt.Println("Insert field type: ", correctTypes)
	for true {
		fmt.Scanln(&f.Type)
		if varutil.IsArrContainStr(correctTypes, f.Type) {
			break
		}
		fmt.Println("Incorrect field type. Available types: ", correctTypes)
	}
	fmt.Println("Min value of field: ")
	fmt.Scan(&f.Min)
	fmt.Println("Max value of field: ")
	fmt.Scan(&f.Max)
	fmt.Println("Insert field attrs: ")
	fmt.Scan(&attrs)
	f.Attrs, err = varutil.SplitWhite(attrs)
	if err != nil {
		return err
	}
	return nil
}
