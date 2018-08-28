package dlib

import (
	"testing"
)

func TestServiceInfoString(t *testing.T) {
	a := NewServiceInfo("test", "test", "test", "0.0.0")
	expected := `{"name":"test","short":"test","long":"test","version":"0.0.0"}`
	result := a.String()
	if result != expected {
		t.Errorf("Expected string: %v, got: %v", expected, result)
	}
}
