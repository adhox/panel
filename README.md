# Panel


Data analysis framework for Go.

#### Examples

Use results from a SQL query. 
```go
package main

import (
	"fmt"
	. "github.com/brydavis/panel"
)

func main() {
	query := `select distinct Id, FirstName, LastName from Persons where FirstName like 'A%'`

	conn := Connect("mssql", "config.json") // Connecting to an Azure DB for example 
	persons := Query(conn, query)

	fmt.Printf(p.Select("FirstName").Unique())

	McNames := persons.Map("LastName", func(s string) string {
		return fmt.Sprintf("Mc%s", s)
	})

	fmt.Println(McNames.Select("FirstName", "LastName"))
}
```
Upload a CSV and map a function
```go
	{ // Read a CSV
		df := Read("data/iris.csv", true)
		df = df.Map("species", func(s string) string {
			return fmt.Sprintf("__%s", s)
		})

		fmt.Println(df["species"])

	}
```






#### Backlog
[X]	df.Select using (int...) OR (string...) OR (interface{}...) 
[X]	df.Subset(0,3,9) returns the first, fourth, and tenth rows
[X]	Row-wise operations (think Map / Filter, but accross rows) -> Iter
[X]	df.Rename -> check if 2 strings OR if map[string]string
[X]	df.Sample(x)
[X]	LeftOuterJoin vs LeftInnerJoin vs LeftOnlyJoin? https://stackoverflow.com/questions/406294/left-join-vs-left-outer-join-in-sql-server


[ ]	Adapt `adhox.go` to connect and query 
[ ]	Documentation (preferably GoDoc standard) 
[ ]	df.Arrange(panel.OrderBy("A","B").Descend(), panel.OrderBy("C").Ascend())
[ ]	Standardize design patterns (i.e. use similar variable names for similar actions)
[ ]	add error messaging
[ ]	CrossJoin (Cartesian, no clause)
[ ]	Ensure as much chaining ability as possible
[ ]	Pivot and CrossTabs... GroupBy+?
[ ]	Supported data types
		[X]	Int
		[X]	Float
		[X]	String
		[X]	Bool
		[X]	Time
		[ ]	Others???
[ ]	Modify with MontanaFlynn's stats package
[ ]	Update SQL interface to be smoother; allow for inline connection string panel.Connect("serverX",4444, "usernom", "passw")
[ ]	df.Query -> use SQLite backend to enable querying the panel
[ ]	panel.Query -> return multiple panels if multiple returns from SQL

[ ]	Align New and Load as mutual gateways  
[ ]	Import file format 
		[X]	CSV
		[ ]	TSV (default)
		[X]	XLSX
		[ ]	JSON
		[ ]	XML
		[ ]	SQL 
		[ ]	type gonum.Matrix
		[X]	type [][]string
		[ ]	multiple types ([][]string{...}, map[string]string{...})
		[ ]	type strings.NewReader("... here's some CSV...")
[ ]	Export file format 
		[X]	CSV
		[X]	TSV (default)
		[X]	XLSX
		[X]	JSON
		[ ]	XML
		[ ]	SQL 
[ ]	special export of XLSX (with multiple tabs; naming ability)		
[ ]	create metadata manager: track things like order of columns, etc.
[ ]	TEST, TEST, TEST!!! (look at 'gota' package)
[ ]	handle when uploaded column names are empty strings ... "_"?
[ ]	fix Concat to check for conflicting / misaligned columns
[ ]	Implement STATS
		[ ]	Mean float64
		[ ]	Median float64
		[ ]	Percentile / Quantile float64
		


<br>
<br>

<hr>
<!-- &copy; 2015-2016  -->
<!-- [MIT License](https://www.github.com/openwonk/mit_license) -->

[I'm an inline-style link](https://www.github.com/openwonk/mit_license)
