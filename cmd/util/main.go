package main

import (
	"context"
	"flag"
	"os"

	"github.com/google/subcommands"
)

func main() {
	ctx := context.Background()
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(newRunConvertCmd(), "")
	flag.Parse()
	os.Exit(int(subcommands.Execute(ctx)))
}
