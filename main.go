package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"text/template"

	"github.com/nasjp/scafolding-sample/genport"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(0)
}

func run() error {
	jsonF, err := os.Open("get_user.json")
	if err != nil {
		return err
	}
	defer jsonF.Close()
	m := map[string]interface{}{}
	if err := json.NewDecoder(jsonF).Decode(&m); err != nil {
		return err
	}

	str := fmt.Sprintf("type %s %s", m["name"].(string), genport.MapToStruct(m["request"], 0))

	buf := bytes.NewBuffer([]byte{})

	tmpl, err := template.ParseFiles("./examples/template.go")
	if err != nil {
		return err
	}

	if err := tmpl.Execute(buf, map[string]string{"response": str}); err != nil {
		return err
	}

	fset := token.NewFileSet()
	astF, err := parser.ParseFile(fset, "", buf, parser.ParseComments)
	if err != nil {
		return err
	}

	goF, err := os.Create("examples/response.go")
	if err != nil {
		return err
	}
	defer goF.Close()

	if err := format.Node(goF, fset, astF); err != nil {
		return err
	}

	return nil
}
