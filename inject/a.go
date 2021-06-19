package inject

import (
	"context"

	"github.com/suzuito/blog1-go/bgcp/fdb"
	"github.com/suzuito/blog1-go/bgcp/storage"
	"github.com/suzuito/blog1-go/local"
	"github.com/suzuito/blog1-go/setting"
	"github.com/suzuito/blog1-go/usecase"
	"golang.org/x/xerrors"
)

type GlobalDepends struct {
	MDConverter usecase.MDConverter
}

func NewGlobalDepends(ctx context.Context, env *setting.Environment) (*GlobalDepends, func(), error) {
	closeFuncs := []func(){}
	closeFunc := func() {
		for _, cf := range closeFuncs {
			cf()
		}
	}
	r := GlobalDepends{}
	r.MDConverter = &local.BlackFridayMDConverter{}
	return &r, closeFunc, nil
}

type ContextDepends struct {
	DB      usecase.DB
	Storage usecase.Storage
}

func NewContextDepends(ctx context.Context, env *setting.Environment) (*ContextDepends, func(), error) {
	closeFuncs := []func(){}
	closeFunc := func() {
		for _, cf := range closeFuncs {
			cf()
		}
	}
	r := ContextDepends{}
	cliFirestore, err := fdb.NewResource(ctx, env)
	if err != nil {
		return nil, closeFunc, xerrors.Errorf("%w", err)
	}
	closeFuncs = append(closeFuncs, func() { cliFirestore.Close() })
	cliStorage, err := storage.NewResource(ctx, env)
	if err != nil {
		return nil, closeFunc, xerrors.Errorf("%w", err)
	}
	closeFuncs = append(closeFuncs, func() { cliStorage.Close() })
	r.Storage = storage.New(cliStorage, env.GCPBucketArticle)
	r.DB = fdb.NewClient(cliFirestore)
	return &r, closeFunc, nil
}
