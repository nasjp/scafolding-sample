package main

import (
	"encoding/json"
	"fmt"
	"os"

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
	// f, err := os.Open("sample.json")
	f, err := os.Open("get_user.json")
	if err != nil {
		return err
	}
	defer f.Close()
	m := map[string]interface{}{}
	if err := json.NewDecoder(f).Decode(&m); err != nil {
		return err
	}

	// fmt.Println(toJsonString(m, 0, false, true))
	fmt.Printf("type Request%s ", m["name"])
	fmt.Println(genport.MapToStruct(m["request"], 0))
	return nil
}
