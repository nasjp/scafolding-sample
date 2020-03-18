package genport

import (
	"fmt"
	"math"
	"strings"
)

// https://github.com/golang/go/wiki/CodeReviewComments#initialisms
var initialisms = map[string]bool{
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SSH":   true,
	"TLS":   true,
	"TTL":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
	"NTP":   true,
	"DB":    true,
}

func MapToStruct(i interface{}, nest int) string {
	switch i.(type) {
	case string:
		return "string"

	case float64:
		return getNumberType(i.(float64))

	case bool:
		return "bool"

	case nil:
		panic("NO expect null value")

	case []interface{}:
		s := i.([]interface{})
		if len(s) == 0 {
			panic("NO expect no value array")
		}
		return "[]" + MapToStruct(s[0], nest)

	case map[string]interface{}:
		m := i.(map[string]interface{})
		txt := "struct {\n"
		for k, v := range m {
			txt += fmt.Sprintf("%s%s %s `json:\"%s\"`\n", strings.Repeat("  ", nest+1), toCamelCase(k), MapToStruct(v, nest+1), k)
		}
		txt += strings.Repeat("  ", nest) + "}"
		return txt
	}
	return ""
}

func getNumberType(f float64) string {
	if math.Floor(f) == f {
		return "int64"
	}
	return "float64"
}

func toCamelCase(raw string) string {
	var (
		txt     string
		nowWord string
		// isToUpper bool
	)

	for i, r := range raw {
		// forst rune must be upper case
		if i == 0 {
			nowWord = strings.ToUpper(string(r))
			continue
		}

		// for snake case
		// if isToUpper {
		// 	nowWord += strings.ToUpper(string(r))
		// 	isToUpper = false
		// 	continue
		// }

		// for snake case
		// if r == '_' {
		// 	isToUpper = true
		// 	if upNowWord := strings.ToUpper(nowWord); initialisms[upNowWord] {
		// 		nowWord = upNowWord
		// 	}
		// 	txt += nowWord
		// 	nowWord = ""
		// 	continue
		// }

		// check initialisms
		if string(r) == strings.ToUpper(string(r)) {
			if upNowWord := strings.ToUpper(nowWord); initialisms[upNowWord] {
				nowWord = upNowWord
			}
			txt += nowWord
			nowWord = string(r)
			continue
		}

		// normal rune
		nowWord += string(r)
	}

	// check initialisms
	if upNowWord := strings.ToUpper(nowWord); initialisms[upNowWord] {
		nowWord = upNowWord
	}

	return txt + nowWord
}
