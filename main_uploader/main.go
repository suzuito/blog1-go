package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/suzuito/blog1-go/inject"
	"github.com/suzuito/blog1-go/local"
	"github.com/suzuito/blog1-go/setting"
	"github.com/suzuito/blog1-go/usecase"
)

func main() {
	ctx := context.Background()
	env, err := setting.NewEnvironment()
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
	gdeps, gcloseFunc, err := inject.NewGlobalDepends(ctx, env)
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
	defer gcloseFunc()
	cdeps, ccloseFunc, err := inject.NewContextDepends(ctx, env)
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
	defer ccloseFunc()

	u := usecase.NewImpl(
		env,
		cdeps.DB,
		cdeps.Storage,
		gdeps.MDConverter,
	)
	mode := flag.String("target", "changed-only", "'all', 'changed-only', 'fixed'")
	flag.Parse()

	var areader usecase.ArticleReader
	switch *mode {
	case "all":
		areader = local.NewArticleReaderAll(path.Join("data", "articles")) // FIXME data => env
	default:
		areaderFix := local.NewArticleReaderFix()
		for _, filePath := range flag.Args() {
			areaderFix.AddFilePath(filePath)
		}
		areader = areaderFix
	}
	defer areader.Close()

	if err := u.SyncArticles(ctx, areader); err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}

	if err := u.WriteArticleHTMLs(ctx, areader); err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
}
