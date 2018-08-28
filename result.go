package dlib

import (
	"encoding/json"
	"time"
)

// Result values contain information about the result of a command.
type Result struct {
	Opt  interface{} `json:"options,omitempty"`
	Val  interface{} `json:"value,omitempty"`
	Type string      `json:"type,omitempty"`
	Num  int         `json:"number,omitempty"`
	Msg  string      `json:"message,omitempty"`
	Err  error       `json:"error,omitempty"`
	Data interface{} `json:"data,omitempty"`
	Time *time.Time  `json:"time,string,omitempty"`
}

// NewResult creates a new result value.
func NewResult(
	opt interface{},
	val interface{},
	typ string,
	num int,
	msg string,
	err error,
	data interface{}) *Result {
	return &Result{
		Opt:  opt,
		Val:  val,
		Type: typ,
		Num:  num,
		Msg:  msg,
		Err:  err,
		Data: data,
		Time: Now(),
	}
}

// NewErrorResult creates a new result value specifcally to return an error.
func NewErrorResult(err error) *Result {
	return &Result{Type: "error", Err: err, Time: Now()}
}

// String returns the Result object as a string.
func (r *Result) String() string {
	str, err := json.Marshal(r)
	if err != nil {
		return ""
	}

	return string(str)
}

// Equals tests for deep equality between result values.
func (r *Result) Equals(b *Result) bool {
	switch {
	case r == nil || b == nil:
		return false
	case r == b:
		return true
	case r.Opt != b.Opt:
		return false
	case r.Val != b.Val:
		return false
	case r.Type != b.Type:
		return false
	case r.Num != b.Num:
		return false
	case r.Msg != b.Msg:
		return false
	case r.Data != b.Data:
		return false
	case r.Err != b.Err:
		return false
	case (r.Time == nil && b.Time != nil) || (r.Time != nil && b.Time == nil):
		return false
	case r.Time != nil && b.Time != nil && *r.Time != *b.Time:
		return false
	default:
		return true
	}
}
