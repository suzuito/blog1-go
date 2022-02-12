package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/google/subcommands"
	"github.com/suzuito/blog1-go/deployment/gcf"
	"github.com/suzuito/blog1-go/pkg/inject"
	"github.com/suzuito/blog1-go/pkg/setting"
)

type runBlogUpdateArticleCmd struct {
	gdeps   *inject.GlobalDepends
	env     *setting.Environment
	dirBase string
}

func newRunBlogUpdateArticleCmd(gdeps *inject.GlobalDepends, env *setting.Environment) *runBlogUpdateArticleCmd {
	return &runBlogUpdateArticleCmd{gdeps: gdeps, env: env}
}

func (c *runBlogUpdateArticleCmd) Name() string { return "update-article" }
func (c *runBlogUpdateArticleCmd) Synopsis() string {
	return "記事を更新する\n"
}
func (c *runBlogUpdateArticleCmd) Usage() string {
	return c.Synopsis()
}

func (c *runBlogUpdateArticleCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.dirBase, "input-dir", "", "")
}

func (c *runBlogUpdateArticleCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	entries, err := os.ReadDir(c.dirBase)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot read dir %s : %+v\n", c.dirBase, err)
		return subcommands.ExitFailure
	}
	for _, entry := range entries {
		eventBytes, err := ioutil.ReadFile(path.Join(c.dirBase, entry.Name()))
		if err != nil {
			fmt.Fprintf(os.Stderr, "cannot read file : %+v\n", err)
			return subcommands.ExitFailure
		}
		event := gcf.GCSEvent{}
		if err := json.Unmarshal(eventBytes, &event); err != nil {
			fmt.Fprintf(os.Stderr, "unmarshal : %+v\n", err)
			return subcommands.ExitFailure
		}
		if err := gcf.BlogUpdateArticle(ctx, event); err != nil {
			fmt.Fprintf(os.Stderr, ": %+v\n", err)
			return subcommands.ExitFailure
		}
	}
	return subcommands.ExitSuccess
}
