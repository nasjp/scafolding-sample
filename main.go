package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/ChimeraCoder/gojson"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(0)
}

func run() error {
	// return part1()
	return part2()
}

func part1() error {
	f, err := os.Open("sample.json")
	if err != nil {
		return err
	}
	defer f.Close()

	b, err := gojson.Generate(f, gojson.ParseJson, strings.Split(f.Name(), ".")[0], "_", []string{"json"}, true, true)
	if err != nil {
		return err
	}
	sc := bufio.NewScanner(bytes.NewBuffer(b))
	for sc.Scan() {
		if sc.Text() == "" {
			continue
		}
		if strings.HasPrefix(sc.Text(), "package") {
			continue
		}
		fmt.Println(sc.Text())
	}
	return nil
}

func part2() error {
	// f, err := os.OpenFile("sample.json", os.O_RDWR, 0664)
	f, err := os.Open("sample.json")
	if err != nil {
		return err
	}
	defer f.Close()
	m := map[string]interface{}{}
	if err := json.NewDecoder(f).Decode(&m); err != nil {
		return err
	}

	s := parser(m, 0)

	fmt.Println(s)
	return nil
}

var indentStr = "  "

func parser(i interface{}, nest int) string {
	switch i.(type) {
	case string:
		return fmt.Sprintf("%s\"%s\",\n", strings.Repeat(indentStr, nest), i.(string))
	case float64:
		return fmt.Sprintf("%s%s,\n", strings.Repeat(indentStr, nest), numberPerser(i.(float64)))
	case bool:
		return fmt.Sprintf("%s%t,\n", strings.Repeat(indentStr, nest), i.(bool))
	case nil:
		return fmt.Sprintf("%s%s,\n", strings.Repeat(indentStr, nest), "<nil>")
	case []interface{}:
		s := i.([]interface{})
		txt := fmt.Sprintf("%s%s\n", strings.Repeat(indentStr, nest), "[")
		for _, v := range s {
			txt += parser(v, nest+1)
		}
		txt += fmt.Sprintf("%s%s,\n", strings.Repeat(indentStr, nest), "]")
		return txt
	case map[string]interface{}:
		m := i.(map[string]interface{})
		txt := fmt.Sprintf("%s%s\n", strings.Repeat(indentStr, nest), "{")
		for k, v := range m {
			txt += fmt.Sprintf("%s%s: ", strings.Repeat(indentStr, nest+1), k)
			txt += parser(v, nest+1)
		}
		txt += fmt.Sprintf("%s%s,\n", strings.Repeat(indentStr, nest), "}")
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

// func getSpace()
