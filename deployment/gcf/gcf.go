package gcf

import (
	"context"
	"os"

	"github.com/getsentry/sentry-go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/suzuito/blog1-go/pkg/inject"
	"github.com/suzuito/blog1-go/pkg/setting"
)

var closeFunc func()
var env *setting.Environment
var gdeps *inject.GlobalDepends

func init() {
	zerolog.LevelFieldName = "severity"
	zerolog.TimestampFieldName = "zerolog_timestamp"
	ctxGlobal := context.Background()
	var err error
	env, err = setting.NewEnvironment()
	if err != nil {
		log.Error().AnErr("message", err).Send()
		return
	}
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:         "https://4344a150d04a4393aa5cb94f1098e1ca@o336494.ingest.sentry.io/6268788",
		Environment: env.Env,
		Release:     os.Getenv("COMMIT_SHA"),
	}); err != nil {
		panic(err)
	}
	gdeps, closeFunc, err = inject.NewGlobalDepends(ctxGlobal, env)
	if err != nil {
		log.Error().AnErr("message", err).Send()
		return
	}
}
