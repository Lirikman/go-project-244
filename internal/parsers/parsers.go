package parsers

import (
	"encoding/json"
	"errors"
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

// определяем вложенную карту для хранения данных
var parsData map[string]map[string]any

// функция парсинга файлов .json в переменную parsData
func ReadJson(path string, jsonByte []byte) error {
	// создаём временную карту для десериализации json
	var tempMap map[string]any
	//  десериализуем json файл в карту
	err := json.Unmarshal(jsonByte, &tempMap)
	if err != nil {
		return errors.New("unable to deserialize JSON file")
	}
	// получаем имя файла
	fileName := filepath.Base(path)
	// если карта пустая, то инициализируем её
	if parsData == nil {
		parsData = make(map[string]map[string]any)
	}
	//добавляем данные в основную карту
	parsData[fileName] = tempMap
	return nil
}

// функция парсинга файлов .yml в переменную parsData
func ReadYaml(path string, yamlByte []byte) error {
	// создаём временную карту для десериализации yaml
	var tempMap map[string]any
	//  десериализуем json файл в карту
	err := yaml.Unmarshal(yamlByte, &tempMap)
	if err != nil {
		return errors.New("unable to deserialize YAML file")
	}
	// получаем имя файла
	fileName := filepath.Base(path)
	// если карта пустая, то инициализируем её
	if parsData == nil {
		parsData = make(map[string]map[string]any)
	}
	//добавляем данные в основную карту
	parsData[fileName] = tempMap
	return nil
}

// функция чтения и парсинга файлов .json и .yaml
func ReadData(path string) (map[string]map[string]any, error) {
	var res error
	data, err := ReadFiles(path)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	// проверяем что расширение файла json
	if strings.HasSuffix(path, ".json") {
		res = ReadJson(path, data)
		if res != nil {
			return nil, fmt.Errorf("%w", res)
		}
	}
	// проверяем что расширение файла yml
	if strings.HasSuffix(path, ".yml") || strings.HasSuffix(path, ".yaml") {
		res = ReadYaml(path, data)
		if res != nil {
			return nil, fmt.Errorf("%w", res)
		}
	}
	return parsData, nil
}
