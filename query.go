package panel

// Connect and query databases for data

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	// Import to allow for dynamic use of databases
	_ "github.com/alexbrainman/odbc"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/mattn/go-sqlite3"
)

// Config ...
type Config struct {
	Password string
	Port     int
	Server   string
	User     string
	Database string
}

// Connect ...
func Connect(driver string, config interface{}) *sql.DB {
	var connString string

	switch t := config.(type) {
	case Config:
		connString = fmt.Sprintf(
			"server=%s;user id=%s;password=%s;port=%d;database=%s",
			t.Server,
			t.User,
			t.Password,
			t.Port,
			t.Database,
		)
	case map[string]string:
		for k, v := range t {
			connString += fmt.Sprintf("%s=%v;", k, v)
		}
	case map[string]interface{}:
		for k, v := range t {
			connString += fmt.Sprintf("%s=%v;", k, v)
		}
	case string:
		if t[len(t)-5:] == ".json" {

			b, err := ioutil.ReadFile(t)
			if err != nil {
				fmt.Println(err.Error())
			}

			var config map[string]interface{}
			if err = json.Unmarshal(b, &config); err != nil {
				panic(err)
			}

			for k, v := range config {
				connString += fmt.Sprintf("%s=%v;", k, v)
			}

			// connString = fmt.Sprintf(
			// 	"server=%s;user id=%s;password=%s;port=%d;database=%s",
			// 	config["server"],
			// 	config["user"],
			// 	config["password"],
			// 	int(config["port"].(float64)),
			// 	config["database"],
			// )
		} else {
			connString = t
		}
	}

	conn, err := sql.Open(driver, connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	return conn
}

// Query ...
func Query(conn *sql.DB, code string) Panel {
	tempPanel := make(Panel)

	for _, query := range strings.Split(code, ";") {
		query = strings.TrimSpace(query)
		if query != "" {
			rows, err := conn.Query(query)
			if err != nil {
				// fmt.Printf("%s: \"%s\"\n", err.Error(), query)
				continue
			}
			defer rows.Close()

			columns, _ := rows.Columns()
			count := len(columns)
			values := make([]interface{}, count)
			valuePtrs := make([]interface{}, count)

			for rows.Next() {
				for i := range columns {
					valuePtrs[i] = &values[i]
				}
				rows.Scan(valuePtrs...)
				for i, col := range columns {
					var v interface{}
					val := values[i]
					b, ok := val.([]byte)

					if ok {
						v = string(b)
					} else {
						v = val
					}
					tempPanel[col] = append(tempPanel[col], v)
				}
			}

		}
	}
	return tempPanel
}
