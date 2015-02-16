package norway

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/steevel/idno-parsers"
)

var weights = [][]int{
	{3, 7, 6, 1, 8, 9, 4, 5, 2},
	{5, 4, 3, 2, 7, 6, 5, 4, 3, 2},
}

// Norway specific error messages
var (
	ErrInvalidYearInd = errors.New("Invalid year/ind combination")
)

// Norway national identification number types
const (
	Regular = "Regular"
	D       = "D"
	H       = "H"
)

func init() {
	parsers.Register("norway", Parse)
}

// Parse id
/**
 * DD MM YY III KK (11 digits)
 * |  |  |  |   |
 * |  |  |  |   `- Two check digits
 * |  |  |  `- Individual number (3 digits, last one is sex, woman/man even/odd)
 * |  |  `- Year (2 digits)
 * |  `- Month (2 digits)
 * `- Day (2 digits)
 */
func Parse(id string) (*parsers.IdNo, error) {
	n := parsers.New(id)

	if _, e := strconv.Atoi(n.Source); e != nil || len(n.Source) != 11 {
		return n, parsers.ErrInvalidFormat
	}

	tt := strings.Split(n.Source, "")
	source := make([]int, len(tt))
	for i := range tt {
		source[i], _ = strconv.Atoi(tt[i])
	}

	// Control checkdigits
	for _, weight := range weights {
		sum := 0
		for i, w := range weight {
			sum += source[i] * w
		}

		checkdigit := 11 - (sum % 11)
		if checkdigit == 11 {
			checkdigit = 0
		}

		if source[len(weight)] != checkdigit {
			return n, parsers.ErrInvalidCheckdigit
		}
	}

	// Extract year and individual number
	year := source[4]*10 + source[5]
	day := source[0]*10 + source[1]
	month := source[2]*10 + source[3]
	ind := source[6]*100 + source[7]*10 + source[8]

	// Add to year based on year/ind combination
	if ind <= 499 || (ind >= 900 && ind <= 999 && year >= 40 && year <= 99) {
		year += 1900
	} else if ind >= 500 && ind <= 749 && year >= 54 && year <= 99 {
		year += 1800
	} else if ind >= 500 && ind <= 999 && year <= 39 {
		year += 2000
	} else {
		return n, ErrInvalidYearInd
	}

	// Handle D and H numbers
	ninType := Regular
	if day >= 40 && day <= 71 {
		day -= 40
		ninType = D
	} else if month >= 40 && month <= 52 {
		month -= 40
		ninType = H
	}

	// Sanity check, is this a correct date?
	dateString := fmt.Sprintf("%d-%02d-%02d", year, month, day)
	if _, e := time.Parse("2006-01-02", dateString); e != nil {
		return n, parsers.ErrInvalidDate
	}

	// Everything looks OK, go ahead and populate return value
	n.Valid = true
	n.Data["gender"] = []string{parsers.Female, parsers.Male}[source[8]%2]
	n.Data["type"] = ninType
	n.Data["year"] = strconv.Itoa(year)
	n.Data["month"] = fmt.Sprintf("%02s", strconv.Itoa(month))
	n.Data["day"] = fmt.Sprintf("%02s", strconv.Itoa(day))

	return n, nil
}
