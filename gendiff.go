package code

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// функция чтения файлов
func ReadFiles(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return []byte{}, fmt.Errorf("file read error")
	}
	return data, nil
}

// переменная для хранения данных
var parsData map[string]any

// функция парсинга файлов .json в переменную parsData
func ReadJson(jsonByte []byte) error {
	err := json.Unmarshal(jsonByte, &parsData)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}
	return nil
}

// функция парсинга файлов .yaml в переменную
func ReadYaml(yamlByte []byte) error {
	err := yaml.Unmarshal(yamlByte, &parsData)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}
	return nil
}

// функция чтения и парсинга файлов .json и .yaml
func ReadData(path string) error {
	var res error
	fileName := filepath.Base(path)
	data, _ := ReadFiles(path)
	if strings.HasSuffix(path, ".json") {
		res = ReadJson(data)
		if res == nil {
			fmt.Printf("Файл: %s\n", fileName)
			for k, v := range parsData {
				fmt.Printf("key:%v\tvalue:%v\n", k, v)
			}
		}
	}
	if strings.HasSuffix(path, ".yml") {
		res = ReadYaml(data)
		if res == nil {
			fmt.Printf("Файл: %s\n", fileName)
			for k, v := range parsData {
				fmt.Printf("key:%v\tvalue:%v\n", k, v)
			}
		}
	}
	return res
}
