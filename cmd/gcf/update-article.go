package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/google/subcommands"
	"github.com/suzuito/blog1-go/deployment/gcf"
	"github.com/suzuito/blog1-go/pkg/inject"
	"github.com/suzuito/blog1-go/pkg/setting"
	"golang.org/x/xerrors"
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
	if err := readDirTest(c.dirBase, func(d *testData) error {
		if err := gcf.BlogUpdateArticle(ctx, &d.Metadata, d.Event); err != nil {
			return xerrors.Errorf(": %+v\n", err)
		}
		return nil
	}); err != nil {
		fmt.Fprintf(os.Stderr, "failed: %+v\n", err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}
