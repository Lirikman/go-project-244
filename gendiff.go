package code

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadFiles(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return []byte{}, fmt.Errorf("file read error")
	}
	os.Stdout.Write(data)
	return data, nil
}

var data map[string]interface{}

func ParsJson(data []byte) error {
	err := json.Unmarshal([]byte(data), &data)
	if err != nil {
		return fmt.Errorf("error parsing json:", err)
	}
	return nil
}
