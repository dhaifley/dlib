package dlib

import "strings"

// XMLEncode encodes an xml string into a string with xml special characters.
func XMLEncode(xmlstr string) string {
	xmlstr = strings.Replace(xmlstr, "&", "&amp;", -1)
	xmlstr = strings.Replace(xmlstr, "<", "&lt;", -1)
	xmlstr = strings.Replace(xmlstr, ">", "&gt;", -1)
	xmlstr = strings.Replace(xmlstr, "\"", "&quot;", -1)
	xmlstr = strings.Replace(xmlstr, "'", "&apos;", -1)
	return xmlstr
}

// XMLDecode decodes a string with xml special characters into an xml string.
func XMLDecode(xmlstr string) string {
	xmlstr = strings.Replace(xmlstr, "&amp;", "&", -1)
	xmlstr = strings.Replace(xmlstr, "&lt;", "<", -1)
	xmlstr = strings.Replace(xmlstr, "&gt;", ">", -1)
	xmlstr = strings.Replace(xmlstr, "&quot;", "\"", -1)
	xmlstr = strings.Replace(xmlstr, "&apos;", "'", -1)
	return xmlstr
}

// PadString pads a string to the specified length with the
// specified pad string.
func PadString(str string, length int, pad string) string {
	res := ""
	for i := 0; i < length; i++ {
		res += pad
	}

	res += str
	return res[len(res)-length:]
}

// FormatUPC formats a string as a valid UPC number optionally trimming
// the check digit.
func FormatUPC(upc string, length int, check bool) string {
	if len(strings.TrimLeft(upc, " 0")) <= 5 {
		return upc
	}

	if check {
		return PadString(upc, length, "0")
	}

	return PadString(upc[0:len(upc)-1], length, "0")
}
