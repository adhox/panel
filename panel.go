package panel

<<<<<<< HEAD
import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

=======
>>>>>>> b7d41b383db941fc86e7c093029e2c5b2be6c9e8
// Panel is a basically tabular-shaped data s
type Panel map[string][]interface{}

// New creates a new Panel
func New(data interface{}) Panel {
	p := make(Panel)
	switch t := data.(type) {
	case CSV:
	case XML:
	case JSON:
	case TSV:
	case SQL:

	case map[string][]interface{}:
		p = Panel(p)
	case map[string][]string:
		for head := range t {
			p.Add(head, t[head])
		}
	case map[string][]bool:
		for head := range t {
			p.Add(head, t[head])
		}

	case map[string][]int:
		for head := range t {
			p.Add(head, t[head])
		}

	case map[string][]float64:
		for head := range t {
			p.Add(head, t[head])
		}

	case [][]string:
		columns, body := t[0], t[1:]
		for row := range body {
			for c, col := range columns {
				p[col] = append(p[col], body[row][c])
			}
		}
	case []map[string]interface{}:
		// ***** TODO: check if interface is string or slice of data???
		// columns := getStringKeys(t)
		// fmt.Println("here")
		// // columns, body := t[0], t[1:]
		// for row := range t {
		// 	for c, col := range columns {
		// 		p[col] = append(p[col], body[row][c])
		// 	}
		// }
	default:

	}

	p = p.Clean()
	Meta = Meta.Append(&p)
	return p

}

// Select ...
func (p Panel) Select(c ...interface{}) Panel {
	cols := []string{}
	columns := p.Columns()

	for i := range c {
		switch t := c[i].(type) {
		case string:
			if t == "*" {
				return p
			}
			cols = append(cols, t)
		case int:
			cols = append(cols, columns[t])
		case []interface{}:
			for x := range t {
				switch tt := t[x].(type) {
				case string:
					if tt == "*" {
						return p
					}
					cols = append(cols, tt)
				case int:
					cols = append(cols, columns[tt])
				}
			}
		case []string:
			for x := range t {
				cols = append(cols, t[x])
			}
		case []int:
			for x := range t {
				cols = append(cols, columns[t[x]])
			}
		}
	}

	if len(cols) == 0 {
		return p
	}

	tempPanel := make(Panel)

	for _, col := range cols {
		if col == "*" {
			p.Clone(tempPanel)
			return tempPanel
		}
		tempPanel.Add(col, p[col])
	}
	return tempPanel
}

// Clone ...
func (p Panel) Clone(dst Panel) {
	for header, series := range p {
		dst.Add(header, series)
	}
}

// Clean sanitizes data types within each series
func (p Panel) Clean(cols ...string) Panel {
	return Clean(p, cols...)
}

