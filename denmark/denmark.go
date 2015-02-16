package denmark

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
)

var (
	reId = regexp.MustCompile("^(\\d{6})-?(\\d{3,4})$")
)

func init() {
	parsers.Register("denmark", Parse)
}

// Parse id
/**
 * DDMMYY-?FFGC?
 */
func Parse(id string) (*parsers.IdNo, error) {
	n := parsers.New(id)

	var parsed []string
	foo := reId.FindAllStringSubmatch(id, -1)
	if len(foo) == 0 {
		return n, parsers.ErrInvalidFormat
	}
	parsed = foo[0]

	tt := strings.Split(parsed[1]+parsed[2], "")
	source := make([]int, len(tt))
	for i := range tt {
		source[i], _ = strconv.Atoi(tt[i])
	}

	// Apply weights
	sum := 0
	weight := 9
	for _, num := range source[0:9] {
		sum += num * weight
		weight--
	}

	// We got a checkdigit, so let's check if it's correct
	if len(source) == 10 {
		if sum%11 != 0 {
			return n, parsers.ErrInvalidCheckdigit
		}
	}

	year := source[0]*10 + source[1]
	month := source[2]*10 + source[3]
	day := source[4]*10 + source[5]

	year += 1900

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
	n.Data["type"] = Regular

	return n, nil
}
