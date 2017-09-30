package panel

import (
	"fmt"
)

const (
	// String column data type
	String = iota
	// Int column data type
	Int
	// Float column data type
	Float
	// Bool column data type
	Bool
	// Date column data type
	Date
)

// Convert ...
func (p Panel) Convert(col string, kind int) Panel {
	return Convert(p, col, kind)
}

// Convert ...
func Convert(p Panel, col string, kind int) Panel {
	switch kind {
	case String:
		tempCol := make([]interface{}, len(p[col]))
		for row, val := range p[col] {
			tempCol[row] = fmt.Sprintf("%v", val)
		}
		p[col] = tempCol

	// NEEDS IMPLEMENTATION
	case Int:

	// NEEDS IMPLEMENTATION
	case Float:

	// NEEDS IMPLEMENTATION
	case Bool:

	// NEEDS IMPLEMENTATION
	case Date:

	// NEEDS IMPLEMENTATION
	default:

	}

	return p
}
