package dlib

import (
	"testing"
	"time"

	"google.golang.org/grpc/credentials"
)

// The certificate and key data in this file was created for testing purposes
// only. They are not real or used in any actual projects.

const testCrt = `
-----BEGIN CERTIFICATE-----
MIIFNTCCBB2gAwIBAgIIUdhUn2xoanEwDQYJKoZIhvcNAQELBQAwgbQxCzAJBgNV
BAYTAlVTMRAwDgYDVQQIEwdBcml6b25hMRMwEQYDVQQHEwpTY290dHNkYWxlMRow
GAYDVQQKExFHb0RhZGR5LmNvbSwgSW5jLjEtMCsGA1UECxMkaHR0cDovL2NlcnRz
LmdvZGFkZHkuY29tL3JlcG9zaXRvcnkvMTMwMQYDVQQDEypHbyBEYWRkeSBTZWN1
cmUgQ2VydGlmaWNhdGUgQXV0aG9yaXR5IC0gRzIwHhcNMTcwNTA1MTU1MzAwWhcN
MjAwNTA1MTU1MzAwWjA+MSEwHwYDVQQLExhEb21haW4gQ29udHJvbCBWYWxpZGF0
ZWQxGTAXBgNVBAMMECoucm95YWxmYXJtcy5jb20wggEiMA0GCSqGSIb3DQEBAQUA
A4IBDwAwggEKAoIBAQCs9UN/tRNf0LKwgvSqbvvZU31K4H/VOq6sywP2cCM1+ZZs
r87tjm3+HOJWnXKJjgnpXrC29pM1rlyx/EwLPoEMy0BooApHlmTq4VO+vrnTLkj1
+27a4rAjpT3kzwTul3mVjdGVf+xJfzPOkX0iu1+wxEherPDfCBCqS5FCJBFTvggk
bHADL2S4ScX6rBIU4/YlPM+55pWFVIuuX5Va8iFpHbtB75ZiQ1vF7cJ7T4x+dwze
aPCrsTqQLlS4x06NvISP1Wn+jCy1Ii3vGua4u1o//KrkWquaQfXftTEz+jRaP4Qq
Tiv2JSdIyanLc/8bWFtAWY6d5torWM40HrKbjF4rAgMBAAGjggG+MIIBujAMBgNV
HRMBAf8EAjAAMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAOBgNVHQ8B
Af8EBAMCBaAwNwYDVR0fBDAwLjAsoCqgKIYmaHR0cDovL2NybC5nb2RhZGR5LmNv
bS9nZGlnMnMxLTUwMi5jcmwwXQYDVR0gBFYwVDBIBgtghkgBhv1tAQcXATA5MDcG
CCsGAQUFBwIBFitodHRwOi8vY2VydGlmaWNhdGVzLmdvZGFkZHkuY29tL3JlcG9z
aXRvcnkvMAgGBmeBDAECATB2BggrBgEFBQcBAQRqMGgwJAYIKwYBBQUHMAGGGGh0
dHA6Ly9vY3NwLmdvZGFkZHkuY29tLzBABggrBgEFBQcwAoY0aHR0cDovL2NlcnRp
ZmljYXRlcy5nb2RhZGR5LmNvbS9yZXBvc2l0b3J5L2dkaWcyLmNydDAfBgNVHSME
GDAWgBRAwr0njsw0gzCiM9f7bLPwtCyAzjArBgNVHREEJDAighAqLnJveWFsZmFy
bXMuY29tgg5yb3lhbGZhcm1zLmNvbTAdBgNVHQ4EFgQUWN1+hI1gcguUsEvDQDrS
l3mTyCowDQYJKoZIhvcNAQELBQADggEBAEyAI8qQwmIZn1eeeuwmqXGbhK7/gXFh
XEzHsEUXrMi0t/wFcAG5gDgn+YZK9AD/lJkLLMnlcIjAzkqtwXbqdV3Nm/yNfq9C
hc2IzqHQ/7L3vGY6DTVHjEdOWGTJkGt0hqbF+LO+tM+wBx+G57E7EzGGTR7wTSQE
oRQLmesKFKbdrX22UANpULKlf1yfun/wz5Zear0RAIVfb/xBR/Bfhto6R38dv4Da
4QCAZTPQyC2xPzgYpfZJuzHIdMw6/vRPIqxUEHSHVsoE0MeTHS49VhEfAW8/p5X9
2CKOcjBYk9xQ2kKCYUZtpGyJJ318qiYckG+eu2dLvUo61OvboABTgiY=
-----END CERTIFICATE-----`

