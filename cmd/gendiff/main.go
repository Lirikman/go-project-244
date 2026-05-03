package main

import (
	"code"
	formatters "code/internal/formatters"
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
				Name:    "format",
				Aliases: []string{"f"},
				Value:   "stylish",
				Usage:   "supported formats: stylish, plain",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
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
			data, err := parser.ReadData(path1)
			if err != nil {
				return err
			}
			data, err = parser.ReadData(path2)
			if err != nil {
				return err
			}
			// разделение на две карты
			data1, data2 := code.SplitNestedMap(data)
			// очистка данных парсинга
			clear(data)
			// построение дерева отличий
			deffTree := code.GenDiff(data1, data2)
			// построение ответа в выбранном формате (по умолчанию - "stylish")
			fmt.Println(formatters.FormatMessage(deffTree, cmd.String("format")))
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
