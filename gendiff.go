package code

import (
	"fmt"
	"slices"
)

// функция рекурсивного обхода карт
func walkMap(m map[string]any, depth int) {
	for k, v := range m {
		switch value := v.(type) {
		case map[string]any:
			fmt.Printf("%d. %s\n", depth, k)
			walkMap(value, depth+1)
		default:
			fmt.Printf("%d. %s: %v\n", depth, k, v)
		}
	}
}

// функция сравнения файлов
func GenDiff(dataMap map[string]map[string]any) error {
	// переменная для хранения имён файлов
	var allNames []string
	// переменная для хранения уникальных ключей текущего уровня
	var allUniqueKeys []string
	// получаем имена файлов и уникальные ключи текущего уровня вложенности
	for filename, data := range dataMap {
		allNames = append(allNames, filename)
		for key := range data {
			if !slices.Contains(allUniqueKeys, key) {
				allUniqueKeys = append(allUniqueKeys, key)
			}
		}
	}

	// сортируем ключи по возрастанию
	slices.Sort(allUniqueKeys)
	// переменная для хранения дерева различий
	var diff map[string]map[string]any
	// проходим в цикле по обеим картам
	for _, nameKey := range allUniqueKeys {
		// получаем значения ключей и их наличие
		val1, ok1 := dataMap[allNames[0]][nameKey]
		val2, ok2 := dataMap[allNames[1]][nameKey]
		// значение отсутствует в первой карте
		if !ok1 {
			if diff == nil {
				// если карта пустая, то инициализируем её
				diff = make(map[string]map[string]any)
			}
			diff[nameKey]["type"] = "added"
			diff[nameKey]["value"] = val2
		}
		// 	значение отсутствует во второй карте
		if !ok2 {
			if diff == nil {
				// если карта пустая, то инициализируем её
				diff = make(map[string]map[string]any)
			}
			diff[nameKey]["type"] = "deleted"
			diff[nameKey]["value"] = val1
		}
		// значения присутствуют в обеих картах
		if ok1 && ok2 {
			// проверяем тип обеих значений
			type1 := fmt.Sprintf("%T", val1)
			type2 := fmt.Sprintf("%T", val2)
			fmt.Println(type1)
			fmt.Println(type2)
		}
	}

	return nil
}
