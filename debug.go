package panel

import (
	"fmt"
	"reflect"
)

// Dtypes displays data types
func (p Panel) Dtypes(cols ...string) {
	Dtypes(p, cols...)
}

// TODO
// limit print to head in heads

// Dtypes displays data types
func Dtypes(p Panel, cols ...string) {
	if cols[0] == "*" {
		cols = p.Columns()
	}

	// for col, series := range p.Select(cols...) {
	for col, series := range p.Select(cols) {
		for row, val := range series {

			switch t := val.(type) {
			case float64:
				ref := reflect.TypeOf(t)
				fmt.Printf("col(%s) row(%d) val(%v) ref(%s)\n", col, row, val, ref)
			case int:
				ref := reflect.TypeOf(t)
				fmt.Printf("col(%s) row(%d) val(%v) ref(%s)\n", col, row, val, ref)
			case bool:
				ref := reflect.TypeOf(t)
				fmt.Printf("col(%s) row(%d) val(%v) ref(%s)\n", col, row, val, ref)
			case string:
				ref := reflect.TypeOf(t)
				fmt.Printf("col(%s) row(%d) val(%v) ref(%s)\n", col, row, val, ref)
			default:
				fmt.Printf("col(%s) row(%d) val(%v) ref(%s)\n", col, row, val, reflect.TypeOf(val))
			}
		}
	}
}
