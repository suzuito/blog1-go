package inject

import (
	"context"

	"github.com/pkg/errors"
	"github.com/suzuito/blog1-go/internal/bgcp/fdb"
	"github.com/suzuito/blog1-go/internal/bgcp/storage"
	"github.com/suzuito/blog1-go/pkg/setting"
	"github.com/suzuito/blog1-go/pkg/usecase"
	"github.com/suzuito/common-go/cgcp"
	"github.com/suzuito/common-go/cmarkdown"
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
	gcpResources, err := cgcp.NewGCPResourceGenerator().
		GCPS(env.GCPProjectID).
		GCS().
		GCF(env.GCPProjectID).
		Gen(ctx)
	if err != nil {
		return nil, closeFunc, errors.Wrapf(err, "cannot generate google resource clients")
	}
	closeFuncs = append(closeFuncs, func() { gcpResources.Close() })
	r.Storage = storage.New(gcpResources.GCS, env.GCPBucketArticle)
	r.DB = fdb.NewClient(gcpResources.GCF)
	return &r, closeFunc, nil
}
