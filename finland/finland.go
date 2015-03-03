package finland

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/steevel/idno-parsers"
)

func init() {
	parsers.Register("finland", Parse)
}

const (
	ctrlChrs = "0123456789ABCDEFHJKLMNPRSTUVWXY"
	reTpl    = "^(\\d{6})([-+A])(\\d{3})([%s])$"
)

var (
	reID = regexp.MustCompile(fmt.Sprintf(reTpl, ctrlChrs))
)

// Parse id
/*
 * Accepts one format
 *  - DDMMYY[+-A]FFFC
 */
func Parse(id string) (*parsers.IdNo, error) {
	var idNo = parsers.New(id)

	// Parse and see if valid
	match := reID.FindAllStringSubmatch(id, -1)
	if len(match) == 0 {
		return idNo, parsers.ErrInvalidFormat
	}

	// Extract the date, century, increment and check character
	idDate := match[0][1]
	idCent := match[0][2]
	idIncr := match[0][3]
	idCtrl := match[0][4]

	// Combine date of birth with incremental number, divide by 31 to get
	// check character
	dateAndIncr, _ := strconv.Atoi(idDate + idIncr)
	ctrlNum := dateAndIncr % 31
	ctrlChr := ctrlChrs[ctrlNum : ctrlNum+1]

	// Check authenticity
	if ctrlChr != idCtrl {
		return idNo, parsers.ErrInvalidCheckdigit
	}

	year, month, day, err := parseDate(idDate, idCent)
	if err != nil {
		return idNo, err
	}

	// The incremental number is uneven for men and even for women
	gender := parsers.Male
	if i, _ := strconv.Atoi(idIncr); i%2 == 0 {
		gender = parsers.Female
	}

	// Everything looks OK, go ahead and populate return value
	idNo.Valid = true
	idNo.Data["gender"] = gender
	idNo.Data["year"] = year
	idNo.Data["month"] = month
	idNo.Data["day"] = day

	return idNo, nil
}

// Calculate century and extract year, month and day from the date
func parseDate(date string, centMark string) (year string, month string, day string, err error) {
	var century string

	switch centMark {
	case "+":
		century = "18"
	case "-":
		century = "19"
	case "A":
		century = "20"
	}

	day = date[:2]
	month = date[2:4]
	year = century + date[4:6]

	// Sanity check, is this a correct date?
	dateString := fmt.Sprintf("%s-%s-%s", year, month, day)
	if _, e := time.Parse("2006-01-02", dateString); e != nil {
		err = parsers.ErrInvalidDate
	}

	return
}
