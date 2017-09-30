package panel

<<<<<<< HEAD
import "fmt"

func main() {

	{
		df := New(map[string][]string{
			"Q": []string{"l", "o", "r", "q"},
			"N": []string{"5", "2", "9", "7"},
			"M": []string{"false", "true", "true", "false"},
		})

		df.Dtypes()
		fmt.Println(df.Sum("N"))

		port := 8844
		df.Serve(port, "test1") // localhost:8844

	}

	// 	{
	// 		df := Read("./data/gota.csv", true)

	// 		subset := df.
	// 			Filter("a", func(a string) bool {
	// 				return a == "a"
	// 			}).
	// 			Filter("b", func(b float64) bool {
	// 				return b >= 4.0

	// 			}).
	// 			Filter("d", func(d bool) bool {
	// 				return d == true
	// 			})

	// 			// df.PrintRowTypes()
	// 		fmt.Println(subset)

	// 	}

	// 	{ // Read a CSV file with headers via URL
	// 		// link := "https://vincentarelbundock.github.io/Rdatasets/csv/datasets/infert.csv"
	// 		link := "https://vincentarelbundock.github.io/Rdatasets/csv/Stat2Data/SATGPA.csv"
	// 		// link := "http://www.thestranger.com/seattle/Rss.xml?section=529161"

	// 		df := Read(link, true)
	// 		fmt.Println(df["gpa"])
	// 	}

	// 	// { // Read a CSV file with headers stored locally
	// 	// 	df := Read("data/iris.csv", true)
	// 	// 	df = df.Map("species", func(s string) string {
	// 	// 		return fmt.Sprintf("__%s", s)
	// 	// 	})
	// 	// 	fmt.Println(df["species"][:5])
	// 	// 	df.Serve(port, "iris")
	// 	// }

	// 	{ // Read a XLSX file with headers stored locally
	// 		df := Read("data/iris.xlsx", true)
	// 		df = df.Map("species", func(s string) string {
	// 			return fmt.Sprintf("**%s", s)
	// 		})

	// 		fmt.Println(df["species"][:5])

	// 		df = df.Rename(map[string]string{
	// 			"species":     "kind",
	// 			"petal.width": "pwidth",
	// 			"unnamed: 0":  "",
	// 		})

	// 		df.Serve(port, "iris")
	// 	}

	// 	// { // Try reading file that doesn't exist
	// 	// 	df := Read("hello.json", false)
	// 	// 	if df == nil {
	// 	// 		fmt.Println("nothing there")
	// 	// 	}
	// 	// }

	// }

	/////////////////////////////////////////////////////////////////

	// port := 8844

	// {
	// 	df := panel.New(
	// 		[][]string{
	// 			[]string{"A", "B", "C", "D"},
	// 			[]string{"a", "4", "5.1", "true"},
	// 			[]string{"k", "5", "7.0", "true"},
	// 			[]string{"k", "4", "6.0", "true"},
	// 			[]string{"a", "2", "7.1", "false"},
	// 		})

	// 	fmt.Println(df["A"])
	// 	fmt.Println("---------------")
	// 	df.GroupBy("A").Squeeze("B", "|")
	// 	// fmt.Println(panel.Meta)
	// }

	// {
	// 	df := panel.New(

	// 		[]map[string]interface{}{
	// 			map[string]interface{}{
	// 				"A": "a",
	// 				"B": 1,
	// 				"C": true,
	// 				"D": 0,
	// 			},
	// 			map[string]interface{}{
	// 				"A": "b",
	// 				"B": 2,
	// 				"C": true,
	// 				"D": 0.5,
	// 			},
	// 		})

	// 	fmt.Println(df)
	// }

	// {
	// 	df := panel.New(map[string][]string{
	// 		"Q": []string{"l", "o", "r", "q"},
	// 		"N": []string{"5", "2", "9", "7"},
	// 		"M": []string{"false", "true", "true", "false"},
	// 	})

	// 	fmt.Println("---------------")

	// 	// 	// 	// df.Dtypes()
	// 	// 	// 	// fmt.Println(df.Sum("N"))

	// 	fmt.Println(df.Subset(1, 3, 10))
	// 	fmt.Println("---------------")

	// 	// df = df.Iter(func(row panel.Row) panel.Row {
	// 	// 	row["X"] = fmt.Sprintf("%v.%v", row["Q"], row["N"])
	// 	// 	return row
	// 	// })

	// 	// df = df.Iter(func(row map[string]interface{}) map[string]interface{} {
	// 	// 	row["X"] = fmt.Sprintf("%v..%v", row["Q"], row["N"])
	// 	// 	return row
	// 	// })

	// 	// df = df.Iter(func(row panel.Row) map[string]interface{} {
	// 	// 	row["X"] = fmt.Sprintf("%v...%v", row["Q"], row["N"])
	// 	// 	return row
	// 	// })

	// 	df = df.Iter(func(row map[string]interface{}) panel.Row {
	// 		row["X"] = fmt.Sprintf("%v....%v", row["Q"], row["N"])
	// 		return row
	// 	})

	// 	fmt.Println(df)

	// 	df.Unload("panel_test.txt")

	// 	// df.Serve(port+1, "test1")
	// 	// fmt.Println(panel.Meta)
	// }

	// {
	// 	df1 := panel.Read("/Users/brydavis/go/src/github.com/adhox/panel/data/gota.csv", true)
	// 	df2 := panel.Read("/Users/brydavis/go/src/github.com/adhox/panel/data/gota3.csv", true)

	// 	// 	// 	// 	subset := df.
	// 	// 	// 	// 		Filter("a", func(a string) bool {
	// 	// 	// 	// 			return a == "a"
	// 	// 	// 	// 		}).
	// 	// 	// 	// 		Filter("b", func(b float64) bool {
	// 	// 	// 	// 			return b >= 4.0

	// 	// 	// 	// 		}).
	// 	// 	// 	// 		Filter("d", func(d bool) bool {
	// 	// 	// 	// 			return d == true
	// 	// 	// 	// 		})

	// 	// 	// 	// 		// df.PrintRowTypes()
	// 	// 	// 	// 	fmt.Println(subset)

	// 	// 	// 	// df.Dtypes()

	// 	merged := df1.InnerJoin(df2, "a")
	// 	// 	// 	// merged := df1.LeftJoin(df2, "a")
	// 	// 	// 	// merged := df1.RightJoin(df2, "a")
	// 	// 	// 	merged := df1.FullJoin(df2, "a")
	// 	// 	// 	// fmt.Println(merged.Select("a", "c", "f"))
	// 	// 	// 	groupBy := merged.GroupBy("a", "c", "f")
	// 	// 	// 	fmt.Printf("%T\n", groupBy)
	// 	// 	// 	fmt.Println(groupBy)

	// 	fmt.Println(merged)
	// 	fmt.Println("---------------")
	// 	// 	fmt.Println(df2)
	// 	// 	fmt.Println("---------------")
	// 	// 	fmt.Println(df1.FullJoin(df2, "a", "b"))
	// 	// 	fmt.Println("---------------")
	// 	// 	// fmt.Println(df1.Rows(0, 3))

	// 	df3 := panel.Read("/Users/brydavis/go/src/github.com/adhox/panel/data/iris.csv", true)

	// 	// 	for _, col := range df3.Columns() {
	// 	// 		fmt.Printf("~%s~\n", col)
	// 	// 	}

	// 	// 	// fmt.Println(df3.Head(7))
	// 	// 	// fmt.Println("---------------")
	// 	// 	// fmt.Println(df3.Tail(7))

	// 	// 	// df3.Dtypes()

	// 	// 	// fmt.Println("---------------")
	// 	// 	// fmt.Println(df3.Select("*"))
	// 	// 	// fmt.Println(df3.Select("petal.length", "petal.width", "sepal.length", "sepal.width", "species"))
	// 	// 	// fmt.Println(df3.Select())
	// 	// 	// fmt.Println(df3.Select(3, "sepal.length", 2, true).Head())

	// 	// 	// fmt.Println("---------------")
	// 	// 	// fmt.Println(df3.Select("petal.length", "sepal.length", "species").Sample(9))

	// 	// 	fmt.Println("---------------")
	// 	// 	fmt.Println(df3.Rename("species", "kind").Head())

	// 	// 	fmt.Println("---------------")
	// 	// 	renameCols := map[string]string{"species": "kind", "sepal.length": "sepal_length"}
	// 	// 	df4 := df3.Rename(renameCols)
	// 	// 	fmt.Println("---------------")

	// 	// 	fmt.Println(df4.Size().Length)
	// 	// 	// fmt.Println(df4)

	// 	// 	// for _, col := range df4.Columns() {
	// 	// 	// 	fmt.Println(col, len(df4[col]))
	// 	// 	// }

	// 	// }

	// 	// { // Read a CSV file with headers via URL
	// 	// 	// link := "https://vincentarelbundock.github.io/Rdatasets/csv/datasets/infert.csv"
	// 	// 	link := "https://vincentarelbundock.github.io/Rdatasets/csv/Stat2Data/SATGPA.csv"
	// 	// 	// link := "http://www.thestranger.com/seattle/Rss.xml?section=529161"

	// 	// 	df := panel.Read(link, true)
	// 	// 	fmt.Println(df["gpa"])
	// 	// 	df = df.Convert("gpa", panel.String)
	// 	// 	fmt.Printf("%T\n", df["gpa"][0])

	// 	df3.GroupBy("species").Squeeze("petal.length", ", ")

	// }

	// { // Read a CSV file with headers stored locally
	// 	df := panel.Read("/Users/brydavis/go/src/github.com/adhox/panel/data/iris.csv", true)
	// 	// df = df.Map("species", func(s string) string {
	// 	// 	return fmt.Sprintf("__%s", s)
	// 	// })

	// 	// fmt.Println(df["species"][:5])
	// 	fmt.Printf("%T\n", df.GroupBy("species"))
	// 	// df.Serve(port, "iris")
	// }

	// {
	// 	// Read a CSV file with headers stored locally
	// 	// df := panel.Read("/Users/brydavis/code/mock_car_sales.csv", true)

	// 	// fmt.Println(df["car_make"][:5])
	// 	// fmt.Printf("%T\n", df.GroupBy("car_make", "car_color", "buyer_gender"))

	// 	// df.GroupBy("car_make", "car_color", "buyer_gender").Count("price_k")
	// 	// df.GroupBy("car_make", "car_color", "buyer_gender").Sum("price_k")
	// 	// df.GroupBy("car_make", "car_color", "buyer_gender").Min("price_k")
	// 	// df1 := df.GroupBy("car_make", "car_color", "buyer_gender").Sum("price_k")
	// 	// df2 := df.GroupBy("car_make", "car_color", "buyer_gender").Count("price_k")
	// 	// df = df.GroupBy("car_make", "car_color", "buyer_gender").Mean("price_k")

	// 	// df1.Write("panel_agg_export_test.csv")

	// 	// fmt.Println(df)
	// }

	// { // adhox-style program
	// 	query := `select distinct Id, FirstName, LastName from Persons where FirstName like 'A%'`
	// 	conn := panel.Connect("mssql", "config.json") // Connecting to an Azure DB for example
	// 	p := conn.Query(conn, query)                  // what about multiple return datasets???
	// 	p.Write("/some/path/export.xlsx")
	// }

	// { // Read a XLSX file with headers stored locally
	// 	df := Read("data/iris.xlsx", true)
	// 	df = df.Map("species", func(s string) string {
	// 		return fmt.Sprintf("**%s", s)
	// 	})

	// 	fmt.Println(df["species"][:5])

	// 	df = df.Rename(map[string]string{
	// 		"species":     "kind",
	// 		"petal.width": "pwidth",
	// 		"unnamed: 0":  "",
	// 	})

	// 	df.Serve(port, "iris")
	// }

	// { // Try reading file that doesn't exist
	// 	df := Read("hello.json", false)
	// 	if df == nil {
	// 		fmt.Println("nothing there")
	// 	}
}
=======
import (
	"fmt"
)

