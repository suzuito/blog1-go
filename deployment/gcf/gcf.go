package gcf

import (
	"context"
	"os"

	"github.com/getsentry/sentry-go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/suzuito/blog1-go/pkg/inject"
	"github.com/suzuito/blog1-go/pkg/setting"
	"github.com/suzuito/blog1-go/pkg/usecase"
)

var closeFunc func()
var gdeps *usecase.GlobalDepends
var u usecase.Usecase

func init() {
	zerolog.LevelFieldName = "severity"
	zerolog.TimestampFieldName = "zerolog_timestamp"
	ctxGlobal := context.Background()
	var err error
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:         "https://4344a150d04a4393aa5cb94f1098e1ca@o336494.ingest.sentry.io/6268788",
		Environment: setting.E.Env,
		Release:     os.Getenv("COMMIT_SHA"),
	}); err != nil {
		panic(err)
	}
	gdeps, closeFunc, err = inject.NewGlobalDepends(ctxGlobal)
	if err != nil {
		log.Error().AnErr("message", err).Send()
		return
	}
	u = usecase.NewImpl(gdeps.DB, gdeps.Storage, gdeps.MDConverter, gdeps.HTMLMediaFetcher, gdeps.HTMLEditor, gdeps.HTMLTOCExtractor)
}
