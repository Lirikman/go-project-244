package code

import (
	"fmt"
	"slices"
)

// функция сравнения json файлов
func GenDiff(dataFile map[string]map[string]any) string {
	// переменная для сообщения
	var res string
	// переменная для хранения всех уникальных ключей
	var allUniqueKeys []string
	// переменная для хранения имён файлов
	var allNames []string
	// получаем все уникальные ключи из вложенной карты
	for name, dataMap := range dataFile {
		allNames = append(allNames, name)
		for key := range dataMap {
			if !slices.Contains(allUniqueKeys, key) {
				allUniqueKeys = append(allUniqueKeys, key)
			}
		}
	}
	// сортируем ключи по возрастанию
	slices.Sort(allUniqueKeys)
	// fmt.Println(allUniqueKeys)
	// проходим по всем ключам и сравниваем
	// формируем сообщение
	res += "{\n"
	for _, keyName := range allUniqueKeys {
		// получаем ключи и наличие в карте
		val1, ok1 := dataFile[allNames[0]][keyName]
		val2, ok2 := dataFile[allNames[1]][keyName]
		// проверяем условия и фомируем сообщение
		// наличие в обеих картах и значения совпадают
		if ok1 && ok2 && val1 == val2 {
			res += fmt.Sprintf("    %s: %v\n", keyName, val1)
			// наличие в обеих картах и значения не совпадают
		} else if ok1 && ok2 && val1 != val2 {
			res += fmt.Sprintf("  - %s: %v\n", keyName, val1)
			res += fmt.Sprintf("  + %s: %v\n", keyName, val2)
			// наличие ключа только в первой карте
		} else if ok1 && !ok2 {
			res += fmt.Sprintf("  - %s: %v\n", keyName, val1)
			// наличие ключа только во второй карте
		} else if !ok1 && ok2 {
			res += fmt.Sprintf("  + %s: %v\n", keyName, val2)
		}
	}
	res += ("}\n")
	return res
}
