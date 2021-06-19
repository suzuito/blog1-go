package fdb

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/suzuito/blog1-go/setting"
	"github.com/suzuito/blog1-go/usecase"
	"golang.org/x/xerrors"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var (
	// CollArticles ...
	CollArticles = "articles"
	// CollTags ...
	CollTags = "tags"
)

func NewResource(ctx context.Context, env *setting.Environment) (*firestore.Client, error) {
	cli, err := firestore.NewClient(ctx, env.GCPProjectID)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}
	return cli, nil
}

// Client ...
type Client struct {
	cli *firestore.Client
}

// NewClient ...
func NewClient(cli *firestore.Client) *Client {
	return &Client{
		cli: cli,
	}
}

func newFirestoreOrder(o usecase.CursorOrder) firestore.Direction {
	if o == usecase.CursorOrderAsc {
		return firestore.Asc
	}
	return firestore.Desc
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
	if err := shp.DataTo(doc); err != nil {
		return xerrors.Errorf("%s : %w", err.Error(), err)
	}
	return nil
}

func getDocByTx(
	tx *firestore.Transaction,
	coll *firestore.CollectionRef,
	docID string,
	doc interface{},
) error {
	ref := coll.Doc(docID)
	shp, err := tx.Get(ref)
	if err != nil {
		if grpc.Code(err) == codes.NotFound {
			return xerrors.Errorf("Document '%s' is not found : %w", docID, usecase.ErrNotFound)
		}
		return xerrors.Errorf("%s : %w", err.Error(), usecase.ErrNotFound)
	}
	if err := shp.DataTo(doc); err != nil {
		return xerrors.Errorf("%s : %w", err.Error(), err)
	}
	return nil
}

func getDocs(
	ctx context.Context,
	q *firestore.Query,
	f func(*firestore.DocumentSnapshot) error,
) error {
	it := q.Documents(ctx)
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
		return xerrors.Errorf("Cannot set doc '%s/%s' : %w", coll.Path, docid, err)
	}
	return nil
}

func setDocByTx(tx *firestore.Transaction, coll *firestore.CollectionRef, docid string, v interface{}) error {
	ref := coll.Doc(docid)
	err := tx.Set(ref, v)
	if err != nil {
		return xerrors.Errorf("Cannot set doc '%s/%s' : %w", coll.Path, docid, err)
	}
	return nil
}

func deleteDocByTx(tx *firestore.Transaction, coll *firestore.CollectionRef, docid string) error {
	ref := coll.Doc(docid)
	err := tx.Delete(ref)
	if err != nil {
		return xerrors.Errorf("Cannot delete doc '%s/%s' : %w", coll.Path, docid, err)
	}
	return nil
}
