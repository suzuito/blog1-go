package inject

import (
	"context"
	"net/http"

	"cloud.google.com/go/firestore"
	gstorage "cloud.google.com/go/storage"
	"github.com/pkg/errors"
	"github.com/suzuito/blog1-go/internal/bgcp/fdb"
	"github.com/suzuito/blog1-go/internal/bgcp/storage"
	"github.com/suzuito/blog1-go/internal/bhtml"
	"github.com/suzuito/blog1-go/internal/cmarkdown"
	"github.com/suzuito/blog1-go/pkg/setting"
	"github.com/suzuito/blog1-go/pkg/usecase"
)

func NewGlobalDepends(ctx context.Context) (*usecase.GlobalDepends, func(), error) {
	closeFuncs := []func(){}
	closeFunc := func() {
		for _, cf := range closeFuncs {
			cf()
		}
	}
	r := usecase.GlobalDepends{}
	r.MDConverter = cmarkdown.NewV1()
	gcli, err := gstorage.NewClient(ctx)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "cannot storage.NewClient")
	}
	closeFuncs = append(closeFuncs, func() { gcli.Close() })
	fcli, err := firestore.NewClient(ctx, setting.E.GCPProjectID)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "cannot storage.NewClient")
	}
	closeFuncs = append(closeFuncs, func() { fcli.Close() })
	r.Storage = storage.New(gcli, setting.E.GCPBucketArticle)
	r.DB = fdb.NewClient(fcli)
	r.HTMLEditor = &bhtml.Editor{}
	r.HTMLMediaFetcher = &bhtml.MediaFetcher{Cli: http.DefaultClient}
	r.HTMLTOCExtractor = &bhtml.TOCExtractor{}
	return &r, closeFunc, nil
}
