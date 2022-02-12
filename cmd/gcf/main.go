package main

import (
	"context"
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"path"

	"cloud.google.com/go/functions/metadata"
	"github.com/google/subcommands"
	"github.com/rs/zerolog/log"
	"github.com/suzuito/blog1-go/deployment/gcf"
	"github.com/suzuito/blog1-go/pkg/inject"
	"github.com/suzuito/blog1-go/pkg/setting"
	"golang.org/x/xerrors"
)

type testData struct {
	Event    gcf.GCSEvent      `json:"event"`
	Metadata metadata.Metadata `json:"metadata"`
}

func readDirTest(dir string, f func(d *testData) error) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		xerrors.Errorf("cannot read dir %s : %w", dir, err)
	}
	for _, entry := range entries {
		eventBytes, err := ioutil.ReadFile(path.Join(dir, entry.Name()))
		if err != nil {
			return xerrors.Errorf("cannot read file : %w", err)
		}
		d := testData{}
		if err := json.Unmarshal(eventBytes, &d); err != nil {
			return xerrors.Errorf("unmarshal : %w", err)
		}
		if err := f(&d); err != nil {
			return xerrors.Errorf("f : %w", err)
		}
	}
	return nil
}

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
	subcommands.Register(newRunBlogDeleteArticleCmd(gdeps, env), "")
	flag.Parse()
	os.Exit(int(subcommands.Execute(ctx)))
}
