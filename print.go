package panel

// May modify with delimit options: func (p Panel) String(delim string) string
// For example, p.String(",")... How would you pass delim in fmt.Println(p)???
//

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

// JSONify is just a pretty print wrapper
// for a special use case
func Pretty(d map[string]map[string]interface{}) string {
	// var i interface{} = d
	// m := i.(map[string]map[string]interface{})
	j, _ := json.MarshalIndent(d, "", "	")
	// fmt.Println(j)
	return string(j)
}

// Print .... (placeholder for pretty printing; undone)
func (p Panel) Print() {
	// Pretty way????
	fmt.Println(p)
}

// String returns printable text
func (p Panel) String() string {
	length := p.Size().Length

	var (
		res  string
		cols []string
	)

	// collect headers
	for col := range p {
		cols = append(cols, col)
	}

	// sort headers
	sort.Strings(cols)

	// add headers to output
	res = fmt.Sprintf("%s\n", strings.Join(cols, "\t"))
	// for each row
	for i := 0; i < length; i++ {
		var row []string
		for _, col := range cols {
			row = append(row, fmt.Sprintf("%v", p[col][i]))
		}
		res += fmt.Sprintf("%s\n", strings.Join(row, "\t"))
	}

	return res
}

// HTML returns markup table
func (p Panel) HTML() string {
	length := p.Size().Length

	var (
		res  string
		cols []string
	)

	// collect headers
	for col := range p {
		cols = append(cols, col)
	}

	// sort headers
	sort.Strings(cols)

	// add headers to output
	res = "<table><thead><tr>"
	for _, col := range cols {
		res += fmt.Sprintf("<th>%v</th>", col)
	}
	res += "</tr></thead><tbody>"

	// for each row
	for row := 0; row < length; row++ {
		s := "<tr>"
		for _, col := range cols {
			if col != "" {
				s += fmt.Sprintf("<td contenteditable=\"true\">%v</td>", p[col][row])
			} else {
				s += fmt.Sprintf("<td><strong>%v</strong></td>", p[col][row])
			}
		}
		res += fmt.Sprintf("%s</tr>", s)
	}

	return res + `</tbody></table>
	<style>
		table {
			font-family: arial, sans-serif;
			border-collapse: collapse;
			width: 100%;
		}
		
		td, th {
			border: 1px solid #dddddd;
			text-align: center;
			padding: 8px;
		}
		
		tr:nth-child(even) {
			background-color: #dddddd;
		}
	</style>
	<script>
		var updateButton = document.getElementById("update")
	</script>
	
	
	`
}
