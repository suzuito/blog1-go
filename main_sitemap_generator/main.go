package main

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/firestore"
	gstorage "cloud.google.com/go/storage"
	"github.com/suzuito/blog1-go/application"
	"github.com/suzuito/blog1-go/bgcp/fdb"
	"github.com/suzuito/blog1-go/bgcp/storage"
	"github.com/suzuito/blog1-go/setting"
	"github.com/suzuito/blog1-go/usecase"
	"github.com/suzuito/common-go/clogger"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You must set site url origin")
		os.Exit(1)
	}
	origin := os.Args[1]

	env, err := setting.NewEnvironment()
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}

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

	db := fdb.NewClient(fcli)
	logger := clogger.LoggerPrint{}
	sscli := storage.New(scli, app.GCPBucket)

	u := usecase.NewImpl(env, &logger, db, sscli, nil)

	result, err := u.GenerateBlogSiteMap(ctx, origin)
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
	body, err := result.Marshal()
	if err != nil {
		fmt.Printf("Cannot marshal xml : %+v\n", err)
		os.Exit(1)
	}
	fmt.Println(body)
}
