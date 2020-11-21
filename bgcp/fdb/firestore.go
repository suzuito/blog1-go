package fdb

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/suzuito/blog1-go/usecase"
	"golang.org/x/xerrors"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var (
	// CollArticles ...
	CollArticles = "articles"
)

// ClientFirestore ...
type ClientFirestore struct {
	cli *firestore.Client
}

func getDoc(
	ctx context.Context,
	coll *firestore.CollectionRef,
	docID string,
	doc interface{},
) error {
	ref := coll.Doc(docID)
	shp, err := ref.Get(ctx)
	if err != nil {
		if grpc.Code(err) == codes.NotFound {
			return xerrors.Errorf("Document '%s' is not found : %w", docID, usecase.ErrNotFound)
		}
		return xerrors.Errorf("%s : %w", err.Error(), usecase.ErrNotFound)
	}
	if err := s.DataTo(ctx, doc); err != nil {
		return xerrors.Errorf("%s : %w", err.Error(), err)
	}
	return nil
}

func getDocs(
	ctx context.Context,
	coll *firestore.CollectionRef,
	f func(*firestore.DocumentSnapshot) error,
) error {
	it := coll.Documents(ctx)
	for {
		doc, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		if err := f(doc); err != nil {
			return err
		}
	}
	return nil
}

func setDoc(ctx context.Context, coll *firestore.CollectionRef, docid string, v interface{}) error {
	ref := coll.Doc(docid)
	_, err := ref.Set(ctx, v)
	if err != nil {
		return err
	}
	return nil
}
