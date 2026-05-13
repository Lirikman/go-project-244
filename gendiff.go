package code

import (
	formatters "code/internal/formatters"
	parser "code/internal/parsers"
	"errors"
	"reflect"
	"slices"
)

// функция разделения вложенной карты на две отдельные карты
func SplitNestedMap(dataMap map[int]map[string]any) (map[string]any, map[string]any, error) {
	errTwoArgEx := errors.New("exactly two arguments are required")
	// проверка на наличие всего двух записей
	if len(dataMap) != 2 {
		return map[string]any{}, map[string]any{}, errTwoArgEx
	}
	// переменная для хранения первого файла
	data1 := dataMap[0]
	// переменная для хранения второго файла
	data2 := dataMap[1]
	return data1, data2, nil
}

// функция построения дерева различий
func TreeBuildDiff(data1 map[string]any, data2 map[string]any) map[string]map[string]any {
	// переменная для хранения уникальных ключей текущего уровня
	var allUniqueKeys []string
	// получаем ключи с первой карты
	for key := range data1 {
		if !slices.Contains(allUniqueKeys, key) {
			allUniqueKeys = append(allUniqueKeys, key)
		}
	}
	// получаем ключи со второй карты
	for key := range data2 {
		if !slices.Contains(allUniqueKeys, key) {
			allUniqueKeys = append(allUniqueKeys, key)
		}
	}
	// создаём и инициализируем карту для хранения дерева различий
	diff := make(map[string]map[string]any)
	// проходим в цикле по обеим картам
	for _, nameKey := range allUniqueKeys {
		// получаем значения ключей и их наличие
		val1, ok1 := data1[nameKey]
		val2, ok2 := data2[nameKey]
		// значение отсутствует в первой карте
		if !ok1 {
			// если карта пустая, то инициализируем её
			if diff[nameKey] == nil {
				diff[nameKey] = make(map[string]any)
			}
			diff[nameKey]["type"] = "added"
			diff[nameKey]["value2"] = val2
		}
		// 	значение отсутствует во второй карте
		if !ok2 {
			// если карта пустая, то инициализируем её
			if diff[nameKey] == nil {
				diff[nameKey] = make(map[string]any)
			}
			diff[nameKey]["type"] = "deleted"
			diff[nameKey]["value1"] = val1
		}
		// значения присутствуют в обеих картах
		if ok1 && ok2 {
			// проверяем тип обеих значений
			m1, typeOk1 := val1.(map[string]any)
			m2, typeOk2 := val2.(map[string]any)
			// оба значения являются картами
			if typeOk1 && typeOk2 {
				// если карта пустая, то инициализируем её
				if diff[nameKey] == nil {
					diff[nameKey] = make(map[string]any)
				}
				diff[nameKey]["type"] = "nested"
				diff[nameKey]["children"] = TreeBuildDiff(m1, m2)
				// одно или оба значения не являются картами
			} else {
				// если карта пустая, то инициализируем её
				if diff[nameKey] == nil {
					diff[nameKey] = make(map[string]any)
				}
				// сравниваем с помощью рефлексии
				if reflect.DeepEqual(val1, val2) {
					diff[nameKey]["type"] = "unchanged"
					diff[nameKey]["value1"] = val1
				} else {
					diff[nameKey]["type"] = "changed"
					diff[nameKey]["value1"] = val1
					diff[nameKey]["value2"] = val2
				}
			}
		}
	}
	return diff
}

// функция генерации отличий по заданному формату
func GenDiff(filepath1, filepath2, formatName string) (string, error) {
	// парсинг данных из файлов
	_, err := parser.ReadData(filepath1)
	if err != nil {
		return "", err
	}
	data, err := parser.ReadData(filepath2)
	if err != nil {
		return "", err
	}
	// разделение на две картыcl
	data1, data2, err := SplitNestedMap(data)
	if err != nil {
		return "", errors.New("data partitioning error")
	}
	// построение дерева отличий
	deffTree := TreeBuildDiff(data1, data2)
	// вывод сообщения в выбранном формате
	result := formatters.FormatMessage(deffTree, formatName)
	return result, nil
}
