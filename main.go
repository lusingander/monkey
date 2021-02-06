package main

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/lusingander/monkey/repl"
)

var in, out = os.Stdin, os.Stdout

func run(args []string) error {
	user, err := user.Current()
	if err != nil {
		return err
	}

	fmt.Fprintf(out, "Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Fprintf(out, "Feel free to type in commands\n")

	repl.Start(in, out)
	return nil
}

func main() {
	if err := run(os.Args); err != nil {
		log.Fatal(err)
	}
}
