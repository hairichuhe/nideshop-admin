package tools

import (
	"strings"
)

func GetSuffix(s string) string {
	start := strings.LastIndex(s, ".")
	end := len(s)
	return SubStr(s, start, end)
}

func GetFileName(s string) string {
	start := strings.LastIndex(s, "/") + 1
	end := len(s)
	return SubStr(s, start, end)
}

func SubStr(s string, start, end int) string {
	rs := []byte(s)
	rl := len(rs)

	if start < 0 {
		start = 0
	}

	if start > end {
		start, end = end, start
	}

	if end > rl {
		end = rl
	}

	if start > rl {
		start = rl
	}

	if end < 0 {
		end = 0
	}

	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

func NTos(n string) string {
	result := ""
	for i := 0; i < len(n); i++ {
		result += nTs(string(n[i]))
	}
	return result
}

func STon(n string) string {
	result := ""
	for i := 0; i < len(n); i++ {
		result += sTn(string(n[i]))
	}
	return result
}

func nTs(n string) string {
	switch n {
	case "0":
		return "f"
	case "1":
		return "b"
	case "2":
		return "h"
	case "3":
		return "w"
	case "4":
		return "k"
	case "5":
		return "n"
	case "6":
		return "a"
	case "7":
		return "p"
	case "8":
		return "u"
	case "9":
		return "s"
	default:
		return ""
	}
}

func sTn(n string) string {
	switch n {
	case "f":
		return "0"
	case "b":
		return "1"
	case "h":
		return "2"
	case "w":
		return "3"
	case "k":
		return "4"
	case "n":
		return "5"
	case "a":
		return "6"
	case "p":
		return "7"
	case "u":
		return "8"
	case "s":
		return "9"
	default:
		return ""
	}
}
