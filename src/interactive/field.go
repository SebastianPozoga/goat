package interactive

import (
	"fmt"
	"github.com/goatcms/goat-core/varutil"
	"strings"
)

var (
	correctTypes = []string{"string", "int", "datetime", "blob"}
)

type Field struct {
	Type     string   `json:"type"`
	Required bool     `json:"required"`
	Min      int      `json:"min"`
	Max      int      `json:"max"`
	Attrs    []string `json:"attrs"`
}

func (f *Field) Scan() error {
	var attrs string
	var err error
	fmt.Print(correctTypes, "?: ")
	for true {
		fmt.Scanln(&f.Type)
		f.Type = strings.ToLower(f.Type)
		if varutil.IsArrContainStr(correctTypes, f.Type) {
			break
		}
		fmt.Print(correctTypes, "?: ")
	}
	fmt.Print("Min value of field: ")
	fmt.Scan(&f.Min)
	fmt.Print("Max value of field: ")
	fmt.Scan(&f.Max)

	var isAttrs string
	fmt.Print("Do you want add attribiutes? ")
	fmt.Scan(&isAttrs)
	isAttrs = strings.ToLower(isAttrs)
	if isAttrs == "yes" || isAttrs == "y" {
		fmt.Print("Insert field attrs (separated with space): ")
		fmt.Scan(&attrs)
		f.Attrs, err = varutil.SplitWhite(attrs)
		if err != nil {
			return err
		}
	}
	return nil
}
