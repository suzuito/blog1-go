package main

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/suzuito/blog1-go/application"
	"github.com/suzuito/blog1-go/bgcp/fdb"
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

	db := fdb.NewClient(fcli)
	logger := clogger.LoggerPrint{}

	u := usecase.NewImpl(&logger, db, nil, nil)

	if len(os.Args) <= 1 {
		fmt.Println("You must set site url origin")
		os.Exit(1)
	}
	origin := os.Args[1]

	result, err := u.GenerateBlogSiteMap(ctx, origin)
	if err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
	fmt.Println(result)
}
