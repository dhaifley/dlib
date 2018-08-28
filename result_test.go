package dlib

import (
	"testing"
	"time"
)

func TestNewResult(t *testing.T) {
	r := NewResult(nil, nil, "test", 0, "", nil, nil)
	if r.Type != "test" {
		t.Errorf("Expected string: test, got: %v", r.Type)
	}

	r = NewErrorResult(NewError(0, "test"))
	if r.Type != "error" {
		t.Errorf("Expected string: error, got: %v", r.Type)
	}
}

func TestResultString(t *testing.T) {
	d := time.Date(1983, 2, 2, 0, 0, 0, 0, time.Local)
	r := Result{
		Opt:  DateOptions{Range: 1},
		Msg:  "testmsg",
		Err:  &Error{Code: 401, Msg: "Unauthorized"},
		Time: &d,
	}

	exp := `{"options":{"range":1},"message":"testmsg",` +
		`"error":{"code":401,"message":"Unauthorized"},"time":"1983-02-02T00:00:00-05:00"}`
	if r.String() != exp {
		t.Errorf("Expected string: %v, got: %v", exp, r.String())
	}
}

func TestResultEquals(t *testing.T) {
	cases := []struct {
		a        *Result
		b        *Result
		expected bool
	}{
		{
			a:        NewResult(nil, nil, "test", 0, "test", nil, nil),
			b:        NewResult(nil, nil, "test", 0, "test", nil, nil),
			expected: true,
		},
		{
			a:        NewResult(nil, nil, "test", 0, "test", nil, nil),
			b:        NewResult(nil, nil, "test", 0, "test2", nil, nil),
			expected: false,
		},
	}

	for _, c := range cases {
		c.a.Time = nil
		c.b.Time = nil
		result := c.a.Equals(c.b)
		if result != c.expected {
			t.Errorf("Expected bool: %v, got: %v", c.expected, result)
		}
	}
}