func ExamplePanel() {
	h := "Hello"
	if h != "Hi" {
		fmt.Println("Wanted Hi not %s", h)
	}
}

// print string
//
//
// func TestHello(t *testing.T) {
// 	h := "Hello"
// 	if h != "Hi" {
// 		t.Errorf("Wanted Hi not %s", h)
// 	}
// }

// // TODO
// // - Transform Go structures into frames

// func main() {
// 	port := 8844

// 	{
// 		df := New(map[string][]string{
// 			"Q": []string{"l", "o", "r", "q"},
// 			"N": []string{"5", "2", "9", "7"},
// 			"M": []string{"false", "true", "true", "false"},
// 		})

// 		df.Dtypes()
// 		fmt.Println(df.Sum("N"))

// 		go df.Serve(port+1, "test1")

// 	}

// 	{
// 		df := Read("./data/gota.csv", true)

// 		subset := df.
// 			Filter("a", func(a string) bool {
// 				return a == "a"
// 			}).
// 			Filter("b", func(b float64) bool {
// 				return b >= 4.0

// 			}).
// 			Filter("d", func(d bool) bool {
// 				return d == true
// 			})

// 			// df.PrintRowTypes()
// 		fmt.Println(subset)

// 	}

// 	{ // Read a CSV file with headers via URL
// 		// link := "https://vincentarelbundock.github.io/Rdatasets/csv/datasets/infert.csv"
// 		link := "https://vincentarelbundock.github.io/Rdatasets/csv/Stat2Data/SATGPA.csv"
// 		// link := "http://www.thestranger.com/seattle/Rss.xml?section=529161"

