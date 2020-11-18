package main

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/common-go/cgin"
	"github.com/suzuito/terakoya-go/tgin"
)

func main() {
	ctx := context.Background()
	app, err := newApplicationImpl(ctx)
	if err != nil {
		panic(err)
	}
	root := gin.New()
	if os.Getenv("GAE_APPLICATION") == "" {
		root.Use(gin.Logger())
	}
	root.Use(gin.Recovery())
	cgin.UseCORS(app, root)
	tgin.SetUpRoute(root)
	tgin.SetUpGeoRoute(root)
	tgin.SetUpRouteForGeo(root)
	if err := root.Run(); err != nil {
		panic(err)
	}
}
