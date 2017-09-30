package panel

// Filter ...
func (p Panel) Filter(cols string, fn interface{}) Panel {
	tempPanel := make(Panel)
	for header := range p {

		c := []interface{}{}
		tempPanel.Add(header, c)
	}

	switch t := fn.(type) {
	case func(bool) bool:
		for key, val := range p[cols] {
			if t(val.(bool)) {
				for header := range p {
					tempPanel[header] = append(tempPanel[header], p[header][key])
				}
			}
		}
		return tempPanel

	case func(string) bool:
		for key, val := range p[cols] {
			if t(val.(string)) {
				for header := range p {
					tempPanel[header] = append(tempPanel[header], p[header][key])
				}
			}
		}
		return tempPanel

	case func(int) bool:
		for key, val := range p[cols] {
			if t(val.(int)) {
				for header := range p {
					tempPanel[header] = append(tempPanel[header], p[header][key])
				}
			}
		}
		return tempPanel

	case func(float64) bool:
		for key, val := range p[cols] {
			if t(val.(float64)) {
				for header := range p {
					tempPanel[header] = append(tempPanel[header], p[header][key])
				}
			}
		}
		return tempPanel

	case func(uint) bool:
		for key, val := range p[cols] {
			if t(val.(uint)) {
				for header := range p {
					tempPanel[header] = append(tempPanel[header], p[header][key])

				}
			}
		}
		return tempPanel

	case func(interface{}) bool:
		for key, val := range p[cols] {
			if t(val) {
				for header := range p {
					tempPanel[header] = append(tempPanel[header], p[header][key])
				}
			}
		}
		return tempPanel

	default:
		def := t.(func(interface{}) bool)
		for key, val := range p[cols] {
			if def(val) {
				for header := range p {
					tempPanel[header] = append(tempPanel[header], p[header][key])
				}
			}
		}
		return tempPanel
	}

	return tempPanel
}
