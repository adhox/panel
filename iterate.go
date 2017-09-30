package panel

// Row  ...
type Row map[string]interface{}

// Iter ...
func (p Panel) Iter(fn interface{}) Panel {
	var t1 []map[string]interface{}
	for i := 0; i < p.Size().Length; i++ {
		row := map[string]interface{}{}
		for col, val := range p {
			row[col] = val[i]
		}
		t1 = append(t1, row)
	}

	t2 := make([]Row, len(t1))
	switch f := fn.(type) {
	case func(Row) Row:
		for i, x := range t1 {
			t2[i] = f(Row(x))
		}
	case func(map[string]interface{}) map[string]interface{}:
		for i, x := range t1 {
			t2[i] = Row(f(x))
		}
	case func(Row) map[string]interface{}:
		for i, x := range t1 {
			t2[i] = Row(f(Row(x)))
		}
	case func(map[string]interface{}) Row:
		for i, x := range t1 {
			t2[i] = Row(f(x))
		}
	}

	tempPanel := New(nil)
	for i := range t2 {
		row := t2[i]
		for k, v := range row {
			tempPanel[k] = append(tempPanel[k], v)
		}
	}
	return tempPanel
}
