package dlib

import (
	"database/sql/driver"
	"encoding/xml"
	"strconv"
	"time"
)

// NullTime values represent nullable date/time field values in sql rows.
type NullTime struct {
	Time  time.Time
	Valid bool
}

// Scan implements the Scanner interface for NullTime values.
func (nt *NullTime) Scan(value interface{}) error {
	nt.Time, nt.Valid = value.(time.Time)
	return nil
}

// Value implements the sql driver Valuer interface for NullTime values.
func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}

	return nt.Time, nil
}

// XMLTime values represent time data that will unmarshall from XML data.
type XMLTime struct {
	time.Time
}

// UnmarshalXML implements the XMLUnmarshaller interface for XMLTime values.
func (xt *XMLTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)
	if v == "" {
		return nil
	}

	parse, err := time.ParseInLocation("2006-01-02T15:04:05.000", v, time.Local)
	if err != nil {
		return err
	}

	*xt = XMLTime{parse}
	return nil
}

// UnmarshalXMLAttr implements the XMLUnmarshaller interface for XMLTime values
// representing attributes.
func (xt *XMLTime) UnmarshalXMLAttr(attr xml.Attr) error {
	if attr.Value == "" {
		return nil
	}

	parse, err := time.ParseInLocation("2006-01-02T15:04:05.000", attr.Value, time.Local)
	if err != nil {
		return err
	}

	*xt = XMLTime{parse}
	return nil
}

// Scan implements the Scanner interface for XMLTime values.
func (xt *XMLTime) Scan(value interface{}) error {
	val, ok := value.(time.Time)
	if !ok {
		return NewError(500, "invalid time")
	}

	xt.Time = val
	return nil
}

// Value implements the sql driver Valuer interface for XMLTime values.
func (xt *XMLTime) Value() (driver.Value, error) {
	if xt != nil {
		return xt.Time, nil
	}

	return nil, nil
}

// XMLDate values represent date data that will unmarshall from XML data.
type XMLDate struct {
	time.Time
}

// UnmarshalXML implements the XMLUnmarshaller interface for XMLDate values.
func (xd *XMLDate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)
	if v == "" {
		return nil
	}

	parse, err := time.ParseInLocation("2006-01-02", v, time.Local)
	if err != nil {
		return err
	}

	*xd = XMLDate{parse}
	return nil
}

// UnmarshalXMLAttr implements the XMLUnmarshaller interface for XMLDate values
// representing attributes.
func (xd *XMLDate) UnmarshalXMLAttr(attr xml.Attr) error {
	if attr.Value == "" {
		return nil
	}

	parse, err := time.ParseInLocation("2006-01-02", attr.Value, time.Local)
	if err != nil {
		return err
	}

	*xd = XMLDate{parse}
	return nil
}

// Scan implements the Scanner interface for XMLDate values.
func (xd *XMLDate) Scan(value interface{}) error {
	val, ok := value.(time.Time)
	if !ok {
		return NewError(500, "invalid date")
	}

	xd.Time = val
	return nil
}

// Value implements the sql driver Valuer interface for XMLDate values.
func (xd *XMLDate) Value() (driver.Value, error) {
	if xd != nil {
		return xd.Time, nil
	}

	return nil, nil
}

// DateOptions instaces describe a range of days in a standard format.
type DateOptions struct {
	Start  *time.Time `json:"startDate,string,omitempty"`
	End    *time.Time `json:"endDate,string,omitempty"`
	Offset int        `json:"offset,omitempty"`
	Range  int        `json:"range,omitempty"`
}

// Equals checks whether two DateOptions instances are equivalent.
func (do DateOptions) Equals(b DateOptions) bool {
	if &do == &b {
		return true
	}

	if (do.Start == nil && b.Start != nil) || (do.Start != nil && b.Start == nil) {
		return false
	}
	if do.Start != nil && b.Start != nil && *do.Start != *b.Start {
		return false
	}

	if (do.End == nil && b.End != nil) || (do.End != nil && b.End == nil) {
		return false
	}
	if do.End != nil && b.End != nil && *do.End != *b.End {
		return false
	}

	if do.Offset != b.Offset {
		return false
	}

	if do.Range != b.Range {
		return false
	}

	return true
}

// Calculate derives all values for a DateOptions inststance when only
// a partial set of values are provided.
func (do *DateOptions) Calculate() {
	if do.Start == nil {
		st := time.Now()
		do.Start = &st
	}

	if do.End == nil {
		et := time.Now()
		do.End = &et
	}

	if do.Start == do.End {
		et := *do.End
		do.Start = &et
	}

	if do.Offset != 0 {
		*do.Start = do.Start.Add(time.Hour * time.Duration(-24*do.Offset))
		*do.End = do.End.Add(time.Hour * time.Duration(-24*do.Offset))
	}

	if do.Range != 0 {
		*do.Start = do.Start.Add(time.Hour * time.Duration(-24*do.Range))
	}

	if do.Start.Unix() > do.End.Unix() {
		*do.Start = *do.End
	}
}

// ParseDateTimeCode parses a Time structure from a date time code string.
func ParseDateTimeCode(s string) (*time.Time, error) {
	var t time.Time
	if len(s) < 14 {
		return nil, &Error{Code: 500, Msg: "invalid date time code string"}
	}

	year, err := strconv.ParseInt(s[0:4], 10, 64)
	if err != nil {
		return nil, err
	}

	mon, err := strconv.ParseInt(s[4:6], 10, 64)
	if err != nil {
		return nil, err
	}

	day, err := strconv.ParseInt(s[6:8], 10, 64)
	if err != nil {
		return nil, err
	}

	hour, err := strconv.ParseInt(s[8:10], 10, 64)
	if err != nil {
		return nil, err
	}

	min, err := strconv.ParseInt(s[10:12], 10, 64)
	if err != nil {
		return nil, err
	}

	sec, err := strconv.ParseInt(s[12:14], 10, 64)
	if err != nil {
		return nil, err
	}

	t = time.Date(int(year), time.Month(mon), int(day),
		int(hour), int(min), int(sec), 0, time.Local)
	return &t, nil
}

// ParseDateCode parses a Time structure from a date code string.
func ParseDateCode(s string) (*time.Time, error) {
	var t time.Time
	if len(s) < 8 {
		return nil, &Error{Code: 500, Msg: "invalid date code string"}
	}

	year, err := strconv.ParseInt(s[0:4], 10, 64)
	if err != nil {
		return nil, err
	}

	mon, err := strconv.ParseInt(s[4:6], 10, 64)
	if err != nil {
		return nil, err
	}

	day, err := strconv.ParseInt(s[6:8], 10, 64)
	if err != nil {
		return nil, err
	}

	t = time.Date(int(year), time.Month(mon), int(day), 0, 0, 0, 0, time.Local)
	return &t, nil
}

// GetBusinessDate returns the business date of a sepecifed time.
func GetBusinessDate(t *time.Time) *time.Time {
	bd := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	sbd := bd.Add(time.Hour * 5)
	if t.Unix() < sbd.Unix() {
		bd = bd.Add(time.Hour * -24)
	}

	return &bd
}

// Now returns a pointer to a time.Time structure containing the current time.
func Now() *time.Time {
	t := time.Now()
	return &t
}
