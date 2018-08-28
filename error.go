package dlib

import (
	"encoding/json"
	"time"
)

// Error values contain information about error conditions.
type Error struct {
	Code int        `json:"code,omitempty"`
	Msg  string     `json:"message,omitempty"`
	Time *time.Time `json:"time,string,omitempty"`
}

// NewError creates a new error value.
func NewError(code int, msg string) *Error {
	return &Error{Code: code, Msg: msg, Time: Now()}
}

// String returns the Error object as a string.
func (e *Error) String() string {
	str, err := json.Marshal(e)
	if err != nil {
		return ""
	}

	return string(str)
}

// Error returns the Error object formatted as a JSON string.
func (e *Error) Error() string {
	return e.String()
}

// Equals tests for deep equality between error values.
func (e *Error) Equals(b *Error) bool {
	switch {
	case e == nil || b == nil:
		return false
	case e == b:
		return true
	case e.Code != b.Code:
		return false
	case e.Msg != b.Msg:
		return false
	case (e.Time == nil && b.Time != nil) || (e.Time != nil && b.Time == nil):
		return false
	case e.Time != nil && b.Time != nil && *e.Time != *b.Time:
		return false
	default:
		return true
	}
}
