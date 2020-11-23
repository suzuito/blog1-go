package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path"

	"cloud.google.com/go/firestore"
	gstorage "cloud.google.com/go/storage"
	"github.com/suzuito/blog1-go/application"
	"github.com/suzuito/blog1-go/bgcp/fdb"
	"github.com/suzuito/blog1-go/bgcp/storage"
	"github.com/suzuito/blog1-go/local"
	"github.com/suzuito/blog1-go/usecase"
	"github.com/suzuito/common-go/clogger"
)

func main() {
	ctx := context.Background()
	app, err := application.NewApplication(ctx)
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}

	fcli, err := firestore.NewClient(ctx, app.GCPProjectID)
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
	defer fcli.Close()

	scli, err := gstorage.NewClient(ctx)
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
	defer scli.Close()

	areader := local.NewArticleReader(path.Join(app.DirData, "articles"))
	defer areader.Close()

	converter := local.BlackFridayMDConverter{}
	db := fdb.NewClient(fcli)
	str := storage.New(scli, app.GCPBucket)

	logger := clogger.LoggerPrint{}

	u := usecase.NewImpl(&logger, db, str, &converter)

	syncDB := flag.Bool("syncdb", false, "Sync articles")
	syncStorage := flag.Bool("syncstorage", false, "Sync storage")
	flag.Parse()

	if *syncDB {
		if err := u.SyncArticles(ctx, areader); err != nil {
			fmt.Printf("%+v\n", err)
			os.Exit(1)
		}
	}

	if *syncStorage {
		if err := u.UploadArticleMDs(ctx, areader); err != nil {
			fmt.Printf("%+v\n", err)
			os.Exit(1)
		}
	}
}
