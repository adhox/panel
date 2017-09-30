package panel


import "net/http"
import "strconv"
import "fmt"

// Serve serves data in the
// Panel to a simple webapp
func (p Panel) Serve(port int, sheet string) {
	fmt.Printf("Starting PANEL Server @ %d\n", port)
	http.HandleFunc("/"+sheet, func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, p.HTML())
	})
	if err := http.ListenAndServe(fmt.Sprintf(":%s", strconv.Itoa(port)), nil); err != nil {
		fmt.Println(err)
	}
}

// func handleSheet(p Panel) func(http.ResponseWriter, *http.Request) {
// 	return func(res http.ResponseWriter, req *http.Request) {
// 		fmt.Fprint(res, p.HTML())
// 	}
// }