// Clean sanitizes data types within each series
func Clean(p Panel, cols ...string) Panel {
	tempPanel := make(Panel) // <- implement goroutine for each column
	if len(cols) == 0 {
		for header, series := range p {
			// tally up most frequent
			// data type per column
			cnt := make(map[string]int)
			for _, val := range series {
				switch t := val.(type) {
				case string:
					if f, err := strconv.ParseFloat(t, 64); err == nil {
						ref := reflect.TypeOf(f).String()
						if _, ok := cnt[ref]; ok {
							cnt[ref]++
						} else {
							cnt[ref] = 1
						}
					} else if b, err := strconv.ParseBool(t); err == nil {
						ref := reflect.TypeOf(b).String()
						if _, ok := cnt[ref]; ok {
							cnt[ref]++
						} else {
							cnt[ref] = 1
						}
					} else if code, yes := isDate(t); yes {
						format, _ := dateFormat(code)
						d, _ := time.Parse(format, t)
						ref := reflect.TypeOf(d).String()
						if _, ok := cnt[ref]; ok {
							cnt[ref]++
						} else {
							cnt[ref] = 1
						}
					} else {
						ref := reflect.TypeOf(val).String()
						if _, ok := cnt[ref]; ok {
							cnt[ref]++
						} else {
							cnt[ref] = 1
						}
					}
				default:
					ref := reflect.TypeOf(t).String()
					if _, ok := cnt[ref]; ok {
						cnt[ref]++
					} else {
						cnt[ref] = 1
					}
				}
			}

			var mxHead string
			var mxVal int
			for s, i := range cnt {
				if i > mxVal { // || s == "float64" {
					mxHead = s
					mxVal = i
				}
			}

			switch mxHead {
			case "float64":
				for _, val := range series {
					ref := reflect.TypeOf(val)
					var f float64
					if ref.String() == "string" {
						f, _ = strconv.ParseFloat(val.(string), 64)
					} else {
						f = val.(float64)
					}
					tempPanel[header] = append(tempPanel[header], f)
				}
			case "int":
				for _, val := range series {
					i, _ := strconv.Atoi(val.(string))
					tempPanel[header] = append(tempPanel[header], i)
				}
			case "string":
				for _, val := range series {
					tempPanel[header] = append(tempPanel[header], val.(string))
				}
			case "bool":
				for _, val := range series {
					b, _ := strconv.ParseBool(val.(string))
					tempPanel[header] = append(tempPanel[header], b)
				}
			case "time.Time":
				for _, val := range series {
					ref := reflect.TypeOf(val)
					var d time.Time
					if ref.String() == "string" {
						s := val.(string)
						code, _ := isDate(s)
						format, _ := dateFormat(code)
						d, _ = time.Parse(format, s)
					} else {
						d = val.(time.Time)
					}
					tempPanel[header] = append(tempPanel[header], d)
				}
			default:
				for _, val := range series {
					tempPanel[header] = append(tempPanel[header], val)
				}
			}

		}
		return tempPanel
	}

	// else {
	specs := make(map[string]bool, len(cols))
	for _, col := range cols {
		specs[col] = true
	}

	for header, series := range p {
		if specs[header] {
			// tally up most frequent
			// data type per column
			cnt := make(map[string]int)
			for _, val := range series {
				switch t := val.(type) {
				case string:
					if f, err := strconv.ParseFloat(t, 64); err == nil {
						ref := reflect.TypeOf(f).String()
						if _, ok := cnt[ref]; ok {
							cnt[ref]++
						} else {
							cnt[ref] = 1
						}
					} else if b, err := strconv.ParseBool(t); err == nil {
						ref := reflect.TypeOf(b).String()
						if _, ok := cnt[ref]; ok {
							cnt[ref]++
						} else {
							cnt[ref] = 1
						}
					} else if code, yes := isDate(t); yes {
						format, _ := dateFormat(code)
						d, _ := time.Parse(format, t)
						ref := reflect.TypeOf(d).String()
						if _, ok := cnt[ref]; ok {
							cnt[ref]++
						} else {
							cnt[ref] = 1
						}
					} else {
						ref := reflect.TypeOf(val).String()
						if _, ok := cnt[ref]; ok {
							cnt[ref]++
						} else {
							cnt[ref] = 1
						}
					}
				default:
					ref := reflect.TypeOf(t).String()
					if _, ok := cnt[ref]; ok {
						cnt[ref]++
					} else {
						cnt[ref] = 1
					}
				}
			}

			var mxHead string
			var mxVal int
			for s, i := range cnt {
				if i > mxVal || s == "float64" {
					mxHead = s
					mxVal = i
				}
			}

			switch mxHead {
			case "float64":
				for _, val := range series {
					ref := reflect.TypeOf(val)
					var f float64
					if ref.String() == "string" {
						f, _ = strconv.ParseFloat(val.(string), 64)
					} else {
						f = val.(float64)
					}
					tempPanel[header] = append(tempPanel[header], f)
				}
			case "int":
				for _, val := range series {
					i, _ := strconv.Atoi(val.(string))
					tempPanel[header] = append(tempPanel[header], i)
				}
			case "string":
				for _, val := range series {
					tempPanel[header] = append(tempPanel[header], val.(string))
				}
			case "bool":
				for _, val := range series {
					b, _ := strconv.ParseBool(val.(string))
					tempPanel[header] = append(tempPanel[header], b)
				}
			case "time.Time":
				for _, val := range series {
					ref := reflect.TypeOf(val)
					var d time.Time
					if ref.String() == "string" {
						s := val.(string)
						code, _ := isDate(s)
						format, _ := dateFormat(code)
						d, _ = time.Parse(format, s)
					} else {
						d = val.(time.Time)
					}
					tempPanel[header] = append(tempPanel[header], d)
				}
			default:
				for _, val := range series {
					tempPanel[header] = append(tempPanel[header], val)
				}
			}
		} else {
			tempPanel[header] = series
		}
	}
	return tempPanel
	// }
}

