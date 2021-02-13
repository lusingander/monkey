package command

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"

	"github.com/lusingander/monkey/repl"
	"github.com/urfave/cli/v2"
)

var in, out = os.Stdin, os.Stdout

var ReplCommand = &cli.Command{
	Name:  "repl",
	Usage: "Start REPL",
	Action: func(c *cli.Context) error {
		user, err := user.Current()
		if err != nil {
			return err
		}

		fmt.Fprintf(out, "Hello %s! This is the Monkey programming language!\n", user.Username)
		fmt.Fprintf(out, "Feel free to type in commands\n")

		repl.Start(in, out)
		return nil
	},
}

var RunCommand = &cli.Command{
	Name:  "run",
	Usage: "Run Monkey program",
	Action: func(c *cli.Context) error {
		if c.NArg() != 1 {
			return cli.Exit("File not specified", 1)
		}
		filename := c.Args().First()
		content, err := ioutil.ReadFile(filename)
		if err != nil {
			return err
		}
		return run(string(content))
	},
}
