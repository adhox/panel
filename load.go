package panel

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/tealeg/xlsx"
)

// TSV is a simple type wrapper
// to assist in the parsing logic
type TSV string

// CSV is a simple type wrapper
// to assist in the parsing logic
type CSV string

// XML is a simple type wrapper
// to assist in the parsing logic
type XML string

// JSON is a simple type wrapper
// to assist in the parsing logic
type JSON string

// SQL is a simple type wrapper
// to assist in the parsing logic
type SQL string

// Load ...
// 1. Loading data from file or memory
// 2. Perform data cleaning and standardization
func Load(fname string, head bool) Panel {
	p := make(Panel)

	// check if filename is URL
	if u, err := url.Parse(fname); err == nil && fname[:4] == "http" {
		fmt.Printf("downloading (%s)\n", u.String())
		// if a URL, download file
		// and reassign fname
		fname = download(*u)
	}

	fmt.Printf("reading file (%s)\n", fname)

	switch path.Ext(fname) {
	case ".csv":
		return readCSV(fname, head).Clean()
	case ".xml", ".html":
		fmt.Printf("reading %s", fname)
		b, _ := ioutil.ReadFile(fname)
		fmt.Println(string(b))
		return nil

	case ".tsv":
		return nil

	case ".json":
		return nil

	case ".xlsx":
		// fmt.Println("reading EXCEL")
		return readXLSX(fname, head).Clean()

	default: // text file; treat like Hadoop
		// file, err := xlsx.OpenFile(fname)
		// if err != nil {
		// 	fmt.Println(err)
		// }
	}

	return p.Clean()
}

// Read does same as Load
func Read(fname string, head bool) Panel {
	return Load(fname, head)
}

func download(u url.URL) string {

	res, err := http.Get(u.String())
	if err != nil {
		return ""
	}
	defer res.Body.Close()

	fname := strings.ToLower(fmt.Sprintf("%s.%s", u.Host, strings.Split(path.Base(u.String()), "?")[0]))

	file, _ := os.Create(fname)
	body, _ := ioutil.ReadAll(res.Body)
	file.Write(body)

	return fname
}

func readCSV(fname string, head bool) Panel {
	p := make(Panel)

	h := []string{}
	f, err := os.Open(fname)
	if err != nil {
		fmt.Println(err)
	}

	headless := false
	if !head {
		headless = true
	}

	r := csv.NewReader(bufio.NewReader(f))
	r.FieldsPerRecord = -1
	for {
		if head {
			line, _ := r.Read()
			for _, val := range line {
				h = append(h, strings.ToLower(val))
			}
			head = false
		} else {
			line, err := r.Read()
			if err == io.EOF {
				break
			}
			for key, val := range line {
				if headless {
					k := strconv.Itoa(key)
					p[k] = append(p[k], val)
				} else {
					p[h[key]] = append(p[h[key]], val)
				}
			}
		}
	}
	return p
}

func readXLSX(fname string, head bool) Panel {
	p := make(Panel)
	h := []string{}

	file, err := xlsx.OpenFile(fname)
	if err != nil {
		fmt.Println(err)
	}

	headless := false
	if !head {
		headless = true
	}

	for _, sheet := range file.Sheets {
		for _, row := range sheet.Rows {
			if head {
				for _, cell := range row.Cells {
					val, _ := cell.String()
					h = append(h, strings.ToLower(val))
				}
				head = false
			} else {
				for cn, cell := range row.Cells {
					val, _ := cell.String()
					if headless {
						k := strconv.Itoa(cn)
						p[k] = append(p[k], val)
					} else {
						p[h[cn]] = append(p[h[cn]], val)
					}
				}
			}
		}
	}
	return p
}

func (p Panel) Write(fname string) { // ADD ERROR
	Unload(p, fname)
}

// Dump is method short hand for Unload function
func (p Panel) Dump(fname string) { // ADD ERROR
	Unload(p, fname)
}

// Export is method short hand for Unload function
func (p Panel) Export(fname string) { // ADD ERROR
	Unload(p, fname)
}

// Unload is method short hand for Unload function
func (p Panel) Unload(fname string) { // ADD ERROR
	Unload(p, fname)
}

