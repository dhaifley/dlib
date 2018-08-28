package dlib

import (
	"testing"
	"time"
)

func TestNullTimeScan(t *testing.T) {
	ct := time.Now()
	var nt NullTime
	nt.Scan(ct)
	if !nt.Valid {
		t.Errorf("Bool expected: true, got: false")
	}
}

func TestNullTimeValue(t *testing.T) {
	nt := NullTime{Time: time.Now(), Valid: true}
	_, err := nt.Value()
	if err != nil {
		t.Errorf("Error expected: nil, got: %v", err)
	}
}

func TestDateOptionsEquals(t *testing.T) {
	cases := []struct {
		a   DateOptions
		b   DateOptions
		exp bool
	}{
		{
			a:   DateOptions{},
			b:   DateOptions{},
			exp: true,
		},
		{
			a: DateOptions{
				Offset: 1,
				Range:  2,
			},
			b: DateOptions{
				Offset: 1,
				Range:  1,
			},
			exp: false,
		},
	}

	for _, c := range cases {
		if c.a.Equals(c.b) != c.exp {
			t.Errorf("Bool expected: %v, got: %v", c.exp, c.a.Equals(c.b))
		}
	}
}

func TestDateOptionsCalculate(t *testing.T) {
	d := time.Date(1983, 2, 2, 0, 0, 0, 0, time.Local)
	ed := time.Date(1983, 2, 2, 0, 0, 0, 0, time.Local)
	ed = ed.Add(time.Hour * time.Duration(-24))
	sd := ed.Add(time.Hour * time.Duration(-24*2))
	cases := []struct {
		do  DateOptions
		exp DateOptions
	}{
		{
			do: DateOptions{
				Start: &d,
				End:   &d,
			},
			exp: DateOptions{
				Start: &d,
				End:   &d,
			},
		},
		{
			do: DateOptions{
				Start:  &d,
				End:    &d,
				Offset: 1,
				Range:  2,
			},
			exp: DateOptions{
				Start:  &sd,
				End:    &ed,
				Offset: 1,
				Range:  2,
			},
		},
	}

	for _, c := range cases {
		c.do.Calculate()
		if !c.do.Equals(c.exp) {
			t.Errorf("DateOptions expected: %v, got: %v", c.exp, c.do)
		}
	}
}

func TestParseDateTimeCode(t *testing.T) {
	tm, err := ParseDateTimeCode("19" + "830202051530")
	if err != nil {
		t.Error(err)
	}

	expected := time.Date(1983, 2, 2, 5, 15, 30, 0, time.Local)
	if *tm != expected {
		t.Errorf("Expected time: %v, got: %v", expected, *tm)
	}
}

func TestParseDateCode(t *testing.T) {
	tm, err := ParseDateCode("19830202")
	if err != nil {
		t.Error(err)
	}

	expected := time.Date(1983, 2, 2, 0, 0, 0, 0, time.Local)
	if *tm != expected {
		t.Errorf("Expected time: %v, got: %v", expected, *tm)
	}
}

func TestGetBusinessDate(t *testing.T) {
	tm := time.Date(1983, 2, 2, 4, 15, 30, 0, time.Local)
	bd := GetBusinessDate(&tm)
	expected := time.Date(1983, 2, 1, 0, 0, 0, 0, time.Local)
	if *bd != expected {
		t.Errorf("Expected time: %v, got: %v", expected, *bd)
	}
}

func TestNow(t *testing.T) {
	type i interface {
		Unix() int64
	}

	var tm i
	tm = Now()
	if _, ok := tm.(*time.Time); !ok {
		t.Error("Expected type: *time.Time")
	}
}
