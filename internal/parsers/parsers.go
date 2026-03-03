package parsers

import (
	"encoding/json"
	"fmt"
	"os"
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
var parsData map[string]any

// функция для очистки карты
func clearMap(m map[string]any) {
	for k := range m {
		delete(m, k)
	}
}

// функция парсинга файлов .json в переменную parsData
func ReadJson(path string, jsonByte []byte) error {
	// очищаем карту перед десериализацией json
	clearMap(parsData)
	//  десериализуем json файл в карту
	err := json.Unmarshal(jsonByte, &parsData)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}
	return nil
}

// функция парсинга файлов .yml в переменную parsData
func ReadYaml(path string, yamlByte []byte) error {
	// очищаем карту перед десериализацией нфьд
	clearMap(parsData)
	//  десериализуем yaml файл в карту
	err := yaml.Unmarshal(yamlByte, &parsData)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}
	return nil
}

// функция чтения и парсинга файлов .json и .yaml
func ReadData(path string) (map[string]any, error) {
	var res error
	data, _ := ReadFiles(path)
	// проверяем что расширение файла json
	if strings.HasSuffix(path, ".json") {
		res = ReadJson(path, data)
		if res != nil {
			return nil, fmt.Errorf("error: %w", res)
		}
	}
	// проверяем что расширение файла yml
	if strings.HasSuffix(path, ".yml") {
		res = ReadYaml(path, data)
		if res != nil {
			return nil, fmt.Errorf("error: %w", res)
		}
	}
	return parsData, nil
}
