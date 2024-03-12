package internal

import (
	"bytes"
	"strconv"
	"strings"
)

func FlattenJSON(data []byte) []byte {
	return bytes.Replace(data, []byte("]["), []byte(","), -1)
}
func SSIMtoDate(s string) string {
	var monthMap = make(map[string]string)
	monthMap["JAN"] = "01"
	monthMap["FEB"] = "02"
	monthMap["MAR"] = "03"
	monthMap["APR"] = "04"
	monthMap["MAY"] = "05"
	monthMap["JUN"] = "06"
	monthMap["JUL"] = "07"
	monthMap["AUG"] = "08"
	monthMap["SEP"] = "09"
	monthMap["OCT"] = "10"
	monthMap["NOV"] = "11"
	monthMap["DEC"] = "12"

	var dateString string

	length := len(s)
	switch length {
	case 6:
		// 4JUL24 012345
		dateString += "0" + s[:1] + "-" + monthMap[s[1:4]] + "-20" + s[4:]
	case 7:
		// 19JUL24 0123456
		dateString += s[:2] + "-" + monthMap[s[2:5]] + "-20" + s[5:]
	}
	return dateString
}

// 845 - 840 (14) = 5
func NumberToTime(n int64) string {
	hours := n / 60
	minutes := n - hours*60
	hoursStr := strconv.FormatInt(hours, 10)
	minutesStr := strconv.FormatInt(minutes, 10)

	if hours < 10 {
		hoursStr = "0" + hoursStr
	}
	if minutes < 10 {
		minutesStr = "0" + minutesStr
	}

	return hoursStr + ":" + minutesStr
}
func DaysOfOperation(s string) string {
	return strings.ReplaceAll(s, " ", ".")
}
