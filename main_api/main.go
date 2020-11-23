package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/suzuito/blog1-go/application"
	"github.com/suzuito/blog1-go/bgin"
	"github.com/suzuito/common-go/cgin"
)

func main() {
	ctx := context.Background()
	app, err := application.NewApplication(ctx)
	if err != nil {
		panic(err)
	}
	root := gin.New()
	cgin.UseCORS(app, root)
	bgin.SetUpRoot(root, app)
	if err := root.Run(); err != nil {
		panic(err)
	}
}
