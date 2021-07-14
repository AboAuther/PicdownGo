package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func checkFileIsExits(jsonFileAddr string) bool {
	if _, err := os.Stat(jsonFileAddr); err == nil {
		return true
	}
	return false
}

func reloadParser(jsonFileAddr string) (*Parser, error) {
	var file []byte
	var configuration Parser
	if checkFileIsExits(jsonFileAddr) {
		file, _ = os.ReadFile(jsonFileAddr)
		fmt.Println("Loading local json file...")
	} else {
		err := errors.New("json file is not Existed")
		return nil, err
	}
	if len(file) == 0 {
		fmt.Println("Json file is empty,Loading default parser...")
		file = DefaultJson
	}
	err := json.Unmarshal(file, &configuration)
	if err != nil {
		return nil, fmt.Errorf("json file unmarshal failed,%w", err)
	}
	return &configuration, nil
}
