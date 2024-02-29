package strext

import "strings"

func PadLeft(str, prefix string, desiredLen int) string {
	if len(str) >= desiredLen {
		return str
	}

	return strings.Repeat(prefix, desiredLen/len(prefix)) + str
}

func PadBinary(str string, bits int) string {
	if len(str) >= bits {
		return str
	}

	return PadLeft(str, "0", bits)
}
