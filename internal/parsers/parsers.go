package parsers

import (
	"encoding/json"
	"errors"
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

// определяем вложенную карту для хранения данных
var parsData map[int]map[string]any

// определяем счётчик для сохранения порядка документов
var count int = 0

// функция парсинга файлов .json в переменную parsData
func ReadJson(path string, jsonByte []byte) error {
	// создаём временную карту для десериализации json
	var tempMap map[string]any
	//  десериализуем json файл в карту
	err := json.Unmarshal(jsonByte, &tempMap)
	if err != nil {
		return errors.New("unable to deserialize JSON file")
	}
	// если карта пустая, то инициализируем её
	if parsData == nil {
		parsData = make(map[int]map[string]any)
	}
	//добавляем данные в основную карту
	parsData[count] = tempMap
	//прибавляем счётчик
	count++
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
	// если карта пустая, то инициализируем её
	if parsData == nil {
		parsData = make(map[int]map[string]any)
	}
	//добавляем данные в основную карту
	parsData[count] = tempMap
	//прибавляем счётчик
	count++
	return nil
}

// функция чтения и парсинга файлов .json и .yaml
func ReadData(path string) (map[int]map[string]any, error) {
	var res error
	data, err := ReadFiles(path)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	// проверяем расширение файла
	switch {
	// если расширение .json
	case strings.HasSuffix(path, ".json"):
		res = ReadJson(path, data)
		if res != nil {
			return nil, fmt.Errorf("%w", res)
		}
	// если расширение .yml
	case strings.HasSuffix(path, ".yml"):
		res = ReadYaml(path, data)
		if res != nil {
			return nil, fmt.Errorf("%w", res)
		}
	// если расширение .yaml
	case strings.HasSuffix(path, ".yaml"):
		res = ReadYaml(path, data)
		if res != nil {
			return nil, fmt.Errorf("%w", res)
		}
	// если другое расширение
	default:
		return parsData, fmt.Errorf("unsupported file extension")
	}
	return parsData, nil
}
