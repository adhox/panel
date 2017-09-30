package panel

import (
	"fmt"
	"math"
	"strings"
)

// GroupByPanel is a special Panel with special methods
type GroupByPanel [][]map[string]interface{} // Panel

func hashify(p Panel) (hash []map[string]interface{}) {
	for row := 0; row < p.Size().Length; row++ {
		m := map[string]interface{}{}
		for col := range p {
			m[col] = p[col][row]
		}
		hash = append(hash, m)
	}
	return
}

// GroupBy returns the length of a column
func (p Panel) GroupBy(cols ...string) GroupByPanel {
	return GroupBy(p, cols...)
}

// GroupBy returns the length of a column
func GroupBy(pnl Panel, cols ...string) GroupByPanel {
	return GroupByPanel{
		// hashify(pnl.Select(cols...).Unique()),
		hashify(pnl.Select(cols).Unique()),
		hashify(pnl),
	}
}

// Max ...
func (gp GroupByPanel) Max(col string) GroupByPanel {
	max := make([]map[string]interface{}, len(gp[0]))
	aggColName := fmt.Sprintf("count_%s", col)
	for g, group := range gp[0] {
		agg := map[string]interface{}{}
		for key, val := range group {
			agg[key] = val
		}
		agg[aggColName] = math.Inf(-1)

		for _, row := range gp[1] {
			pass := true
			for key, val := range group {
				if val != row[key] {
					pass = false
					break
				}
			}
			if pass {
				a := row[col].(float64)
				b := agg[aggColName].(float64)
				if a > b {
					agg[aggColName] = a
				}
			}
		}
		max[g] = agg
	}
	fmt.Println(max)
	return gp
}

// Min ...
func (gp GroupByPanel) Min(col string) GroupByPanel {
	min := make([]map[string]interface{}, len(gp[0]))
	aggColName := fmt.Sprintf("count_%s", col)
	for g, group := range gp[0] {
		agg := map[string]interface{}{}
		for key, val := range group {
			agg[key] = val
		}
		agg[aggColName] = math.Inf(1)

		for _, row := range gp[1] {
			pass := true
			for key, val := range group {
				if val != row[key] {
					pass = false
					break
				}
			}
			if pass {
				a := row[col].(float64)
				b := agg[aggColName].(float64)
				if a < b {
					agg[aggColName] = a
				}
			}
		}
		min[g] = agg
	}
	fmt.Println(min)
	return gp
}

// Sum ...
func (gp GroupByPanel) Sum(col string) GroupByPanel {
	summ := make([]map[string]interface{}, len(gp[0]))
	aggColName := fmt.Sprintf("sum_%s", col)
	for g, group := range gp[0] {
		agg := map[string]interface{}{}
		for key, val := range group {
			agg[key] = val
		}
		agg[aggColName] = 0.0

		for _, row := range gp[1] {
			pass := true
			for key, val := range group {
				if val != row[key] {
					pass = false
					break
				}
			}
			if pass {
				agg[aggColName] = row[col].(float64) + agg[aggColName].(float64)
			}
		}
		summ[g] = agg
	}
	fmt.Println(summ)
	return gp
}

// Count ...
func (gp GroupByPanel) Count(col string) GroupByPanel {
	hist := make([]map[string]interface{}, len(gp[0]))
	aggColName := fmt.Sprintf("count_%s", col)
	for g, group := range gp[0] {
		agg := map[string]interface{}{}
		for key, val := range group {
			agg[key] = val
		}
		agg[aggColName] = 0

		for _, row := range gp[1] {
			pass := true
			for key, val := range group {
				if val != row[key] {
					pass = false
					break
				}
			}
			if pass {
				agg[aggColName] = 1 + agg[aggColName].(int)
			}
		}
		hist[g] = agg
	}
	fmt.Println(hist)
	return gp
}

<<<<<<< HEAD
// Squeeze reduces all unique values in p[col] into one string value delimited by 'delim' and grouped by 'groupBy'
// Example: person has a merged master ID ("MID"), but several personal IDs ("PID") for his or her unduplicated selves
//      p.Squeeze("PID", ", ", "MID")
=======
// Squeeze ...
>>>>>>> b7d41b383db941fc86e7c093029e2c5b2be6c9e8
func (gp GroupByPanel) Squeeze(col string, sep string) GroupByPanel {
	hist := make([]map[string]interface{}, len(gp[0]))
	for g, group := range gp[0] {
		agg := map[string]interface{}{}
		for key, val := range group {
			agg[key] = val
		}
		agg[col] = []string{}

		for _, row := range gp[1] {
			pass := true
			for key, val := range group {
				if val != row[key] {
					pass = false
					break
				}
			}
			if pass {
				x := row[col]
				agg[col] = append(agg[col].([]string), fmt.Sprintf("%v", x))
			}
		}
		agg[col] = strings.Join(agg[col].([]string), sep)
		hist[g] = agg
	}
	fmt.Println(hist)
	return gp
}
