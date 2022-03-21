package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/google/subcommands"
	"github.com/pkg/errors"
	"github.com/suzuito/blog1-go/deployment/gcf"
	"github.com/suzuito/blog1-go/pkg/inject"
	"github.com/suzuito/blog1-go/pkg/setting"
)

type runBlogDeleteArticleCmd struct {
	gdeps   *inject.GlobalDepends
	env     *setting.Environment
	dirBase string
}

func newRunBlogDeleteArticleCmd(gdeps *inject.GlobalDepends, env *setting.Environment) *runBlogDeleteArticleCmd {
	return &runBlogDeleteArticleCmd{gdeps: gdeps, env: env}
}

func (c *runBlogDeleteArticleCmd) Name() string { return "delete-article" }
func (c *runBlogDeleteArticleCmd) Synopsis() string {
	return "記事を削除する\n"
}
func (c *runBlogDeleteArticleCmd) Usage() string {
	return c.Synopsis()
}

func (c *runBlogDeleteArticleCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.dirBase, "input-dir", "", "")
}

func (c *runBlogDeleteArticleCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if err := readDirTest(c.dirBase, func(d *testData) error {
		if err := gcf.BlogDeleteArticle(ctx, &d.Metadata, d.Event); err != nil {
			return errors.Wrapf(err, "cannot get delete article %+v", d)
		}
		return nil
	}); err != nil {
		fmt.Fprintf(os.Stderr, "failed: %+v\n", err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}
