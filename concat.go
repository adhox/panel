package panel

// Concat ...
func (pn Panel) Concat(p2 Panel) Panel {
	return Concat(pn, p2)
}

// Concat ...
func Concat(pn, p2 Panel) Panel {
	for col := range pn {
		pn[col] = append(pn[col], p2[col]...)
	}
	return pn
}

// simply stack two panels together
// mimics SQL:
//      select 1 as A, 2 as B
//      union all
//      select 3 as A, 4 as B

// TODO
// change Union and UnionAll signatures to something like...
// `func Union(ps ...Panel) Panel`
// where it loops through endless panels
// ... or maybe just implement `Reduce` func
// that could do the same, given a func

// Union ...
func Union(pn, p2 Panel) Panel {
	return pn.Concat(p2).Unique()
}

// UnionAll ...
func UnionAll(pn, p2 Panel) Panel {
	return pn.Concat(p2)
}

// Union ...
func (pn Panel) Union(p2 Panel) Panel {
	return Union(pn, p2)
}

// UnionAll ...
func (pn Panel) UnionAll(p2 Panel) Panel {
	return UnionAll(pn, p2)
}