// Describe ...
func (p Panel) Describe(cols ...string) map[string]map[string]interface{} {
	return Describe(p, cols)
}

// Describe ...
func Describe(p Panel, cols []string) map[string]map[string]interface{} {
	desc := make(map[string]map[string]interface{})
	if len(cols) != 0 {
		for _, c := range cols {
			subdesc := make(map[string]interface{})
			ref := reflect.TypeOf(p[c][0])
			subdesc["type"] = ref.String()

			switch rs := ref.String(); rs {
			case "bool":
				fmt.Printf("%v => %v\n", c, rs)

			case "string":
				fmt.Printf("%v => %v\n", c, rs)
			case "float64", "int":
				subdesc["max"] = p.Max(c).(float64)
				subdesc["min"] = p.Min(c).(float64)
				subdesc["mean"] = p.Mean(c)
				subdesc["range"] = p.Range(c) // .(float64)
				subdesc["mode"] = p.Mode(c)
				subdesc["sum"] = p.Sum(c)
			}

			desc[c] = subdesc
		}
	} else {
		for h, s := range p {
			subdesc := make(map[string]interface{})
			ref := reflect.TypeOf(s[0])
			subdesc["type"] = ref.String()

			switch rs := ref.String(); rs {
			case "bool":
				fmt.Printf("%v => %v\n", h, rs)

			case "string":
				fmt.Printf("%v => %v\n", h, rs)
			case "float64", "int":

				subdesc["max"] = p.Max(h).(float64)
				subdesc["min"] = p.Min(h).(float64)
				subdesc["mean"] = p.Mean(h)
				subdesc["range"] = p.Range(h)
				subdesc["mode"] = p.Mode(h) // change MODE to all keys are strings
				subdesc["sum"] = p.Sum(h)
			}

			desc[h] = subdesc
		}
	}
	return desc
}

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

// Add a column with data
func Add(p Panel, column string, data interface{}) Panel {
	return p.Add(column, data)
}

// Add a column with data
func (p Panel) Add(column string, data interface{}) Panel {
	switch t := data.(type) {
	case []interface{}:
		p[column] = t
	case []int:
		slice := make([]interface{}, len(t))
		for k, v := range t {
			slice[k] = v
		}
		p[column] = slice
	case []string:
		slice := make([]interface{}, len(t))
		for k, v := range t {
			slice[k] = v
		}
		p[column] = slice
	case []float64:
		slice := make([]interface{}, len(t))
		for k, v := range t {
			slice[k] = v
		}
		p[column] = slice
	case map[int]interface{}:
		slice := make([]interface{}, len(t))
		for k, v := range t {
			slice[k] = v
		}
		p[column] = slice

	}
	// todo: add map[int]interface...
	return p
}

// Rename a column
func (p Panel) Rename(c ...interface{}) Panel {
	return Rename(p, c)
}

