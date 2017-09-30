package panel

import (
	"fmt"
	"reflect"
)

// type description map[string]map[string]interface{}

func (p Panel) Describe(cols ...string) map[string]map[string]interface{} {
	return Describe(p, cols)
}

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
