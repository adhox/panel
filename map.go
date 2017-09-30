package panel


import "time"

func (p Panel) Map(col string, fn interface{}) Panel {
	tempPanel := make(Panel)
	for header, series := range p {
		if header != col {
			tempPanel.Add(header, series)
		} else {
			c := make([]interface{}, len(series))
			tempPanel.Add(header, c)
		}

	}

	switch t := fn.(type) {
	case func(string) string:
		for key, val := range p[col] {
			tempPanel[col][key] = t(val.(string))
		}
	case func(int) int:
		for key, val := range p[col] {
			tempPanel[col][key] = t(val.(int))
		}
	case func(float64) float64:
		for key, val := range p[col] {
			tempPanel[col][key] = t(val.(float64))
		}
	case func(uint) uint:
		for key, val := range p[col] {
			tempPanel[col][key] = t(val.(uint))
		}
	case func(bool) bool:
		for key, val := range p[col] {
			tempPanel[col][key] = t(val.(bool))
		}
	case func(string) int:
		for key, val := range p[col] {
			tempPanel[col][key] = t(val.(string))
		}
	case func(float64) int:
		for key, val := range p[col] {
			tempPanel[col][key] = t(val.(float64))
		}
	case func(uint) int:
		for key, val := range p[col] {
			tempPanel[col][key] = t(val.(uint))
		}
	case func(bool) int:
		for key, val := range p[col] {
			tempPanel[col][key] = t(val.(bool))
		}
	case func(int) string:
		for key, val := range p[col] {
			tempPanel[col][key] = t(val.(int))
		}
	case func(float64) string:
		for key, val := range p[col] {
			tempPanel[col][key] = t(val.(float64))
		}
	case func(uint) string:
		for key, val := range p[col] {
			tempPanel[col][key] = t(val.(uint))
		}
	case func(bool) string:
		for key, val := range p[col] {
			tempPanel[col][key] = t(val.(bool))
		}
	case func(string) float64:
		for key, val := range p[col] {
			tempPanel[col][key] = t(val.(string))
		}
	case func(int) float64:
		for key, val := range p[col] {
			tempPanel[col][key] = t(val.(int))
		}
	case func(uint) float64:
		for key, val := range p[col] {
			tempPanel[col][key] = t(val.(uint))
		}
	case func(bool) float64:
		for key, val := range p[col] {
			tempPanel[col][key] = t(val.(bool))
		}
	case func(string) uint:
		for key, val := range p[col] {
			tempPanel[col][key] = t(val.(string))
		}
	case func(int) uint:
		for key, val := range p[col] {
			tempPanel[col][key] = t(val.(int))
		}
	case func(float64) uint:
		for key, val := range p[col] {
			tempPanel[col][key] = t(val.(float64))
		}
	case func(bool) uint:
		for key, val := range p[col] {
			tempPanel[col][key] = t(val.(bool))
		}
	case func(string) bool:
		for key, val := range p[col] {
			tempPanel[col][key] = t(val.(string))
		}
	case func(int) bool:
		for key, val := range p[col] {
			tempPanel[col][key] = t(val.(int))
		}
	case func(float64) bool:
		for key, val := range p[col] {
			tempPanel[col][key] = t(val.(float64))
		}
	case func(uint) bool:
		for key, val := range p[col] {
			tempPanel[col][key] = t(val.(uint))
		}
	case func(string) interface{}:
		for key, val := range p[col] {
			tempPanel[col][key] = t(val.(string))
		}
	case func(int) interface{}:
		for key, val := range p[col] {
			tempPanel[col][key] = t(val.(int))
		}
	case func(float64) interface{}:
		for key, val := range p[col] {
			tempPanel[col][key] = t(val.(float64))
		}
	case func(uint) interface{}:
		for key, val := range p[col] {
			tempPanel[col][key] = t(val.(uint))
		}
	case func(bool) interface{}:
		for key, val := range p[col] {
			tempPanel[col][key] = t(val.(bool))
		}
	case func(interface{}) string:
		for key, val := range p[col] {
			tempPanel[col][key] = t(val)
		}
	case func(interface{}) int:
		for key, val := range p[col] {
			tempPanel[col][key] = t(val)
		}
	case func(interface{}) float64:
		for key, val := range p[col] {
			tempPanel[col][key] = t(val)
		}
	case func(interface{}) uint:
		for key, val := range p[col] {
			tempPanel[col][key] = t(val)
		}
	case func(interface{}) bool:
		for key, val := range p[col] {
			tempPanel[col][key] = t(val)
		}

	case func(string) time.Time:
		for key, val := range p[col] {
			tempPanel[col][key] = t(val.(string))
		}
	case func(int) time.Time:
		for key, val := range p[col] {
			tempPanel[col][key] = t(val.(int))
		}
	case func(float64) time.Time:
		for key, val := range p[col] {
			tempPanel[col][key] = t(val.(float64))
		}
	case func(uint) time.Time:
		for key, val := range p[col] {
			tempPanel[col][key] = t(val.(uint))
		}
	case func(bool) time.Time:
		for key, val := range p[col] {
			tempPanel[col][key] = t(val.(bool))
		}
	default:
		def := t.(func(interface{}) interface{})
		for key, val := range p[col] {
			tempPanel[col][key] = def(val)
		}
	}

	return tempPanel
}
