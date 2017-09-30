package panel

import (
	"math/rand"
	"time"
)

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

}

func createRandSlice(n, lim int) (list []int) {
	for len(list) < n {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		// x := rand.Intn(lim)
		x := r.Intn(lim)
		if !intInSlice(x, list) {
			list = append(list, x)
		}
	}
	return
}

func intInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
