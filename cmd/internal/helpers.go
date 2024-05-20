package internal

import (
	"bytes"
	"encoding/csv"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
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
func csvReadRecords(fileName string) [][]string {
	// Z,Do,Linia,Numer,Odlot,Przylot,Od,Do,Dni,Samolot,Operator,Typ
	// KRK,FRA,LH,1365,11:15,12:55,01-04-2024,01-04-2024,1......,320,LH,J
	// KRK,FRA,LH,1365,11:15,12:55,02-04-2024,05-04-2024,.2..5..,320,LH,J
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Unable to read file ", err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to read records ", err)
	}

	return records
}

func moreThanOneNumberReg(s string) bool {
	exp := `.*\d.*\d.*`
	pattern := regexp.MustCompile(exp)
	return pattern.Match([]byte(s))

}

func SeparateDays(r [][]string) *os.File {
	for _, row := range r {
		check := row[8]
		if moreThanOneNumberReg(check) {
			var days []int
			// Check if contains
			if strings.Contains(check, "1") {
				days = append(days, 1)
			}
			if strings.Contains(check, "2") {
				days = append(days, 2)
			}
			if strings.Contains(check, "3") {
				days = append(days, 3)
			}
			if strings.Contains(check, "4") {
				days = append(days, 4)
			}
			if strings.Contains(check, "5") {
				days = append(days, 5)
			}
			if strings.Contains(check, "6") {
				days = append(days, 6)
			}
			if strings.Contains(check, "7") {
				days = append(days, 7)
			}
			if strings.Contains(check, "1") {
				days = append(days, 1)
			}
			newLines := performSeparation(row, days)
		}
	}

}

func performSeparation(row []string, d []int) [][]string {
	var newRows [][]string
	for v := range d {
		from := row[6]
		to := row[7]
		cpy := row

		var newRow []string

		from_day, _ := time.Parse("02-01-2006", from)
		f_day_n := int(from_day.Weekday())

		to_day, _ := time.Parse("02-01-2006", to)
		t_day_n := int(to_day.Weekday())
		var l int
		var day int
		// df to from, dt to to, day to v
		if v < f_day_n {
			day += 7
			l = day - f_day_n
		} else {
			l = day - f_day_n
		}
		var m int
		if v > t_day_n {
			day -= 7
			m = day - t_day_n
		} else {
			m = day - t_day_n
		}

		cpy[8] = strings.Repeat(".", v-1) + strconv.Itoa(v) + strings.Repeat(".", 7-v)
		cpy[6] = from_day.AddDate(0, 0, l).Format("02-01-2006")
		cpy[7] = to_day.AddDate(0, 0, m).Format("02-01-2006")

		newRow = append(newRow, cpy...)

		// OD -> do przodu, DO do tylu, sprawdzenie dat
	}
	return newRows
}