// 		df := Read(link, true)
// 		fmt.Println(df["gpa"])
// 	}

// 	// { // Read a CSV file with headers stored locally
// 	// 	df := Read("data/iris.csv", true)
// 	// 	df = df.Map("species", func(s string) string {
// 	// 		return fmt.Sprintf("__%s", s)
// 	// 	})
// 	// 	fmt.Println(df["species"][:5])
// 	// 	df.Serve(port, "iris")
// 	// }

// 	{ // Read a XLSX file with headers stored locally
// 		df := Read("data/iris.xlsx", true)
// 		df = df.Map("species", func(s string) string {
// 			return fmt.Sprintf("**%s", s)
// 		})

// 		fmt.Println(df["species"][:5])

// 		df = df.Rename(map[string]string{
// 			"species":     "kind",
// 			"petal.width": "pwidth",
// 			"unnamed: 0":  "",
// 		})

// 		df.Serve(port, "iris")
// 	}

// 	// { // Try reading file that doesn't exist
// 	// 	df := Read("hello.json", false)
// 	// 	if df == nil {
// 	// 		fmt.Println("nothing there")
// 	// 	}
// 	// }

// }

/////////////////////////////////////////////////////////////////

// port := 8844

// {
// 	df := panel.New(
// 		[][]string{
// 			[]string{"A", "B", "C", "D"},
// 			[]string{"a", "4", "5.1", "true"},
// 			[]string{"k", "5", "7.0", "true"},
// 			[]string{"k", "4", "6.0", "true"},
// 			[]string{"a", "2", "7.1", "false"},
// 		})

