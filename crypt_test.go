package dlib

import "testing"

func TestEncryptDecryptString(t *testing.T) {
	enc, err := EncryptString("test")
	if err != nil {
		t.Error(err)
	}

	dec, err := DecryptString(enc)
	if err != nil {
		t.Error(err)
	}

	expected := "test"
	if dec != expected {
		t.Errorf("Expected string: %s, got: %s", expected, dec)
	}
}

func TestEncodeDecodeString(t *testing.T) {
	enc := EncodeBase64String("test")
	dec, err := DecodeBase64String(enc)
	if err != nil {
		t.Error(err)
	}

	expected := "test"
	if dec != expected {
		t.Errorf("Expected string: %s, got: %s", expected, dec)
	}
}
