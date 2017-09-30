package panel

import (
	"fmt"
)

// Add a column with data
func Add(p Panel, header string, data interface{}) Panel {
	return p.Add(header, data)
}

// Add a column with data
func (p Panel) Add(header string, data interface{}) Panel {
	switch t := data.(type) {
	case []interface{}:
		p[header] = t
	case []int:
		slice := make([]interface{}, len(t))
		for k, v := range t {
			slice[k] = v
		}
		p[header] = slice
	case []string:
		slice := make([]interface{}, len(t))
		for k, v := range t {
			slice[k] = v
		}
		p[header] = slice
	case []float64:
		slice := make([]interface{}, len(t))
		for k, v := range t {
			slice[k] = v
		}
		p[header] = slice
	case map[int]interface{}:
		slice := make([]interface{}, len(t))
		for k, v := range t {
			slice[k] = v
		}
		p[header] = slice

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
	fmt.Printf("%v is a %T\n\n", c, c)

	for i := range c {
		fmt.Printf("%v is a %T\n\n", c[i], c[i])

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
func Columns(p Panel) (h []string) {
	for header := range p {
		h = append(h, header)
	}
	return
}

// // Columns ...
// func Columns(p Panel, cols map[string]string) Panel {
// 	if cols != nil {
// 		tempPanel := make(Panel)
// 		for header, series := range p {
// 			if val, ok := cols[header]; ok {
// 				tempPanel.Add(val, series)
// 			} else {
// 				tempPanel.Add(header, series)
// 			}
// 		}
// 		return tempPanel
// 	}
// 	return p
// }