// 	fmt.Println(df["A"])
// 	fmt.Println("---------------")
// 	df.GroupBy("A").Squeeze("B", "|")
// 	// fmt.Println(panel.Meta)
// }

// {
// 	df := panel.New(

// 		[]map[string]interface{}{
// 			map[string]interface{}{
// 				"A": "a",
// 				"B": 1,
// 				"C": true,
// 				"D": 0,
// 			},
// 			map[string]interface{}{
// 				"A": "b",
// 				"B": 2,
// 				"C": true,
// 				"D": 0.5,
// 			},
// 		})

// 	fmt.Println(df)
// }

// {
// 	df := panel.New(map[string][]string{
// 		"Q": []string{"l", "o", "r", "q"},
// 		"N": []string{"5", "2", "9", "7"},
// 		"M": []string{"false", "true", "true", "false"},
// 	})

// 	fmt.Println("---------------")

// 	// 	// 	// df.Dtypes()
// 	// 	// 	// fmt.Println(df.Sum("N"))

// 	fmt.Println(df.Subset(1, 3, 10))
// 	fmt.Println("---------------")

// 	// df = df.Iter(func(row panel.Row) panel.Row {
// 	// 	row["X"] = fmt.Sprintf("%v.%v", row["Q"], row["N"])
// 	// 	return row
// 	// })

