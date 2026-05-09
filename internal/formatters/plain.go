package formatters

import (
	"fmt"
	"slices"
	"strings"
)

func stringify(value any) string {
	switch v := value.(type) {
	case nil:
		return "null"
	case string:
		return fmt.Sprintf("'%s'", v)
	case bool, int, int64, float64:
		return fmt.Sprintf("%v", v)
	case map[string]any, []any:
		return "[complex value]"
	default:
		return fmt.Sprintf("%v", v)
	}
}

func FormmaterPlain(tree map[string]map[string]any) string {
	var builder strings.Builder
	//получим все ключи корневого уровня
	var allKeys []string
	for key := range tree {
		allKeys = append(allKeys, key)
	}
	// отсортируем ключи
	slices.Sort(allKeys)

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
				// если тип ветки 'added'
				if key == "type" && val == "added" {
					fmt.Fprintf(&builder, "%v", name)
					fmt.Fprintf(&builder, "%v' was added with value: %v\n", nameKey, stringify(data[nameKey]["value2"]))
				}
				// если тип ветки 'deleted'
				if key == "type" && val == "deleted" {
					fmt.Fprintf(&builder, "%v", name)
					fmt.Fprintf(&builder, "%v' was removed\n", nameKey)
				}
				// если тип ветки 'changed'
				if key == "type" && val == "changed" {
					fmt.Fprintf(&builder, "%v", name)
					fmt.Fprintf(&builder, "%v' was updated. From %v to %v\n", nameKey, stringify(data[nameKey]["value1"]), stringify(data[nameKey]["value2"]))
				}
				// если тип ветки 'nested'
				// запускаем рекурсивный проход по значению
				if key == "type" && val == "nested" {
					nameNew := name + nameKey + "."
					if child, ok := data[nameKey]["children"].(map[string]map[string]any); ok {
						recChild(child, nameNew)
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
				fmt.Fprintf(&builder, "Property '%v' was added with value: %v\n", nameKey, stringify(tree[nameKey]["value2"]))
			}
			// если тип ветки 'deleted'
			if k == "type" && v == "deleted" {
				fmt.Fprintf(&builder, "Property '%v' was removed\n", nameKey)
			}
			// если тип ветки 'changed'
			if k == "type" && v == "changed" {
				fmt.Fprintf(&builder, "Property %v was updated. ", nameKey)
				fmt.Fprintf(&builder, "From %v to %v", stringify(tree[nameKey]["value1"]), stringify(tree[nameKey]["value2"]))
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
