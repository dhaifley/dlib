package dlib

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
)

// EncryptString applies an encryption method to the provided string.
// It returns an encrypted string.
func EncryptString(s string) (string, error) {
	key := []byte(">^_^<>^_^<>^_^<@@>^_^<>^_^<>^_^<")
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce, err := hex.DecodeString("a6abc9cb6e899632a5131309")
	if err != nil {
		return "", err
	}

	ciphertext := aesgcm.Seal(nil, nonce, []byte(s), nil)
	hext := hex.EncodeToString(ciphertext)
	return hext, nil
}

// DecryptString applies an decryption method to the provided string.
// It returns an decrypted string.
func DecryptString(s string) (string, error) {
	key := []byte(">^_^<>^_^<>^_^<@@>^_^<>^_^<>^_^<")
	ciphertext, err := hex.DecodeString(s)
	if err != nil {
		return "", err
	}

	nonce, err := hex.DecodeString("a6abc9cb6e899632a5131309")
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

// EncodeBase64String applies base64 encoding to the provided string.
// It returns an encoded string.
func EncodeBase64String(s string) string {
	return base64.URLEncoding.EncodeToString([]byte(s))
}

// DecodeBase64String removes base64 encoding from the provided string.
// It returns a decoded string.
func DecodeBase64String(s string) (string, error) {
	str, err := base64.URLEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}

	return string(str), nil
}