// 	// df = df.Iter(func(row map[string]interface{}) map[string]interface{} {
// 	// 	row["X"] = fmt.Sprintf("%v..%v", row["Q"], row["N"])
// 	// 	return row
// 	// })

// 	// df = df.Iter(func(row panel.Row) map[string]interface{} {
// 	// 	row["X"] = fmt.Sprintf("%v...%v", row["Q"], row["N"])
// 	// 	return row
// 	// })

// 	df = df.Iter(func(row map[string]interface{}) panel.Row {
// 		row["X"] = fmt.Sprintf("%v....%v", row["Q"], row["N"])
// 		return row
// 	})

// 	fmt.Println(df)

// 	df.Unload("panel_test.txt")

// 	// df.Serve(port+1, "test1")
// 	// fmt.Println(panel.Meta)
// }

// {
// 	df1 := panel.Read("/Users/brydavis/go/src/github.com/adhox/panel/data/gota.csv", true)
// 	df2 := panel.Read("/Users/brydavis/go/src/github.com/adhox/panel/data/gota3.csv", true)

// 	// 	// 	// 	subset := df.
// 	// 	// 	// 		Filter("a", func(a string) bool {
// 	// 	// 	// 			return a == "a"
// 	// 	// 	// 		}).
// 	// 	// 	// 		Filter("b", func(b float64) bool {
// 	// 	// 	// 			return b >= 4.0

// 	// 	// 	// 		}).
// 	// 	// 	// 		Filter("d", func(d bool) bool {
// 	// 	// 	// 			return d == true
// 	// 	// 	// 		})

// 	// 	// 	// 		// df.PrintRowTypes()
// 	// 	// 	// 	fmt.Println(subset)

// 	// 	// 	// df.Dtypes()

// 	merged := df1.InnerJoin(df2, "a")
// 	// 	// 	// merged := df1.LeftJoin(df2, "a")
// 	// 	// 	// merged := df1.RightJoin(df2, "a")
// 	// 	// 	merged := df1.FullJoin(df2, "a")
// 	// 	// 	// fmt.Println(merged.Select("a", "c", "f"))
// 	// 	// 	groupBy := merged.GroupBy("a", "c", "f")
// 	// 	// 	fmt.Printf("%T\n", groupBy)
// 	// 	// 	fmt.Println(groupBy)

// 	fmt.Println(merged)
// 	fmt.Println("---------------")
// 	// 	fmt.Println(df2)
// 	// 	fmt.Println("---------------")
// 	// 	fmt.Println(df1.FullJoin(df2, "a", "b"))
// 	// 	fmt.Println("---------------")
// 	// 	// fmt.Println(df1.Rows(0, 3))

// 	df3 := panel.Read("/Users/brydavis/go/src/github.com/adhox/panel/data/iris.csv", true)

// 	// 	for _, col := range df3.Columns() {
// 	// 		fmt.Printf("~%s~\n", col)
// 	// 	}

// 	// 	// fmt.Println(df3.Head(7))
// 	// 	// fmt.Println("---------------")
// 	// 	// fmt.Println(df3.Tail(7))

// 	// 	// df3.Dtypes()

// 	// 	// fmt.Println("---------------")
// 	// 	// fmt.Println(df3.Select("*"))
// 	// 	// fmt.Println(df3.Select("petal.length", "petal.width", "sepal.length", "sepal.width", "species"))
// 	// 	// fmt.Println(df3.Select())
// 	// 	// fmt.Println(df3.Select(3, "sepal.length", 2, true).Head())

