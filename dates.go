package panel

import (
	"fmt"
	"time"
)

var dateStandards = map[int]string{
	100: "Jan 02 2006 03:04AM",
	1:   "01/02/06",
	101: "01/02/2006",
	2:   "06.01.02",
	102: "2006.01.02",
	3:   "02/01/06",
	103: "02/01/2006",
	4:   "02.01.06",
	104: "02.01.2006",
	5:   "02-01-06",
	105: "02-01-2006",
	6:   "02 Jan 06",
	106: "02 Jan 2006",
	7:   "Jan 02, 06",
	107: "Jan 02, 2006",
	108: "15:04:05",
	109: "Jan 02 2006 03:04:05.000PM",
	10:  "06-01-02",
	110: "2006-01-02",
	11:  "06/01/02",
	111: "2006/01/02",
	12:  "060102",
	112: "20060102",
	113: "02 Jan 2006 15:04:05.000",
	114: "15:04:05.000",
	120: "2006-01-02 15:04:05",
	121: "2006-01-02 15:04:05.000",
	126: "2006-01-02T15:04:05.000",
	127: "2006-01-02T15:04:05.000Z",
	130: "02 Jan 2006 03:04:05.000PM",
	131: "02/01/06 03:04:05.000PM",
	// Subsequent are not really standards but should parse as dates
	201: "1/2/2006",
	202: "1-2-2006",
	203: "1/2/06",
	204: "1-2-06",
	205: "2006-1-2",
	206: "06-1-2",
	// 207: "2006-01-02",
}

func isDate(date string) (int, bool) {
	for code, format := range dateStandards {
		_, err := time.Parse(format, date)
		if err == nil {
			return code, true
		}
	}
	// 99 is a throwaway value
	// since 0 is a standard date format
	return 99, false
}

func dateFormat(code int) (string, error) {
	if val, ok := dateStandards[code]; ok {
		return val, nil
	}
	return "", fmt.Errorf("code (%d) is not recognized as standard format\n\r", code)
}
