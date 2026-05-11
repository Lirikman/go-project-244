package main

import (
	"code"
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
				Usage:   "supported formats: stylish, plain, json",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			// проверка корректности путей к файлам
			_, errPath1 := os.Stat(path1)
			_, errPath2 := os.Stat(path2)

			if errPath1 != nil || errPath2 != nil {
				return fmt.Errorf("incorrect paths entered")
			}
			if path1 == "" || path2 == "" {
				return fmt.Errorf("exactly two arguments are required")
			}
			// построение ответа в выбранном формате (по умолчанию - "stylish")
			res, err := code.GenDiff(path1, path2, cmd.String("format"))
			if err == nil {
				fmt.Println(res)
			}
			return err
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
