package main

import (
	"context"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/suzuito/blog1-go/bgin"
	"github.com/suzuito/blog1-go/inject"
	"github.com/suzuito/blog1-go/setting"
)

func main() {
	ctx := context.Background()
	env, err := setting.NewEnvironment()
	if err != nil {
		panic(err)
	}
	gdeps, closeFunc, err := inject.NewGlobalDepends(ctx, env)
	if err != nil {
		panic(err)
	}
	defer closeFunc()
	root := gin.New()
	root.Use(cors.New(cors.Config{
		AllowOrigins:     env.AllowedOrigins,
		AllowMethods:     env.AllowedMethods,
		AllowHeaders:     []string{},
		ExposeHeaders:    []string{},
		AllowCredentials: false,
	}))
	bgin.SetUpRoot(root, env, gdeps)
	if err := root.Run(); err != nil {
		panic(err)
	}
}
