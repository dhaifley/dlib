package dlib

import (
	"testing"
)

func TestErrorString(t *testing.T) {
	e := NewError(401, "Unauthorized")
	e.Time = nil
	expected := `{"code":401,"message":"Unauthorized"}`
	if e.String() != expected {
		t.Errorf("Expected string: %v, got: %v", expected, e.String())
	}
}

func TestErrorEquals(t *testing.T) {
	cases := []struct {
		a        *Error
		b        *Error
		expected bool
	}{
		{
			a:        &Error{Code: 400, Msg: "test"},
			b:        &Error{Code: 400, Msg: "test"},
			expected: true,
		},
		{
			a:        &Error{Code: 400, Msg: "test"},
			b:        &Error{Code: 401, Msg: "test"},
			expected: false,
		},
	}

	for _, c := range cases {
		result := c.a.Equals(c.b)
		if result != c.expected {
			t.Errorf("Expected bool: %v, got: %v", c.expected, result)
		}
	}
}