const testKey = `
-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCs9UN/tRNf0LKw
gvSqbvvZU31K4H/VOq6sywP2cCM1+ZZsr87tjm3+HOJWnXKJjgnpXrC29pM1rlyx
/EwLPoEMy0BooApHlmTq4VO+vrnTLkj1+27a4rAjpT3kzwTul3mVjdGVf+xJfzPO
kX0iu1+wxEherPDfCBCqS5FCJBFTvggkbHADL2S4ScX6rBIU4/YlPM+55pWFVIuu
X5Va8iFpHbtB75ZiQ1vF7cJ7T4x+dwzeaPCrsTqQLlS4x06NvISP1Wn+jCy1Ii3v
Gua4u1o//KrkWquaQfXftTEz+jRaP4QqTiv2JSdIyanLc/8bWFtAWY6d5torWM40
HrKbjF4rAgMBAAECggEASYp3yuNZe5IniHoGQjmyiHPBgIb4k9fB0eL4ql5/+kFH
vqr6V3QKPNRXQPDtdKOaf0ot5X3ulhKvg1Z9lwJaqO/7UQFLnajK/DwW+bHrHWo+
x6jDN/rPXjiymomT1Uc/AWItzV15QL7/bkPaa1w0kdrD0s3CFXM+cspy1yay32Ho
l+uY2h8eFBgKHAh35EoHExWv+hTWK42lY1DBr7Zo9n0WYNJP8Xinz+GUAzzFDV4C
SLRFpkYD+Z2e1b1HH31nWwqDih9pckircd2Uw/HqpPLCn+Lk3fTgXPV0q8oF/PwO
xmdJpuAHMS4phW2PXOKjOjk4rdUK7JB9Ndpv+njY8QKBgQDfzDjPNZU1dZh3AARD
imcgJ1r+O/jLA2DkWsA89f6lgHmMRKmrcj21DfQYJWa2F3DSu2xytuhdmy4OI0ig
uu0uDK9TTSXffagNRiwYBOe2Jnjmp4+8H1hiuqZBAyepBJeHT7b+XZeCdAk9Ixfi
V5rP9ZTrX4WPHkak/mpBz1wNfwKBgQDF2FHA+lxa/z3rjwRyiYO7W6wtCBKlccQq
34Se7SwkYwcZa2X9t6ojgezzqEV5inBf9iGRGLVK+1f8w7jOJ28gMhqU1y96h2a+
8aV0PCck0UZnClfWBlozyXxbUOV9UxoSZuQgXYpDM2Z16/cDZ6bVT8L5ys/4FEwT
p4OyZsudVQKBgQCbMLz8Q4Xyildtvd83DtSwYVoiDmhaLz+TWrMQEu3AyrR+5mZZ
82CDGuf3jogJIXwlRb7QVbIQpzlqqGEGxFkQo884jrkCn9pXSh/tkAk2MLuKBwwP
QhVCcXg0gQGRnROOy2J0RWZ5GgoGET3QwTsjloLsLVMqia4nYB1DAY4t9QKBgQCz
FHgDvlqEf47TvnOfHTwhBfyWEj2WhaCz0fhgRnvzP7O4XY9HFc0qMLuDXWfteNL4
+Xgutve6tEFTaPHJoMYklVWcLh8qwI7O/TnvOVeRKuCX+jPBZaSWRaWLnnDrfKIX
0AxkA2dYJply2bCP422OnZD1u499lRm7nKBHhmZ9tQKBgB/H9T/Lxkc8SzYD65Ud
hv9YiKaun4aVBYWoNUq+5o7wTVGE6aHykjgZKj5/Q/dcnw4x41GGc5qi6Y1A0lKK
OSvyLmrBgAt7jkdkEeeYdwpnm18/B6TNpD2Ng/mXrjY9OWUa5ytLiHgEd9P22nMr
OuPUUter1I/sY71FinjJ1HeB
-----END PRIVATE KEY-----`

func TestGetHTTPSClient(t *testing.T) {
	hc, err := GetHTTPSClient(testCrt)
	if err != nil {
		t.Error(err)
	}

	expected := time.Second * 30
	if hc.Timeout != expected {
		t.Errorf("Timeout expected: %v, got: %v", expected, hc.Timeout)
	}
}

func TestGetGRPCServerCredentials(t *testing.T) {
	cr, err := GetGRPCServerCredentials(testCrt, testKey)
	if err != nil {
		t.Error(err)
	}

	switch cr.(type) {
	case credentials.TransportCredentials:
		break
	default:
		t.Errorf("Invalid return type")
	}
}

func TestGetGRPCClientCredentials(t *testing.T) {
	cr, err := GetGRPCClientCredentials(testCrt)
	if err != nil {
		t.Error(err)
	}

	switch cr.(type) {
	case credentials.TransportCredentials:
		break
	default:
		t.Errorf("Invalid return type")
	}
}
