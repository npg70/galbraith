package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type DateQualifier int

const (
	DATE_EXACT DateQualifier = iota
	DATE_ABOUT
	DATE_SAY
	DATE_BEFORE
	DATE_AFTER
)

type Date struct {
	qualifier DateQualifier
	day       int // 1-31
	month     int // 1-12
	year      int // four digits
}

func ParseTime(t time.Time) Date {
	return Date{
		qualifier: DATE_EXACT,
		day:       t.Day(),
		month:     int(t.Month()),
		year:      t.Year(),
	}
}

func ParseDate(s string) (Date, error) {
	d := &Date{}
	err := d.UnmarshalText([]byte(s))
	return *d, err
}
func (d *Date) UnmarshalText(text []byte) error {

	// TODO use bytes
	s := string(text)

	// replace "." and "-" with spaces
	s = strings.ReplaceAll(s, ".", " ")
	s = strings.ReplaceAll(s, "-", " ")

	// make lower case
	s = strings.ToLower(s)

	//
	parts := strings.Fields(s)
	if len(parts) == 0 {
		return fmt.Errorf("Date unparsable %q", s)
	}

	// Qualifier.. tbd
	switch parts[0] {
	case "abt", "about":
		d.qualifier = DATE_ABOUT
		parts = parts[1:]
	case "say":
		d.qualifier = DATE_SAY
		parts = parts[1:]
	case "before":
		d.qualifier = DATE_BEFORE
		parts = parts[1:]
	case "after":
		d.qualifier = DATE_AFTER
		parts = parts[1:]
	default:
		d.qualifier = DATE_EXACT
	}

	if len(parts) == 0 {
		return fmt.Errorf("Date unparsable %q", s)
	}

	val, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		// not a number.  Maybe a month
		mon := textMonth[parts[0]]
		if mon == 0 {
			return fmt.Errorf("unable to parse month in %q", s)
		}
		d.month = mon
		parts = parts[1:]
		// got month, get year
		val, err = strconv.ParseInt(parts[0], 10, 32)
		if err != nil || val < 1000 || val > 3000 {
			return fmt.Errorf("unable to parse year in %q", s)
		}
		d.year = int(val)
		return nil
	}
	// is it a 4 digit number?  It's a year
	if val > 1000 && val < 3000 {
		d.year = int(val)
		parts = parts[1:]
		if len(parts) == 0 {
			// just a year
			return nil
		}
		// not supporting YYYY MM DD right now
		return fmt.Errorf("GOt year followed by other stuff")
	}
	// not a year.. has to be a day
	if val < 1 || val > 31 {
		return fmt.Errorf("first argument should be a day number but got %d", val)
	}
	d.day = int(val)
	parts = parts[1:]
	if len(parts) != 2 {
		return fmt.Errorf("expected day month year, but got wrong args")
	}
	val, err = strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		// not a number.  Maybe a month
		mon := textMonth[parts[0]]
		if mon == 0 {
			return fmt.Errorf("unable to parse month in %q", s)
		}
		d.month = int(mon)
		parts = parts[1:]
	} else {
		if val < 1 || val > 31 {
			return fmt.Errorf("got month number of %d", val)
		}
		d.month = int(val)
		parts = parts[1:]
	}

	// got day and month.. get year
	val, err = strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return fmt.Errorf("expected year number got %q", parts[0])
	}
	if val < 1000 || val > 3000 {
		return fmt.Errorf("expected 4-digit year number got %q", parts[0])
	}
	d.year = int(val)
	return nil
}

func (d *Date) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

func (d Date) String() string {
	return d.DayMonYear()
}

// returns true if the Date is unknown or uninitized.
func (d Date) Unknown() bool {
	return d.year == 0
}

// qualifier ignored
func (d Date) ISO() string {
	if d.Unknown() {
		// unknown if right idea or not
		return ""
	}

	return fmt.Sprintf("%04d%02d%02d", d.year, d.month, d.day)
}

// ISO 8601, with dash  2022-01-16
// Qualifier ignored
func (d Date) ISODash() string {
	if d.Unknown() {
		// unknown if right idea or not
		return ""
	}

	return fmt.Sprintf("%04d-%02d-%02d", d.year, d.month, d.day)
}

// 1-2 Digit Day, Full Month, Full Year
func (d Date) DayMonthYear() string {
	if d.Unknown() {
		// unknown if right idea or not
		return ""
	}

	parts := []string{}
	switch d.qualifier {
	case DATE_EXACT:
		// NOP
	case DATE_ABOUT:
		parts = append(parts, "about")
	case DATE_SAY:
		parts = append(parts, "say")
	case DATE_BEFORE:
		parts = append(parts, "before")
	case DATE_AFTER:
		parts = append(parts, "after")
	default:
		panic("oops forgot to add qualifier")
	}
	if d.day != 0 {
		parts = append(parts, fmt.Sprintf("%d", d.day))
	}
	if d.month != 0 {
		parts = append(parts, numToMonth(d.month))
	}
	parts = append(parts, fmt.Sprintf("%d", d.year))
	return strings.Join(parts, " ")
}

// 1-2 Digit Day, 3 letter Month, Full Year
func (d Date) DayMonYear() string {
	parts := []string{}
	switch d.qualifier {
	case DATE_EXACT:
		// NOP
	case DATE_ABOUT:
		parts = append(parts, "about")
	case DATE_SAY:
		parts = append(parts, "say")
	case DATE_BEFORE:
		parts = append(parts, "before")
	case DATE_AFTER:
		parts = append(parts, "after")
	default:
		panic("oops forgot to add qualifier")
	}
	if d.day != 0 {
		parts = append(parts, fmt.Sprintf("%d", d.day))
	}
	if d.month != 0 {
		parts = append(parts, numToMonth3(d.month))
	}
	if d.year != 0 {
		parts = append(parts, fmt.Sprintf("%d", d.year))
	}
	return strings.Join(parts, " ")
}

func numToMonth(m int) string {
	return []string{"", "January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November", "December"}[m]
}

func numToMonth3(m int) string {
	return []string{"", "Jan", "Feb", "Mar", "Apr", "May", "Jun",
		"Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}[m]
}

var textMonth = map[string]int{
	"jan":       1,
	"feb":       2,
	"mar":       3,
	"apr":       4,
	"may":       5,
	"jun":       6,
	"jul":       7,
	"aug":       8,
	"sep":       9,
	"oct":       10,
	"nov":       11,
	"dec":       12,
	"january":   1,
	"february":  2,
	"march":     3,
	"april":     4,
	"june":      6,
	"july":      7,
	"august":    8,
	"sept":      9,
	"september": 9,
	"october":   10,
	"november":  11,
	"december":  12,
}
