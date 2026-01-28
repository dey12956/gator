package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/dey12956/gator/internal/cli"
	"github.com/dey12956/gator/internal/config"
	"github.com/dey12956/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)

	newState := cli.State{
		C:  &cfg,
		DB: dbQueries,
	}

	cmds := cli.Commands{
		MapOfCommands: make(map[string]func(*cli.State, cli.Command) error),
	}

	cmds.Register("login", cli.HandlerLogin)
	cmds.Register("register", cli.HandlerRegister)
	cmds.Register("reset", cli.HandlerReset)
	cmds.Register("users", cli.HandlerUsers)
	cmds.Register("agg", cli.HandlerAgg)
	cmds.Register("addfeed", cli.HandlerAddFeed)
	cmds.Register("feeds", cli.HandlerFeeds)

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
