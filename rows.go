package panel

// Rows ...
func (p Panel) Rows(points ...int) Panel {
	return Rows(p, points...)
}

// Rows ...
func Rows(p Panel, points ...int) Panel {
	var start, end int
	tempPanel := Panel{}
	l := len(points)

	switch {
	case l == 1:
		start = points[0]
		end = -1
	case l > 1:
		start = points[0]
		end = points[1]
	default:
		start = 0
		end = -1
	}

	// check if end is greater than panel length
	if ln := p.Size().Length; end >= ln {
		// TODO: check that ln does not cutoff last element
		end = ln
	}

	for header := range p {
		c := []interface{}{}
		if end == -1 {
			for row, val := range p[header] {
				if row >= start {
					// c[row] = val
					c = append(c, val)
				}
			}
			tempPanel.Add(header, c)

		} else {
			for row, val := range p[header] {
				if row >= start && row <= end {
					// c[row] = val
					c = append(c, val)
				}
			}
			tempPanel.Add(header, c)

		}
	}
<<<<<<< HEAD
	return tempPanel
}

// Head returns a frame with the first 'n' number of rows of the original frame.
func (p Panel) Head(n ...int) Panel {
	return Head(p, n...)
}

// Head returns a frame with the first 'n' number of rows of the original frame.
func Head(p Panel, n ...int) Panel {
	if len(n) == 0 {
		return p.Rows(0, 4)
	}
	return p.Rows(0, n[0]-1)
}

// Tail returns a frame with the last 'n' number of rows of the original frame.
func (p Panel) Tail(n ...int) Panel {
	return Tail(p, n...)
}

// Tail returns a frame with the last 'n' number of rows of the original frame.
func Tail(p Panel, n ...int) Panel {
	ln := p.Size().Length
	if len(n) == 0 {
		return p.Rows(ln-5, ln)
	}
	return p.Rows(ln-n[0], ln)
}

// Subset returns data from certain rows
func (p Panel) Subset(n ...int) Panel {
	return Subset(p, n...)
}

// Subset returns data from certain rows
func Subset(p Panel, n ...int) Panel {
	tempPanel := New(nil)
	ln := p.Size().Length
	for row := range n {
		for col := range p {
			if n[row] < ln {
				tempPanel[col] = append(tempPanel[col], p[col][n[row]])
			}
		}
	}
	return tempPanel
}

// Sample returns a random panel with the length of 'n'
func (p Panel) Sample(n int) Panel {
	return Sample(p, n)
}

// Sample returns a random panel with the length of 'n'
func Sample(p Panel, n int) Panel {
	tempPanel := New(nil)
	ln := p.Size().Length
	switch {
	case n == 0:
		return New(nil)
	case n >= ln:
		return p
	case n > 0:
		sliceInt := createRandSlice(n, ln)
		for row := range sliceInt {
			for col := range p {
				tempPanel[col] = append(tempPanel[col], p[col][sliceInt[row]])
			}
		}
		return tempPanel
	default:
		// all else fails, return empty
		return tempPanel
	}

=======
	return tempPanel
}

// Head ...
func (p Panel) Head(i ...int) Panel {
	return Head(p, i...)
}

// Head ...
func Head(p Panel, i ...int) Panel {
	if len(i) == 0 {
		return p.Rows(0, 4)
	}
	return p.Rows(0, i[0]-1)
}

// Tail ...
func (p Panel) Tail(i ...int) Panel {
	return Tail(p, i...)
}

// Tail ...
func Tail(p Panel, i ...int) Panel {
	ln := p.Size().Length
	if len(i) == 0 {
		return p.Rows(ln-5, ln)
	}
	return p.Rows(ln-i[0], ln)
>>>>>>> b7d41b383db941fc86e7c093029e2c5b2be6c9e8
}
