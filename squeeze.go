package panel

// Squeeze reduces all unique values in p[col] into one string value delimited by 'delim' and grouped by 'groupBy'
// Example: person has a merged master ID ("MID"), but several personal IDs ("PID") for his or her unduplicated selves
//      p.Squeeze("PID", ", ", "MID")
// func (pn Panel) Squeeze(col string, delim string) Panel {
// 	return Squeeze(pn, col, delim)
// }

// Squeeze reduces all unique values in p[col] into one string value delimited by 'delim' and grouped by 'groupBy'
// Example: person has a merged master ID ("MID"), but several personal IDs ("PID") for his or her unduplicated selves
//      p.Squeeze("PID", ", ", "MID")
// func Squeeze(pn Panel, col string, delim string) Panel {
// 	cols := []string{}
// 	for _, c := range pn.Columns() {
// 		if c != col {
// 			cols = append(cols, c)
// 		}
// 	}

// 	groups := pn.GroupBy(cols...)

// 	flat, original := groups[0], groups[1]

// 	// for every row in flat
// 	for f := range flat {
// 		flat[f][col] = []string{}
// 		// for every row in original
// 		for o := range original {
// 			// declare row match tally
// 			var matches int
// 			// for every element in flat row
// 			for xf := range flat[f] {
// 				// for every element in original row
// 				for xo := range original[o] {
// 					// if values match
// 					if flat[f][xf] == original[o][xo] {
// 						// increment match tally
// 						matches++
// 					}
// 				}
// 			}

// 			// if each value matches
// 			if matches == len(flat[f]) {
// 				val := original[o][col]
// 				vals := flat[f][col].([]string)
// 				fmt.Println(vals, val)
// 				flat[f][col] = append(vals, fmt.Sprintf("%v", val))
// 			}
// 		}
// 		// flat[f][col] = strings.Join(flat[f][col].([]string), delim)
// 	}
// 	fmt.Println(flat)

// 	return pn
// }
