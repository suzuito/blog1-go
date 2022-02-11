package inject

import (
	"context"

	"github.com/suzuito/blog1-go/internal/bgcp/fdb"
	"github.com/suzuito/blog1-go/internal/bgcp/storage"
	"github.com/suzuito/blog1-go/pkg/setting"
	"github.com/suzuito/blog1-go/pkg/usecase"
	"github.com/suzuito/common-go/cgcp"
	"github.com/suzuito/common-go/cmarkdown"
	"golang.org/x/xerrors"
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
		GCF().
		Gen(ctx)
	if err != nil {
		return nil, closeFunc, xerrors.Errorf("cannot generate google resource clients : %w", err)
	}
	closeFuncs = append(closeFuncs, func() { gcpResources.Close() })
	r.Storage = storage.New(gcpResources.GCS, env.GCPBucketArticle)
	r.DB = fdb.NewClient(gcpResources.GCF)
	return &r, closeFunc, nil
}
