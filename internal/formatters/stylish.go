package formatters

import (
	"fmt"
	"slices"
	"strings"
)

func stringifyNil(value any) string {
	switch v := value.(type) {
	case nil:
		return "null"
	default:
		return fmt.Sprintf("%v", v)
	}
}

// функция форматирования ответа в stylish
func FormatterStylish(tree map[string]map[string]any) string {
	var builder strings.Builder
	builder.WriteString("{\n")

	// анонимная функция для рекурсивного обхода вложенной карты
	var walkMap func(map[string]any, int)
	walkMap = func(m map[string]any, d int) {
		indent := d * 4
		//получим все ключи текущего уровня
		var allKeys []string
		for key := range m {
			allKeys = append(allKeys, key)
		}
		// отсортируем ключи
		slices.Sort(allKeys)
		// проходим по всем ключам и добавляем в сообщение
		for _, namekey := range allKeys {
			if nestedMap, ok := m[namekey].(map[string]any); ok {
				builder.WriteString(strings.Repeat(" ", indent))
				fmt.Fprintf(&builder, "%v: {\n", namekey)
				walkMap(nestedMap, d+1)
			} else {
				builder.WriteString(strings.Repeat(" ", indent))
				fmt.Fprintf(&builder, "%v: %v\n", namekey, stringifyNil(m[namekey]))
			}
		}
		builder.WriteString(strings.Repeat(" ", indent-4))
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
					builder.WriteString(strings.Repeat(" ", indent))
					if n, ok := data[nameKey]["value2"].(map[string]any); ok {
						fmt.Fprintf(&builder, "+ %v: {\n", nameKey)
						walkMap(n, depth+1)
					} else {
						fmt.Fprintf(&builder, "+ %v: %v", nameKey, stringifyNil(data[nameKey]["value2"]))
					}
				}
				// если тип ветки 'deleted'
				if k == "type" && v == "deleted" {
					builder.WriteString(strings.Repeat(" ", indent))
					if n, ok := data[nameKey]["value1"].(map[string]any); ok {
						fmt.Fprintf(&builder, "- %v: {\n", nameKey)
						walkMap(n, depth+1)
					} else {
						fmt.Fprintf(&builder, "- %v: %v", nameKey, stringifyNil(data[nameKey]["value1"]))
					}
				}
				// если тип ветки 'unchanged'
				if k == "type" && v == "unchanged" {
					builder.WriteString(strings.Repeat(" ", indent))
					if n, ok := data[nameKey]["value1"].(map[string]any); ok {
						fmt.Fprintf(&builder, "  %v: {\n", nameKey)
						walkMap(n, depth+1)
					} else {
						fmt.Fprintf(&builder, "  %v: %v", nameKey, stringifyNil(data[nameKey]["value1"]))
					}
				}
				// если тип ветки 'changed'
				if k == "type" && v == "changed" {
					builder.WriteString(strings.Repeat(" ", indent))
					if n, ok := data[nameKey]["value1"].(map[string]any); ok {
						fmt.Fprintf(&builder, "- %v: {\n", nameKey)
						walkMap(n, depth+1)
					} else {
						fmt.Fprintf(&builder, "- %v: %v\n", nameKey, stringifyNil(data[nameKey]["value1"]))
					}
					builder.WriteString(strings.Repeat(" ", indent))
					if n, ok := data[nameKey]["value2"].(map[string]any); ok {
						fmt.Fprintf(&builder, "+ %v: {\n", nameKey)
						walkMap(n, depth+1)
					} else {
						fmt.Fprintf(&builder, "+ %v: %v", nameKey, stringifyNil(data[nameKey]["value2"]))
					}
				}
				// если тип ветки 'nested'
				if k == "type" && v == "nested" {
					builder.WriteString(strings.Repeat(" ", indent))
					fmt.Fprintf(&builder, "  %v: {\n", nameKey)
					if child, ok := data[nameKey]["children"].(map[string]map[string]any); ok {
						recChild(child, depth+1)
					}
					builder.WriteString(strings.Repeat(" ", indent))
					builder.WriteString("  }")
				}
			}
			if !strings.HasSuffix(builder.String(), "\n") {
				builder.WriteString("\n")
			}
		}
	}
	// запускаем рекурсивный обход дерева
	recChild(tree, 1)

	builder.WriteString("}")
	return builder.String()
}
