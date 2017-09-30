package panel

import (
	"fmt"
	"time"
)

// Meta ...
type metadata []struct {
	Panel       *Panel
	Columns     []string
	DateCreated time.Time
}

// Meta ...
var Meta metadata

// Append ...
func (m metadata) Append(p *Panel) metadata {
	// fmt.Println(p)
	// fmt.Println("=========================")

	pn := *p
	d := struct {
		Panel       *Panel
		Columns     []string
		DateCreated time.Time
	}{
		p,
		pn.Columns(),
		time.Now(),
	}

	// fmt.Println(d)

	n := append(m, d)
	// m = &n
	// fmt.Println(m)

	return n

}

// String ...
func (m metadata) String() {
	for p := range m {
		fmt.Println(m[p].DateCreated)
		fmt.Println(m[p].Columns)
		fmt.Println(m[p].Panel)
	}
}