// 	// 	// fmt.Println("---------------")
// 	// 	// fmt.Println(df3.Select("petal.length", "sepal.length", "species").Sample(9))

// 	// 	fmt.Println("---------------")
// 	// 	fmt.Println(df3.Rename("species", "kind").Head())

// 	// 	fmt.Println("---------------")
// 	// 	renameCols := map[string]string{"species": "kind", "sepal.length": "sepal_length"}
// 	// 	df4 := df3.Rename(renameCols)
// 	// 	fmt.Println("---------------")

// 	// 	fmt.Println(df4.Size().Length)
// 	// 	// fmt.Println(df4)

// 	// 	// for _, col := range df4.Columns() {
// 	// 	// 	fmt.Println(col, len(df4[col]))
// 	// 	// }

// 	// }

// 	// { // Read a CSV file with headers via URL
// 	// 	// link := "https://vincentarelbundock.github.io/Rdatasets/csv/datasets/infert.csv"
// 	// 	link := "https://vincentarelbundock.github.io/Rdatasets/csv/Stat2Data/SATGPA.csv"
// 	// 	// link := "http://www.thestranger.com/seattle/Rss.xml?section=529161"

// 	// 	df := panel.Read(link, true)
// 	// 	fmt.Println(df["gpa"])
// 	// 	df = df.Convert("gpa", panel.String)
// 	// 	fmt.Printf("%T\n", df["gpa"][0])

// 	df3.GroupBy("species").Squeeze("petal.length", ", ")

// }

// { // Read a CSV file with headers stored locally
// 	df := panel.Read("/Users/brydavis/go/src/github.com/adhox/panel/data/iris.csv", true)
// 	// df = df.Map("species", func(s string) string {
// 	// 	return fmt.Sprintf("__%s", s)
// 	// })

// 	// fmt.Println(df["species"][:5])
// 	fmt.Printf("%T\n", df.GroupBy("species"))
// 	// df.Serve(port, "iris")
// }

// {
// 	// Read a CSV file with headers stored locally
// 	// df := panel.Read("/Users/brydavis/code/mock_car_sales.csv", true)

// 	// fmt.Println(df["car_make"][:5])
// 	// fmt.Printf("%T\n", df.GroupBy("car_make", "car_color", "buyer_gender"))

// 	// df.GroupBy("car_make", "car_color", "buyer_gender").Count("price_k")
// 	// df.GroupBy("car_make", "car_color", "buyer_gender").Sum("price_k")
// 	// df.GroupBy("car_make", "car_color", "buyer_gender").Min("price_k")
// 	// df1 := df.GroupBy("car_make", "car_color", "buyer_gender").Sum("price_k")
// 	// df2 := df.GroupBy("car_make", "car_color", "buyer_gender").Count("price_k")
// 	// df = df.GroupBy("car_make", "car_color", "buyer_gender").Mean("price_k")

// 	// df1.Write("panel_agg_export_test.csv")

// 	// fmt.Println(df)
// }

// { // adhox-style program
// 	query := `select distinct Id, FirstName, LastName from Persons where FirstName like 'A%'`
// 	conn := panel.Connect("mssql", "config.json") // Connecting to an Azure DB for example
// 	p := conn.Query(conn, query)                  // what about multiple return datasets???
// 	p.Write("/some/path/export.xlsx")
// }

// { // Read a XLSX file with headers stored locally
// 	df := Read("data/iris.xlsx", true)
// 	df = df.Map("species", func(s string) string {
// 		return fmt.Sprintf("**%s", s)
// 	})

// 	fmt.Println(df["species"][:5])

// 	df = df.Rename(map[string]string{
// 		"species":     "kind",
// 		"petal.width": "pwidth",
// 		"unnamed: 0":  "",
// 	})

// 	df.Serve(port, "iris")
// }

// { // Try reading file that doesn't exist
// 	df := Read("hello.json", false)
// 	if df == nil {
// 		fmt.Println("nothing there")
// 	}
// }
>>>>>>> b7d41b383db941fc86e7c093029e2c5b2be6c9e8
