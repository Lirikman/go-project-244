package formatters

import (
	"fmt"
	"slices"
	"strings"
)

func FormmaterPlain(tree map[string]map[string]any) string {
	var builder strings.Builder
	//получим все ключи текущего уровня
	var allKeys []string
	for key := range tree {
		allKeys = append(allKeys, key)
	}
	// отсортируем ключи
	slices.Sort(allKeys)
	// анонимная функция для рекурсивного прохода вложенной карты
	var walkMap func(map[string]any)
	walkMap = func(m map[string]any) {
		// проходим по всем ключам карты
		for key, val := range m {
			// если тип ключа 'added'
			if key == "type" && "val" == "added" {
				fmt.Fprintf(&builder, "' was added with value: %v\n", m["value2"])
				return
			}
			// если тип ключа 'deleted
			if key == "type" && val == "deleted" {
				fmt.Fprintf(&builder, "' was removed\n")
				return
			}
			// если тип ключа 'changed'
			if key == "type" && val == "changed" {
				fmt.Fprintf(&builder, "' was updated. From %v to %v\n", m["value1"], m["value2"])
				return
			}
			// если тип ключа 'nested'
			// запускаем рекурсивный проход по значению
			if key == "type" && val == "nested" {
				fmt.Fprintf(&builder, "%v", key)
				if n, ok := val.(map[string]any); ok {
					walkMap(n)
				}
			}
		}
	}

	// анонимная функция для рекурсивного обхода дерева различий
	var recChild func(map[string]map[string]any, string)
	recChild = func(data map[string]map[string]any, name string) {
		//получим все ключи текущего уровня
		var allKeys []string
		for key := range data {
			allKeys = append(allKeys, key)
		}
		// отсортируем ключи
		slices.Sort(allKeys)
		// проходим по всем ключам вложенной карты
		for _, nameKey := range allKeys {
			for key, val := range data[nameKey] {
				// если значение является картой то запускам функцию 'walkMap'
				if n, ok := val.(map[string]any); ok {
					walkMap(n)
				} else {
					// если тип ветки 'added'
					if key == "type" && val == "added" {
						fmt.Fprintf(&builder, "%v", name)
						// если значение явлется составным
						if _, ok := data[nameKey]["value2"].(map[string]any); ok {
							fmt.Fprintf(&builder, "%v' was added with value: [complex value]\n", nameKey)
							// если значение не является составным
						} else {
							fmt.Fprintf(&builder, "%v' was added with value: '%v'\n", nameKey, data[nameKey]["value2"])
						}
					}
					// если тип ветки 'deleted'
					if key == "type" && val == "deleted" {
						fmt.Fprintf(&builder, "%v", name)
						fmt.Fprintf(&builder, "%v' was removed\n", nameKey)
					}
					// если тип ветки 'changed'
					if key == "type" && val == "changed" {
						fmt.Fprintf(&builder, "%v", name)
						// если первое значение является составным
						if _, ok := data[nameKey]["value1"].(map[string]any); ok {
							fmt.Fprintf(&builder, "%v' was updated. From [complex value] to '%v'\n", nameKey, data[nameKey]["value2"])
							return
						}
						// если второе значение является составным
						if _, ok := data[nameKey]["value2"].(map[string]any); ok {
							fmt.Fprintf(&builder, "%v' was updated. From '%v' to [complex value]\n", nameKey, data[nameKey]["value1"])
							return
						}
						// если оба значения не являются составными
						fmt.Fprintf(&builder, "%v' was updated. From '%v' to '%v'\n", nameKey, data[nameKey]["value1"], data[nameKey]["value2"])
					}
					// если тип ветки 'nested'
					// запускаем рекурсивный проход по значению
					if key == "type" && val == "nested" {
						name = name + nameKey + "."
						if child, ok := data[nameKey]["children"].(map[string]map[string]any); ok {
							recChild(child, name)
						}
					}
				}
			}
		}
	}
	// проходим по всем корневым ключам
	for _, nameKey := range allKeys {
		// проходим по ключам проверям тип и наличие изменений
		for k, v := range tree[nameKey] {
			// если тип ветки 'added'
			if k == "type" && v == "added" {
				fmt.Fprintf(&builder, "Property '%v' was added with value: ", nameKey)
				// если значение является составным
				if _, ok := tree[nameKey]["value2"].(map[string]any); ok {
					fmt.Fprintf(&builder, "[complex value]\n")
					// если значение не является составным
				} else {
					fmt.Fprintf(&builder, "%v\n", tree[nameKey]["value2"])
				}
			}
			// если тип ветки 'deleted'
			if k == "type" && v == "deleted" {
				fmt.Fprintf(&builder, "Property '%v' was removed\n", nameKey)
			}
			// если тип ветки 'changed'
			if k == "type" && v == "changed" {
				fmt.Fprintf(&builder, "Property '%v' was updated. ", nameKey)
				// если первое значение является составным
				if _, ok := tree[nameKey]["value1"].(map[string]any); ok {
					fmt.Fprintf(&builder, "From [complex value] ")
					// если первое значение не является составным
				} else {
					fmt.Fprintf(&builder, "From '%v' ", tree[nameKey]["value1"])
				}
				// если второе значение является составным
				if _, ok := tree[nameKey]["value2"].(map[string]any); ok {
					fmt.Fprintf(&builder, "to [complex value]\n")
					// если второе значение не является составным
				} else {
					fmt.Fprintf(&builder, "to '%v'\n", tree[nameKey]["value2"])
				}
			}
			// если тип ветки 'nested'
			// запускаем рекурсивный проход по значению
			if k == "type" && v == "nested" {
				name := fmt.Sprintf("Property '%v.", nameKey)
				if child, ok := tree[nameKey]["children"].(map[string]map[string]any); ok {
					recChild(child, name)
				}
			}
		}
	}
	// возвращаем сообщение
	return builder.String()
}
