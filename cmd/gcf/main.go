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
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/suzuito/blog1-go/deployment/gcf"
	"github.com/suzuito/blog1-go/pkg/inject"
	"github.com/suzuito/blog1-go/pkg/setting"
)

type testData struct {
	Event    gcf.GCSEvent      `json:"event"`
	Metadata metadata.Metadata `json:"metadata"`
}

func readDirTest(dir string, f func(d *testData) error) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return errors.Wrapf(err, "cannot read dir %s", dir)
	}
	for _, entry := range entries {
		filePath := path.Join(dir, entry.Name())
		eventBytes, err := ioutil.ReadFile(filePath)
		if err != nil {
			return errors.Wrapf(err, "cannot read file %s", filePath)
		}
		d := testData{}
		if err := json.Unmarshal(eventBytes, &d); err != nil {
			return errors.Wrapf(err, "unmarshal %s", filePath)
		}
		if err := f(&d); err != nil {
			return errors.Wrapf(err, "Calling f is failed")
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
