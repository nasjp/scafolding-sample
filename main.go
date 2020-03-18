package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(0)
}

func run() error {
	f, err := os.Open("sample.json")
	if err != nil {
		return err
	}
	defer f.Close()
	m := map[string]interface{}{}
	if err := json.NewDecoder(f).Decode(&m); err != nil {
		return err
	}

	s := parser(m, 0, false, true)

	fmt.Println(s)
	return nil
}

var increment int

func parser(i interface{}, nest int, isSameLine bool, isLastElement bool) string {
	switch i.(type) {
	case string:
		return fmt.Sprintf("%s\"%s\"%s\n", getSpaces(nest, isSameLine), i.(string), getComma(isLastElement))
	case float64:
		return fmt.Sprintf("%s%s%s\n", getSpaces(nest, isSameLine), numberPerser(i.(float64)), getComma(isLastElement))
	case bool:
		return fmt.Sprintf("%s%t%s\n", getSpaces(nest, isSameLine), i.(bool), getComma(isLastElement))
	case nil:
		return fmt.Sprintf("%s%s%s\n", getSpaces(nest, isSameLine), "<nil>", getComma(isLastElement))
	case []interface{}:
		txt := fmt.Sprintf("%s%s\n", getSpaces(nest, isSameLine), "[")
		s := i.([]interface{})
		length := len(s)
		for idx, v := range s {
			idx++
			txt += parser(v, nest+1, false, checkLast(idx, length))
		}
		txt += fmt.Sprintf("%s%s%s\n", getSpaces(nest, false), "]", getComma(isLastElement))
		return txt
	case map[string]interface{}:
		txt := fmt.Sprintf("%s%s\n", getSpaces(nest, isSameLine), "{")
		m := i.(map[string]interface{})
		length := len(m)
		var idx int
		for k, v := range m {
			idx++
			txt += fmt.Sprintf("%s%s:", getSpaces(nest+1, false), k)
			txt += parser(v, nest+1, true, checkLast(idx, length))
		}
		txt += fmt.Sprintf("%s%s%s\n", getSpaces(nest, false), "}", getComma(isLastElement))
		return txt
	}
	return ""
}

func numberPerser(f float64) string {
	if math.Floor(f) == f {
		return fmt.Sprintf("%d", int64(f))
	}
	return fmt.Sprintf("%f", f)
}

func getSpaces(nest int, isSameLine bool) string {
	spaces := strings.Repeat("  ", nest)
	if isSameLine {
		spaces = " "
	}
	return spaces
}

func getComma(isLastElement bool) string {
	if isLastElement {
		return ""
	}
	return ","
}

func checkLast(idx, length int) bool {
	return idx == length
}
