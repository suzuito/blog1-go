package main

import (
	"context"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/suzuito/blog1-go/internal/bgin"
	"github.com/suzuito/blog1-go/pkg/inject"
	"github.com/suzuito/blog1-go/pkg/setting"
)

func main() {
	ctx := context.Background()
	env, err := setting.NewEnvironment()
	if err != nil {
		panic(err)
	}
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:         "https://4344a150d04a4393aa5cb94f1098e1ca@o336494.ingest.sentry.io/6268788",
		Environment: env.Env,
		Release:     os.Getenv("COMMIT_SHA"),
	}); err != nil {
		panic(err)
	}
	defer sentry.Flush(2 * time.Second)
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