// Rename columns
// Quick if only renaming one column:
// 		df.Rename("old", "new")
// If renaming more than one column:
// 		df.Rename(map[string]string{"old1":"new1", "old2":"new2")
func Rename(p Panel, c ...interface{}) Panel {
	removals := []string{}
	pairs := map[string]string{}
	// fmt.Printf("%v is a %T\n\n", c, c)

	for i := range c {
		// fmt.Printf("%v is a %T\n\n", c[i], c[i])

		switch t := c[i].(type) {
		case map[string]string:
			// in case there are multiple maps (e.g. ...map[string]string)
			for k, v := range t {
				pairs[k] = v
			}
		case []interface{}:
			// strings only
			indexCount, key, val := 0, "", ""
			for ii := range t {
				switch tt := t[ii].(type) {
				case map[string]string:
					// in case there are multiple maps (e.g. ...map[string]string)
					for k, v := range tt {
						pairs[k] = v
					}
				case string:
					switch indexCount {
					case 0:
						key = tt
						indexCount++
					case 1:
						val = tt
						indexCount++
					default:
						break
					}
				}
			}

			if indexCount > 1 {
				pairs[key] = val
			}
		case []string:
			k, v := t[0], t[1]
			pairs[k] = v
		case [2]string:
			k, v := t[0], t[1]
			pairs[k] = v
		default:
			return p
		}
	}

	// switch t := cols.(type) {
	// case nil:
	// 	return p
	// case []string:
	// 	from, to := t[0], t[1]
	// 	p.Add(to, p[from])
	// 	removals = append(removals, from)
	// case [2]string:
	// 	from, to := t[0], t[1]
	// 	p.Add(to, p[from])
	// 	removals = append(removals, from)
	// case map[string]string:

	fmt.Println(len(pairs), pairs)
	for from, to := range pairs {
		p.Add(to, p[from])
		removals = append(removals, from)
	}
	// }
	return p.Remove(removals...)
}

// Remove an entire column
func (p Panel) Remove(cols ...string) Panel {
	return Remove(p, cols...)
}

// Remove an entire column
func Remove(p Panel, cols ...string) Panel {
	if len(cols) != 0 {
		for _, col := range cols {
			delete(p, col)
		}
	}
	return p
}

// Columns returns a slice of strings for the column names
func (p Panel) Columns() []string {
	return Columns(p)
}

// Columns returns a slice of strings for the column names
func Columns(p Panel) (cols []string) {
	for header := range p {
		cols = append(cols, header)
	}
	return
}

// FillMissing ...
func (p Panel) FillMissing() Panel {
	return FillMissing(p)
}

// FillMissing ...
func FillMissing(p Panel) Panel {
	for col := range p {
		p = p.Map(col, func(x interface{}) interface{} {
			switch x.(type) {
			case nil:
				return ""
			default:
				return x
			}
		})
	}

	// for col, series := range p {
	// 	for _, line := range series {
	// 		fmt.Printf("%s: %v: %v\n", col, reflect.TypeOf(line), line)
	// 	}
	// }

	return p
}

// Distinct ...
func (p Panel) Distinct() Panel {
	return Unique(p)
}

// Unique ...
func (p Panel) Unique() Panel {
	return Unique(p)
}

// Unique ...
func Unique(p Panel) Panel {
	var tempVal []string
	var tempCol []string

	for col, srs := range p {
		tempCol = append(tempCol, col)

		for row, val := range srs {
			lenTempVal := len(tempVal)

			if row >= lenTempVal {
				tempVal = append(tempVal, fmt.Sprintf("%v", val))
			} else {
				tempVal[row] += fmt.Sprintf("|%v", val)

			}
		}
	}

	m := make(map[string]bool)

	for _, row := range tempVal {
		m[row] = true
	}

	var final [][]string

	for val := range m {
		final = append(final, strings.Split(val, "|"))
	}

	pnl := Panel{}

	for colNum, colName := range tempCol {
		for _, rowVals := range final {

			pnl[colName] = append(pnl[colName], rowVals[colNum])
		}
	}

	return pnl
}
