package main

import (
	"log"
	"os"

	"github.com/lusingander/monkey/command"
	"github.com/urfave/cli/v2"
)

func run(args []string) error {
	app := &cli.App{
		Name:  "monkey",
		Usage: "CLI tool for Monkey programming language",
		Commands: []*cli.Command{
			command.ReplCommand,
			command.RunCommand,
		},
	}
	return app.Run(args)
}

func main() {
	if err := run(os.Args); err != nil {
		log.Fatal(err)
	}
}
