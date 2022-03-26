package main

import (
	"bufio"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/google/subcommands"
	"github.com/suzuito/blog1-go/internal/bhtml"
	"github.com/suzuito/blog1-go/internal/cmarkdown"
	"github.com/suzuito/blog1-go/pkg/entity"
	"github.com/suzuito/blog1-go/pkg/usecase"
)

type runConvertCmd struct {
}

func newRunConvertCmd() *runConvertCmd {
	return &runConvertCmd{}
}

func (c *runConvertCmd) Name() string { return "convert" }
func (c *runConvertCmd) Synopsis() string {
	return "MarkdownをHTMLへ変換する\n"
}
func (c *runConvertCmd) Usage() string {
	return c.Synopsis()
}

func (c *runConvertCmd) SetFlags(f *flag.FlagSet) {
}

func (c *runConvertCmd) Execute(ctx context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	scanner := bufio.NewScanner(os.Stdin)
	srcMD := ""
	for scanner.Scan() {
		srcMD += scanner.Text()
		srcMD += "\n"
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return subcommands.ExitFailure
	}
	var u usecase.Usecase
	u = usecase.NewImpl(
		nil, nil,
		cmarkdown.NewV1(),
		&bhtml.MediaFetcher{
			Cli: http.DefaultClient,
		},
		&bhtml.Editor{},
		&bhtml.TOCExtractor{},
	)
	dstHTML := ""
	article := entity.Article{}
	if err := u.ConvertFromMarkdownToHTML(ctx, []byte(srcMD), &dstHTML, &article); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return subcommands.ExitFailure
	}
	fmt.Println(dstHTML)
	articleBytes, err := json.MarshalIndent(&article, "", " ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return subcommands.ExitFailure
	}
	if err := ioutil.WriteFile(fmt.Sprintf("%s.json", article.ID), articleBytes, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}
