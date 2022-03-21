package fdb

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/pkg/errors"
	"github.com/suzuito/blog1-go/pkg/setting"
	"github.com/suzuito/blog1-go/pkg/usecase"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var (
	// CollArticles ...
	CollArticles = "BlogArticles"
	// CollTags ...
	CollTags = "BlogTags"
)

func NewResource(ctx context.Context, env *setting.Environment) (*firestore.Client, error) {
	cli, err := firestore.NewClient(ctx, env.GCPProjectID)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot firestore.NewClient")
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
			return errors.Wrapf(usecase.ErrNotFound, "Document '%s' is not found")
		}
		return errors.Wrapf(err, "document '%s' is failed", docID)
	}
	if err := shp.DataTo(doc); err != nil {
		return errors.Wrapf(err, "cannot DataTo of document '%s'", docID)
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
			return errors.Wrapf(usecase.ErrNotFound, "Document '%s' is not found", docID)
		}
		return errors.Wrapf(err, "getting document '%s' is failed", docID)
	}
	if err := shp.DataTo(doc); err != nil {
		return errors.Wrapf(err, "cannot DataTo of document '%s'", docID)
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
		return errors.Wrapf(err, "cannot set doc '%s/%s'", coll.Path, docid)
	}
	return nil
}

func setDocByTx(tx *firestore.Transaction, coll *firestore.CollectionRef, docid string, v interface{}) error {
	ref := coll.Doc(docid)
	err := tx.Set(ref, v)
	if err != nil {
		return errors.Wrapf(err, "cannot set doc '%s/%s'", coll.Path, docid)
	}
	return nil
}

func deleteDocByTx(tx *firestore.Transaction, coll *firestore.CollectionRef, docid string) error {
	ref := coll.Doc(docid)
	err := tx.Delete(ref)
	if err != nil {
		return errors.Wrapf(err, "cannot delete doc '%s/%s'", coll.Path, docid)
	}
	return nil
}
