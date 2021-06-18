package main

import (
	"context"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/suzuito/blog1-go/application"
	"github.com/suzuito/blog1-go/bgin"
	"github.com/suzuito/blog1-go/setting"
)

func main() {
	env, err := setting.NewEnvironment()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	app, err := application.NewApplication(ctx)
	if err != nil {
		panic(err)
	}
	root := gin.New()
	root.Use(cors.New(cors.Config{
		AllowOrigins:     env.AllowedOrigins,
		AllowMethods:     env.AllowedMethods,
		AllowHeaders:     []string{},
		ExposeHeaders:    []string{},
		AllowCredentials: false,
	}))
	bgin.SetUpRoot(root, env, app)
	if err := root.Run(); err != nil {
		panic(err)
	}
}
