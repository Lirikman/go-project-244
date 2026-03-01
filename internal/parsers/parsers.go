package parsers

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

// определяем карту для хранения данных
var parsData map[string]map[string]any

// функция парсинга файлов .json в переменную parsData
func ReadJson(path string, jsonByte []byte) error {
	// получаем имя файла
	fileName := filepath.Base(path)
	// создаём временную карту для парсинга одного файла
	var tempData map[string]any
	//  десериализуем json файл во временную карту
	err := json.Unmarshal(jsonByte, &tempData)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}
	// инициализируем карту хранения данных
	parsData = make(map[string]map[string]any)
	if parsData[fileName] == nil {
		parsData[fileName] = make(map[string]any)
	}
	// добавляем данные в карту хранения данных
	parsData[fileName] = tempData
	return nil
}

// функция парсинга файлов .yml в переменную parsData
func ReadYaml(path string, yamlByte []byte) error {
	// получаем имя файла
	fileName := filepath.Base(path)
	// создаём временную карту для парсинга одного файла
	var tempData map[string]any
	//  десериализуем yaml файл во временную карту
	err := yaml.Unmarshal(yamlByte, &parsData)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}
	// инициализируем карту хранения данных
	parsData = make(map[string]map[string]any)
	if parsData[fileName] == nil {
		parsData[fileName] = make(map[string]any)
	}
	// добавляем данные в карту хранения данных
	parsData[fileName] = tempData
	return nil
}

// функция чтения и парсинга файлов .json и .yaml
func ReadData(path string) error {
	var res error
	data, _ := ReadFiles(path)
	// проверяем что расширение файла json
	if strings.HasSuffix(path, ".json") {
		res = ReadJson(path, data)
		if res == nil {
			for name, allData := range parsData {
				fmt.Printf("Файл: %s\n", name)
				for k, v := range allData {
					fmt.Printf("key:%v\tvalue:%v\n", k, v)
				}
			}
		}
	}
	// проверяем что расширение файла yml
	if strings.HasSuffix(path, ".yml") {
		res = ReadYaml(path, data)
		if res == nil {
			for name, allData := range parsData {
				fmt.Printf("Файл: %s\n", name)
				for k, v := range allData {
					fmt.Printf("key:%v\tvalue:%v\n", k, v)
				}
			}
		}
	}
	return res
}
