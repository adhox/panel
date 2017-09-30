package panel

// JoinType provides an easy way to pass the correct join type parameter
var JoinType = struct {
	Inner     string
	Left      string
	Right     string
	LeftOnly  string
	RightOnly string
	Full      string
	Cross     string
}{
	"inner",
	"left",
	"right",
	"leftOnly",
	"rightOnly",
	"full",
	"cross", // need implementation
}

// Join ...
func Join(p1, p2 Panel, joinType string, on ...string) Panel {
	hash := []map[string]interface{}{}
	p := New(nil)
	headers := []string{}

	// get headers ready
	for _, h := range append(p1.Columns(), p2.Columns()...) {
		if !stringInSlice(h, headers) {
			headers = append(headers, h)
		}
	}

	switch joinType {
	case JoinType.Inner:
		for r1 := 0; r1 < p1.Size().Length; r1++ {
			for r2 := 0; r2 < p2.Size().Length; r2++ {

				var matches int
				for i := range on {

					if p1[on[i]][r1] == p2[on[i]][r2] {
						matches++
					}
				}

				if len(on) == matches {
					m := map[string]interface{}{}
					for col := range p1 {
						m[col] = p1[col][r1]
					}
					for col := range p2 {
						m[col] = p2[col][r2]
					}
					hash = append(hash, m)
				}
			}
		}

		for _, h := range headers {
			p[h] = []interface{}{}
		}

		for row := 0; row < len(hash); row++ {
			for _, h := range headers {
				p[h] = append(p[h], hash[row][h])
			}
		}

	case JoinType.Left:
		for r1 := 0; r1 < p1.Size().Length; r1++ {
			joinCount := 0
			for r2 := 0; r2 < p2.Size().Length; r2++ {

				var matches int
				for i := range on {

					if p1[on[i]][r1] == p2[on[i]][r2] {
						matches++
					}
				}
				if len(on) == matches {
					m := map[string]interface{}{}
					for col := range p1 {
						m[col] = p1[col][r1]
					}
					for col := range p2 {
						m[col] = p2[col][r2]
					}
					hash = append(hash, m)
					joinCount++

				}
			}

			// if no joins for a given row in table A,
			// produce at least one row with empty rows from table B
			if joinCount == 0 {
				m := map[string]interface{}{}
				for col := range p1 {
					m[col] = p1[col][r1]
				}
				for col := range p2 {
					if !stringInSlice(col, on) {
						m[col] = nil
					}
				}
				hash = append(hash, m)
				joinCount++
			}
		}

		for _, h := range headers {
			p[h] = []interface{}{}
		}

		for row := 0; row < len(hash); row++ {
			for _, h := range headers {
				p[h] = append(p[h], hash[row][h])
			}
		}

	case JoinType.LeftOnly:
		for r1 := 0; r1 < p1.Size().Length; r1++ {
			for r2 := 0; r2 < p2.Size().Length; r2++ {
				var matches int
				for i := range on {

					if p1[on[i]][r1] == p2[on[i]][r2] {
						matches++
					}
				}
				if len(on) > matches {
					m := map[string]interface{}{}
					for col := range p1 {
						m[col] = p1[col][r1]
					}
					hash = append(hash, m)
				}
			}
		}

		for _, h := range headers {
			p[h] = []interface{}{}
		}

		for row := 0; row < len(hash); row++ {
			for _, h := range headers {
				p[h] = append(p[h], hash[row][h])
			}
		}

	case JoinType.Right:
		return Join(p2, p1, "left", on...)

	case JoinType.RightOnly:
		return Join(p2, p1, "leftOnly", on...)

	case JoinType.Full:
		m1 := Join(p1, p2, "left", on...)
		m2 := Join(p1, p2, "right", on...)
		return m1.Concat(m2).Unique()

	case JoinType.Cross:
		// TODO!!! Rename matching col names
		// since there will some with conflicting values
		for r1 := 0; r1 < p1.Size().Length; r1++ {
			for r2 := 0; r2 < p2.Size().Length; r2++ {
				m := map[string]interface{}{}
				for col := range p1 {
					m[col] = p1[col][r1]
				}
				for col := range p2 {
					m[col] = p2[col][r2]
				}
				hash = append(hash, m)
			}
		}

		for _, h := range headers {
			p[h] = []interface{}{}
		}

		for row := 0; row < len(hash); row++ {
			for _, h := range headers {
				p[h] = append(p[h], hash[row][h])
			}
		}

	// // default is rerun as inner join
	default:
		if joinType != "inner" {
			return Join(p1, p2, "inner", on...)
		}
	}
	return p
}

// InnerJoin ...
func (p1 Panel) InnerJoin(p2 Panel, on ...string) Panel {
	return Join(p1, p2, JoinType.Inner, on...)
}

// LeftJoin ...
func (p1 Panel) LeftJoin(p2 Panel, on ...string) Panel {
	return Join(p1, p2, JoinType.Left, on...)
}

// RightJoin ...
func (p1 Panel) RightJoin(p2 Panel, on ...string) Panel {
	return Join(p1, p2, JoinType.Right, on...)
}

// LeftOnlyJoin ...
func (p1 Panel) LeftOnlyJoin(p2 Panel, on ...string) Panel {
	return Join(p1, p2, JoinType.LeftOnly, on...)
}

// RightOnlyJoin ...
func (p1 Panel) RightOnlyJoin(p2 Panel, on ...string) Panel {
	return Join(p1, p2, JoinType.RightOnly, on...)
}

// FullJoin ...
func (p1 Panel) FullJoin(p2 Panel, on ...string) Panel {
	return Join(p1, p2, JoinType.Full, on...)
}

// CrossJoin returns a Panel where each row of 'p1' is
// joined with each row of 'p2'.  Essentially, a full join
// without conditions.
func (p1 Panel) CrossJoin(p2 Panel, on ...string) Panel {
	return Join(p1, p2, JoinType.Cross, on...)
}

// Concat combines two Panels like length-wise
// (e.g. union in SQL)
func (p1 Panel) Concat(ps ...Panel) Panel {
	return Concat(p1, ps...)
}

// Concat combines two Panels like length-wise
// (e.g. union in SQL)
func Concat(p1 Panel, ps ...Panel) Panel {
	for _, p := range ps {
		for col := range p1 {
			p1[col] = append(p1[col], p[col]...)
		}
	}
	return p1
	// TODO
	// - For columns not in p2 that are p1,
	// 		fill with null values
	// - What happens if columns conflict?
}

// Union is a wrapper for 'Concat' that
// returns Panel with unique rows
func Union(p1 Panel, ps ...Panel) Panel {
	return p1.Concat(ps...).Unique()
}

// UnionAll is a wrapper for 'Concat' that
// returns Panel; does not check remove
// duplicative rows
func UnionAll(p1 Panel, ps ...Panel) Panel {
	return p1.Concat(ps...)
}

// Union is a wrapper for 'Concat' that
// returns Panel with unique rows
func (p1 Panel) Union(ps ...Panel) Panel {
	return p1.Concat(ps...).Unique()
}

// UnionAll is a wrapper for 'Concat' that
// returns Panel; does not check remove
// duplicative rows
func (p1 Panel) UnionAll(ps ...Panel) Panel {
	return UnionAll(p1, ps...)
}
