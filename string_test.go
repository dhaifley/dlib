package dlib

import "testing"

func TestXMLEncode(t *testing.T) {
	exp := "&amp;&lt;&gt;&quot;&apos;"
	res := XMLEncode(`&<>"'`)
	if res != exp {
		t.Errorf("Expected string: %q, got: %q", exp, res)
	}
}

func TestXMLDecode(t *testing.T) {
	exp := `&<>"'`
	res := XMLDecode("&amp;&lt;&gt;&quot;&apos;")
	if res != exp {
		t.Errorf("Expected string: %q, got: %q", exp, res)
	}
}

func TestPadString(t *testing.T) {
	exp := "0ab"
	res := PadString("ab", 3, "0")
	if res != exp {
		t.Errorf("Expected string: %q, got: %q", exp, res)
	}
}

func TestFormatUPC(t *testing.T) {
	exp := "00000testup"
	res := FormatUPC("testupc", 11, false)
	if res != exp {
		t.Errorf("Expected string: %q, got: %q", exp, res)
	}
}