// Unload moves data from memory to file
func Unload(p Panel, fname string) { // ADD ERROR
	// p.Clean()
	switch path.Ext(fname) {
	case ".csv":
		// Convert to CSV-like structure
		width := p.Size().Width
		length := p.Size().Length
		headers := make(map[string]int)
		colnum := 0
		for head := range p {
			headers[head] = colnum
			colnum++
		}

		records := [][]string{}
		record := make([]string, width)
		for head := range p {
			record[headers[head]] = head
		}
		records = append(records, record)

		for i := 0; i < length; i++ {
			record := make([]string, width)
			for head, col := range p {
				var val string
				switch t := col[i].(type) {
				case time.Time:
					d110, _ := dateFormat(110)
					tt := t.Format(d110)
					fmt.Println(tt)
					val = tt
				default:
					val = fmt.Sprintf("%v", t)
				}
				record[headers[head]] = val
			}
			records = append(records, record)
		}

		// fmt.Println(records)
		// Write to file
		file := new(os.File)
		if _, err := os.Stat(fname); err == nil {
			file, _ = os.OpenFile(fname, os.O_RDWR|os.O_APPEND, 0777)
		} else {
			file, _ = os.Create(fname)
		}

		w := csv.NewWriter(file)
		// w.WriteAll(records)

		for _, record := range records {
			if err := w.Write(record); err != nil {
				log.Fatalln("error writing record to csv:", err)
			}
		}

		w.Flush() // Write any buffered data to the underlying writer
		if err := w.Error(); err != nil {
			log.Fatal(err)
		}

	case ".xml":

	case ".json":
		export := []map[string]interface{}{}

		// width := p.Size().Width
		length := p.Size().Length

		for i := 0; i < length; i++ {
			row := make(map[string]interface{})
			for header, series := range p {
				row[header] = series[i]
			}
			export = append(export, row)
		}

		j, _ := json.Marshal(export)

		file := new(os.File)
		if _, err := os.Stat(fname); err == nil {
			file, _ = os.OpenFile(fname, os.O_RDWR|os.O_APPEND, 0777)
		} else {
			file, _ = os.Create(fname)
		}

		_, err := file.Write(j)
		if err != nil {
			fmt.Println(err)
		}

		file.Close()

	case ".xlsx", ".xls":
		// Convert to CSV-like structure
		width := p.Size().Width
		length := p.Size().Length
		headers := make(map[string]int)
		colnum := 0
		for head := range p {
			headers[head] = colnum
			colnum++
		}

		records := [][]string{}
		record := make([]string, width)
		for head := range p {
			record[headers[head]] = head
		}
		records = append(records, record)

		for i := 0; i < length; i++ {
			record := make([]string, width)
			for head, col := range p {
				var val string
				switch t := col[i].(type) {
				case time.Time:
					d110, _ := dateFormat(110)
					tt := t.Format(d110)
					fmt.Println(tt)
					val = tt
				default:
					val = fmt.Sprintf("%v", t)
				}
				record[headers[head]] = val
			}
			records = append(records, record)
		}

		// fmt.Println(records)
		// Write to file
		// file := new(os.File)
		// if _, err := os.Stat(fname); err == nil {
		// 	file, _ = os.OpenFile(fname, os.O_RDWR|os.O_APPEND, 0777)
		// } else {
		// 	file, _ = os.Create(fname)
		// }

		var excel *xlsx.File
		var sheet *xlsx.Sheet
		var row *xlsx.Row
		var cell *xlsx.Cell
		var err error

		excel = xlsx.NewFile()
		sheet, err = excel.AddSheet("Sheet1")
		if err != nil {
			fmt.Printf(err.Error())
		}

		for _, rec := range records {
			row = sheet.AddRow()

			for _, field := range rec {
				cell = row.AddCell()
				cell.Value = fmt.Sprintf("%s", field)
			}
		}

		err = excel.Save(fname)
		if err != nil {
			fmt.Printf(err.Error())
			log.Fatalln("error writing record to xlsx:", err)

		}

	default: // TSV
		// file := new(os.File)
		if _, err := os.Stat(fname); err == nil {
			// file, _ = os.OpenFile(fname, os.O_RDWR|os.O_APPEND, 0777)
			// file, _ = os.OpenFile(fname, os.O_RDWR, 0777)
			if err := os.Remove(fname); err != nil {
				fmt.Println(err)
			}
		}
		file, _ := os.Create(fname)
		file.Write([]byte(fmt.Sprintf("%v", p)))
	}
}

// CopyFile provides a basic process that copies
// the contents of file 'a' into file 'b'
func CopyFile(a, b string, head bool) {
	Load(a, head).Unload(b)
}
