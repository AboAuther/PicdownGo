package main

import (
	"encoding/json"
	"fmt"
	"log"
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
		log.Println("Loading local json file...")
	} else {
		log.Println("json file is not Existed,Loading default parser...")
		file = DefaultJson
	}
	if len(file) == 0 {
		log.Println("Json file is empty,Loading default parser...")
		file = DefaultJson
	}
	err := json.Unmarshal(file, &configuration)
	if err != nil {
		return nil, fmt.Errorf("json file unmarshal failed,%w", err)
	}
	return &configuration, nil
}
