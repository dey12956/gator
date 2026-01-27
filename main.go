package main

import (
	"log"
	"os"

	"github.com/dey12956/gator/internal/cli"
	"github.com/dey12956/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	newState := cli.State{
		C: &cfg,
	}

	cmds := cli.Commands{
		MapOfCommands: make(map[string]func(*cli.State, cli.Command) error),
	}

	cmds.Register("login", cli.HandlerLogin)

	cliArgs := os.Args
	if len(cliArgs) < 2 {
		log.Fatal("no command")
	}

	cmdName := cliArgs[1]
	args := cliArgs[2:]

	cmd := cli.Command{
		Name: cmdName,
		Args: args,
	}

	err = cmds.Run(&newState, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
