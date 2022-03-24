package inject

import (
	"context"

	"cloud.google.com/go/firestore"
	gstorage "cloud.google.com/go/storage"
	"github.com/pkg/errors"
	"github.com/suzuito/blog1-go/internal/bgcp/fdb"
	"github.com/suzuito/blog1-go/internal/bgcp/storage"
	"github.com/suzuito/blog1-go/internal/cmarkdown"
	"github.com/suzuito/blog1-go/pkg/setting"
	"github.com/suzuito/blog1-go/pkg/usecase"
)

type GlobalDepends struct {
	MDConverter cmarkdown.Converter
	DB          usecase.DB
	Storage     usecase.Storage
}

func NewGlobalDepends(ctx context.Context, env *setting.Environment) (*GlobalDepends, func(), error) {
	closeFuncs := []func(){}
	closeFunc := func() {
		for _, cf := range closeFuncs {
			cf()
		}
	}
	r := GlobalDepends{}
	r.MDConverter = cmarkdown.NewV1()
	gcli, err := gstorage.NewClient(ctx)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "cannot storage.NewClient")
	}
	closeFuncs = append(closeFuncs, func() { gcli.Close() })
	fcli, err := firestore.NewClient(ctx, env.GCPProjectID)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "cannot storage.NewClient")
	}
	closeFuncs = append(closeFuncs, func() { fcli.Close() })
	r.Storage = storage.New(gcli, env.GCPBucketArticle)
	r.DB = fdb.NewClient(fcli)
	return &r, closeFunc, nil
}
