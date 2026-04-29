package code

import (
	"fmt"
	"reflect"
	"slices"
	"strings"
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

// функция построения дерева различий
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
	// переменная для хранения дерева различий
	var diff map[string]map[string]any
	// инициализируем карту
	diff = make(map[string]map[string]any)
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
				diff[nameKey]["children"] = GenDiff(m1, m2)
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

// функция форматирования ответа в stylish
func FormatterStylish(tree map[string]map[string]any) string {
	var builder strings.Builder
	builder.WriteString("{\n")

	// анонимная функция для рекурсивного обхода вложенной карты
	var walkMap func(map[string]any, int)
	walkMap = func(m map[string]any, d int) {
		indent := d * 4
		for key, val := range m {
			if nestedMap, ok := val.(map[string]any); ok {
				builder.WriteString(strings.Repeat(".", indent))
				fmt.Fprintf(&builder, "%v: {\n", key)
				walkMap(nestedMap, d+1)
			} else {
				builder.WriteString(strings.Repeat(".", indent))
				fmt.Fprintf(&builder, "%v: %v\n", key, val)
			}
		}
		builder.WriteString(strings.Repeat(".", indent-4))
		builder.WriteString("}\n")
	}

	// анонимная функция для рекурсивного обхода дерева различий
	var recChild func(map[string]map[string]any, int)
	recChild = func(data map[string]map[string]any, depth int) {
		//получим все ключи текущего уровня
		var allKeys []string
		for key := range data {
			allKeys = append(allKeys, key)
		}
		// отсортируем ключи
		slices.Sort(allKeys)

		// проходим по всем ключам добавляем их и отступы в сообщение
		for _, nameKey := range allKeys {
			// проходим по ключам проверям тип, ставим отступы и нужный знак
			for k, v := range data[nameKey] {
				// рассчитываем отступ на текущем уровне
				indent := depth*4 - 2
				// если типа ветки 'added'
				if k == "type" && v == "added" {
					builder.WriteString(strings.Repeat(".", indent))
					if n, ok := data[nameKey]["value2"].(map[string]any); ok {
						fmt.Fprintf(&builder, "+ %v: {\n", nameKey)
						walkMap(n, depth+1)
					} else {
						fmt.Fprintf(&builder, "+ %v: %v", nameKey, data[nameKey]["value2"])
					}
				}
				// если тип ветки 'deleted'
				if k == "type" && v == "deleted" {
					builder.WriteString(strings.Repeat(".", indent))
					if n, ok := data[nameKey]["value1"].(map[string]any); ok {
						fmt.Fprintf(&builder, "- %v: {\n", nameKey)
						walkMap(n, depth+1)
					} else {
						fmt.Fprintf(&builder, "- %v: %v", nameKey, data[nameKey]["value1"])
					}
				}
				// если тип ветки 'unchanged'
				if k == "type" && v == "unchanged" {
					builder.WriteString(strings.Repeat(".", indent))
					if n, ok := data[nameKey]["value1"].(map[string]any); ok {
						fmt.Fprintf(&builder, "  %v: {\n", nameKey)
						walkMap(n, depth+1)
					} else {
						fmt.Fprintf(&builder, "  %v: %v", nameKey, data[nameKey]["value1"])
					}
				}
				// если тип ветки 'changed'
				if k == "type" && v == "changed" {
					builder.WriteString(strings.Repeat(".", indent))
					if n, ok := data[nameKey]["value1"].(map[string]any); ok {
						fmt.Fprintf(&builder, "- %v: {\n", nameKey)
						walkMap(n, depth+1)
					} else {
						fmt.Fprintf(&builder, "- %v: %v\n", nameKey, data[nameKey]["value1"])
					}
					builder.WriteString(strings.Repeat(".", indent))
					if n, ok := data[nameKey]["value2"].(map[string]any); ok {
						fmt.Fprintf(&builder, "+ %v: {\n", nameKey)
						walkMap(n, depth+1)
					} else {
						fmt.Fprintf(&builder, "+ %v: %v", nameKey, data[nameKey]["value2"])
					}
				}
				// если тип ветки 'nested'
				if k == "type" && v == "nested" {
					builder.WriteString(strings.Repeat(".", indent))
					fmt.Fprintf(&builder, "  %v: {\n", nameKey)
					if child, ok := data[nameKey]["children"].(map[string]map[string]any); ok {
						recChild(child, depth+1)
					}
					builder.WriteString(strings.Repeat(".", indent))
					builder.WriteString("  }")
				}
			}
			builder.WriteString("\n")
		}
	}
	// запускаем рекурсивный обход дерева
	recChild(tree, 1)

	builder.WriteString("}")
	return builder.String()
}
