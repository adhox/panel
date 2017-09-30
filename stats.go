package panel

import (
	"fmt"
	"math"
	"strconv"
)

// Size is a struct that holds the Length and Width of a frame
type Size struct {
	Length,
	Width int
}

// Size returns a Size struct of a given frame
func (p Panel) Size() Size {
	size := Size{}
	for _, col := range p {
		size.Width++
		if len(col) > size.Length {
			size.Length = len(col)
		}
	}
	return size
}

// Min produces minimum value for a given column
func (p Panel) Min(col string) interface{} {
	return Min(p, col)
}

// Min produces minimum value for a given column
func Min(p Panel, col string) interface{} {
	switch a := p[col][0].(type) {
	case int:
		mn := float64(a)
		for _, val := range p[col][1:] {
			if b := float64(val.(int)); b < mn {
				mn = b
			}
		}
		return mn
	case float64:
		mn := a
		for _, val := range p[col][1:] {
			if b := val.(float64); b < mn {
				mn = b
			}
		}
		return mn
	case string:
		mn := a
		ln := len(a)
		for _, val := range p[col][1:] {
			if b := len(val.(string)); b < ln {
				mn = val.(string)
				ln = b
			}
		}
		return mn
	default:
		// return

	}

	return nil
}

// Max ...
func (p Panel) Max(col string) interface{} {
	return Max(p, col)
}

// Max ...
func Max(p Panel, col string) interface{} {
	// fmt.Printf("%v => %T\n", p[col][0], p[col][0])
	switch a := p[col][0].(type) {
	case int:
		var mx float64
		for _, val := range p[col] {
			if b := float64(val.(int)); b > mx {
				mx = b
			}
		}
		return mx
	case float64:
		var mx float64
		for _, val := range p[col] {
			if b := val.(float64); b > mx {
				mx = b
			}
		}
		return mx

	case string:
		mx := a
		ln := len(a)
		for _, val := range p[col][1:] {
			if b := len(val.(string)); b > ln {
				mx = val.(string)
				ln = b
			}
		}
		return mx
	}

	return nil
}

// Count returns the length of a column
func (p Panel) Count(col string) float64 {
	return Count(p, col)
}

// Count returns the length of a column
func Count(p Panel, col string) float64 {
	return float64(len(p[col]))
}

// Sum ...
func (p Panel) Sum(col string) float64 {
	return Sum(p, col)
}

// Sum ...
func Sum(p Panel, col string) (sum float64) {
	switch p[col][0].(type) {
	case int:
		for _, val := range p[col] {
			sum += float64(val.(int))
		}
	case float64:
		for _, val := range p[col] {
			sum += val.(float64)
		}
	}
	return
}

// Mode ...
func (p Panel) Mode(col string) map[string]int {
	return Mode(p, col)
}

// Mode ...
func Mode(p Panel, col string) map[string]int {
	counter := make(map[string]int)
	for _, val := range p[col] {
		vval := fmt.Sprintf("%v", val)
		if counter[vval] > 0 {
			counter[vval]++
		} else {
			counter[vval] = 1
		}
	}

	mx := 0
	var top map[string]int

	for i, c := range counter {
		if c > mx {
			mx = c
			top = map[string]int{i: c}
		} else if c == mx {
			top[i] = c
		}
	}
	return top
}

// Round ...
func (p Panel) Round(col string, places ...int) Panel {
	l := len(places)
	switch {
	case l > 0:
		for row, val := range p[col] {
			p[col][row] = Round(val, places[0])
		}
	default:
		for row, val := range p[col] {
			p[col][row] = Round(val, 0)
		}
	}
	return p
}

// Round ...
func Round(v interface{}, places ...int) float64 {
	pl := 0
	if len(places) > 0 {
		pl = places[0]
	}

	shift := math.Pow(10, float64(pl))

	switch t := v.(type) {
	case int:
		return float64(t)
	case string:
		pf, _ := strconv.ParseFloat(t, 64)
		f := math.Floor((pf * shift) + 0.5)
		return f / shift
	default:
		f := math.Floor((v.(float64) * shift) + 0.5)
		return f / shift
	}

}

// Median is NOT IMPLEMENTED
func (p Panel) Median(col string) float64 {
	return 0.0
}

// Median is NOT IMPLEMENTED
func Median(p Panel, col string) float64 {
	return 0.0
}

// Mean ...
func (p Panel) Mean(col string) float64 {
	return Mean(p, col)
}

// Mean ...
func Mean(p Panel, col string) float64 {
	return p.Sum(col) / p.Count(col)
}

/////////////////////////////////////////////////////////

// // Percentiles
// func (p Panel) Percentile(col string) Panel {
// 	// TODO
// 	// - JOIN percentiles with associated values in Panel
// 	p.AddSeries(fmt.Sprintf("%s_percentiles", col), Percentile(p[col]))
// 	return p
// }

// func Rank(nums interface{}) map[interface{}]int {
// 	switch t := nums.(type) {
// 	case []int:
// 		m := make(map[interface{}]int, len(t))
// 		sort.Ints(t)
// 		for rank, num := range t {
// 			m[num] = rank + 1
// 		}
// 		return m
// 	case []string:
// 		m := make(map[interface{}]int, len(t))
// 		sort.Strings(t)
// 		for rank, num := range t {
// 			m[num] = rank + 1
// 		}
// 		return m
// 	default:
// 		d := make(map[interface{}]int)
// 		return d
// 	}
// }

// func Percentile(nums interface{}) map[interface{}]float64 {

// 	switch t := nums.(type) {
// 	case []string:
// 		sort.Strings(t)
// 		unitSize := 1.0 / float64(len(t))
// 		freqs := Frequency(t)
// 		m := make(map[interface{}]float64, len(t))

// 		for below, val := range t {
// 			sizeBelow := unitSize * float64(below)
// 			midSizeVal := float64(freqs[val]) * unitSize * 0.5
// 			percentRank := sizeBelow + midSizeVal
// 			if _, exists := m[val]; m[val] > percentRank || !exists {

// 				m[val] = percentRank
// 			}
// 		}
// 		return m

// 	case []int:
// 		sort.Ints(t)
// 		unitSize := 1.0 / float64(len(t))
// 		freqs := Frequency(t)
// 		m := make(map[interface{}]float64, len(t))

// 		for below, val := range t {
// 			sizeBelow := unitSize * float64(below)
// 			midSizeVal := float64(freqs[val]) * unitSize * 0.5
// 			percentRank := sizeBelow + midSizeVal
// 			if _, exists := m[val]; m[val] > percentRank || !exists {

// 				m[val] = percentRank
// 			}
// 		}
// 		return m
// 	}
// 	return make(map[interface{}]float64)
// }

// func Frequency(nums interface{}) map[interface{}]int {
// 	switch t := nums.(type) {
// 	case []int:
// 		cnt := make(map[interface{}]int, len(t))
// 		for _, v := range t {
// 			if val, exists := cnt[v]; exists {
// 				cnt[v] = val + 1
// 			} else {
// 				cnt[v] = 1
// 			}
// 		}
// 		return cnt

// 	}

// 	return make(map[interface{}]int)

// }
