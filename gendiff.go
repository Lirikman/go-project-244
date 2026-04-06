package code

import (
	"reflect"
	"slices"
)

// функция разделения вложенной карты на две отдельные карты
func SplitNestedMap(dataMap map[string]map[string]any) (map[string]any, map[string]any) {
	// переменная для хранения имён файлов
	var allNames []string
	// создаём и инициализируем карту для данных первого файла
	data1 := make(map[string]any)
	// создаём и инициализируем карту для данных второго файла
	data2 := make(map[string]any)
	// получаем имена файлов
	for filename := range dataMap {
		allNames = append(allNames, filename)
	}
	// заполняем карты
	data1 = dataMap[allNames[0]]
	data2 = dataMap[allNames[1]]
	return data1, data2
}

// функция сравнения файлов
func GenDiff(data1 map[string]any, data2 map[string]any) map[string]map[string]any {
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
	// сортируем ключи по возрастанию
	slices.Sort(allUniqueKeys)
	// переменная для хранения дерева различий
	var diff map[string]map[string]any
	// проходим в цикле по обеим картам
	for _, nameKey := range allUniqueKeys {
		// получаем значения ключей и их наличие
		val1, ok1 := data1[nameKey]
		val2, ok2 := data2[nameKey]
		// значение отсутствует в первой карте
		if !ok1 {
			if diff == nil {
				// если карта пустая, то инициализируем её
				diff = make(map[string]map[string]any)
			}
			diff[nameKey]["type"] = "added"
			diff[nameKey]["value2"] = val2
		}
		// 	значение отсутствует во второй карте
		if !ok2 {
			if diff == nil {
				// если карта пустая, то инициализируем её
				diff = make(map[string]map[string]any)
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
				diff[nameKey]["type"] = "nested"
				diff[nameKey]["children"] = GenDiff(m1, m2)
				// одно или оба значения не являются картами
				// сравниваем с помощью рефлексии
			} else {
				if reflect.DeepEqual(m1, m2) {
					diff[nameKey]["type"] = "unchanged"
					diff[nameKey]["value1"] = m1
				} else {
					diff[nameKey]["type"] = "changed"
					diff[nameKey]["value1"] = m1
					diff[nameKey]["value2"] = m2
				}
			}
		}
	}

	return diff
}
