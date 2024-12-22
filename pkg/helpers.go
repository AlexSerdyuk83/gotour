package pkg

import (
	"fmt"
	"strings"
)

func Sum(x int, y int) int {
	return x + y
}

func Swap(x, y string) (string, string) {
	return y, x
}

func DomainForLocale(domain, locale string) string {
	subdomain := ""

	if locale == "" {
		subdomain = "en"
	} else {
		subdomain = locale
	}

	return fmt.Sprintf("%s.%s", subdomain, domain)
}

func ModifySpaces(str, separator string) string {
	mod := ""

	switch separator {
	default:
		mod = "*"
	case "dash":
		mod = "-"
	case "underscore":
		mod = "_"
	}

	return strings.ReplaceAll(str, " ", mod)
}

const (
	OK        = 0
	CANCELLED = 1
	UNKNOWN   = 2
)

func ErrorMessageToCode(msg string) int {
	switch msg {
	default:
		return UNKNOWN
	case "OK":
		return OK
	case "CANCELLED":
		return CANCELLED
	}
}
