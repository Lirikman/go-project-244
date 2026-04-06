package main

import (
	"code"
	parser "code/internal/parsers"
	"context"
	"fmt"
	"log"
	"os"

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
			//fmt.Println(data)
			data1, data2 := code.SplitNestedMap(data)
			fmt.Println(code.GenDiff(data1, data2))
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
