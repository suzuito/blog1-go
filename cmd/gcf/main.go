package main

import (
	"context"
	"flag"
	"os"

	"github.com/google/subcommands"
	"github.com/rs/zerolog/log"
	"github.com/suzuito/blog1-go/pkg/inject"
	"github.com/suzuito/blog1-go/pkg/setting"
)

func main() {
	ctx := context.Background()
	env, err := setting.NewEnvironment()
	if err != nil {
		log.Error().AnErr("message", err).Send()
		os.Exit(1)
	}
	gdeps, closeFunc, err := inject.NewGlobalDepends(ctx, env)
	if err != nil {
		log.Error().AnErr("message", err).Send()
		os.Exit(1)
	}
	defer closeFunc()
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(newRunBlogUpdateArticleCmd(gdeps, env), "")
	flag.Parse()
	os.Exit(int(subcommands.Execute(ctx)))
}
