package sweden

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/steevel/idno-parsers"
)

// National identification number types
var (
	Regular = "Regular"
	TP      = "TP"
)

var (
	reID = regexp.MustCompile("^(\\d{2})?(\\d{6})([+-])?(\\d{4})$")
)

func init() {
	parsers.Register("sweden", Parse)
}

// Parse id
/*
 * Takes two forms
 * - YYYYMMDD[-]?FFGC
 * - YYMMDD[-+]?FFGC (the correct form
 */
func Parse(id string) (*parsers.IdNo, error) {
	n := parsers.New(id)

	// Parse and see if valid
	match := reID.FindAllStringSubmatch(id, -1)
	if len(match) == 0 {
		return n, parsers.ErrInvalidFormat
	}
	parsed := match[0]

	// Create int array of the parsed id
	tt := strings.Split(parsed[2]+parsed[4], "")
	source := make([]int, len(tt))
	for i := range tt {
		source[i], _ = strconv.Atoi(tt[i])
	}

	// Apply weights
	str := ""
	for i, num := range source[0 : len(source)-1] {
		str += strconv.Itoa(num * []int{2, 1}[i%2])
	}

	// Create a checkdigit
	checkdigit := 0
	for _, num := range str {
		t, _ := strconv.Atoi(string(num))
		checkdigit += t
	}

	if source[len(source)-1] != 10-checkdigit%10 {
		return n, parsers.ErrInvalidCheckdigit
	}

	year := source[0]*10 + source[1]
	month := source[2]*10 + source[3]
	day := source[4]*10 + source[5]

	Type := Regular
	if day >= 61 && day <= 91 {
		Type = TP
		day -= 60
	}

	if len(parsed[1]) > 0 {
		n, _ := strconv.Atoi(parsed[1])
		year += n * 100
	} else {
		// 100+
		if parsed[3] == "+" {
			// hh := time.Now().Year() / 100 * 100
		}
		year += 1900
	}

	// Sanity check, is this a correct date?
	dateString := fmt.Sprintf("%d-%02d-%02d", year, month, day)
	if _, e := time.Parse("2006-01-02", dateString); e != nil {
		return n, parsers.ErrInvalidDate
	}

	// Everything looks OK, go ahead and populate return value
	n.Valid = true
	n.Data["gender"] = []string{parsers.Female, parsers.Male}[source[8]%2]
	n.Data["year"] = strconv.Itoa(year)
	n.Data["month"] = fmt.Sprintf("%02s", strconv.Itoa(month))
	n.Data["day"] = fmt.Sprintf("%02s", strconv.Itoa(day))
	n.Data["type"] = Type

	return n, nil
}
