package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path"

	"cloud.google.com/go/firestore"
	"github.com/suzuito/blog1-go/application"
	"github.com/suzuito/blog1-go/bgcp/fdb"
	"github.com/suzuito/blog1-go/local"
	"github.com/suzuito/blog1-go/usecase"
)

func main() {
	ctx := context.Background()
	app := application.NewApplication()

	flag.Parse()

	fcli, err := firestore.NewClient(ctx, app.GCPProjectID)
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
	defer fcli.Close()

	areader := local.NewArticleReader(path.Join(app.DirData, "articles"))
	defer areader.Close()

	db := fdb.NewClient(fcli)
	u := usecase.NewImpl(db)
	if err := u.SyncArticles(ctx, areader); err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
}
