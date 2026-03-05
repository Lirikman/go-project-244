package main

import (
	parser "code/internal/parsers"
	"context"
	"fmt"
	"log"
	"os"
	"slices"

	"github.com/urfave/cli/v3"
)

func main() {
	var path1 string
	var path2 string
	cmd := &cli.Command{
		Name:  "gendiff",
		Usage: "Compares two configuration files and shows a difference.",
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name:        "filepath1",
				Destination: &path1,
			},
			&cli.StringArg{
				Name:        "filepath2",
				Destination: &path2,
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "format string",
				Aliases: []string{"f"},
				Value:   "stylish",
				Usage:   "output format",
			},
		},
		Action: func(context.Context, *cli.Command) error {
			// проверка корректности путей к файлам
			_, errPath1 := os.Stat(path1)
			_, errPath2 := os.Stat(path2)

			if errPath1 != nil || errPath2 != nil {
				return fmt.Errorf("files at the entered path were not found")
			}
			if path1 == "" || path2 == "" {
				return fmt.Errorf("invalid files path")
			}

			// парсинг данных из файлов
			_, _ = parser.ReadData(path1)
			data, _ := parser.ReadData(path2)
			//for name, allData := range data {
			//	fmt.Printf("Файл: %s\n", name)
			//	for k, v := range allData {
			//		fmt.Printf("key:%v\tvalue:%v\n", k, v)
			//	}
			//}
			fmt.Println(genDiff(data))
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

// функция сравнения json файлов
func genDiff(dataFile map[string]map[string]any) string {
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
